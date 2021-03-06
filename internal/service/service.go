package service

import (
	"context"
	"log"
	"os"

	auth "go.buf.build/grpc/go/corux/gps-tracker-auth/auth/v1"
	position "go.buf.build/grpc/go/corux/gps-tracker-position/position/v1"

	"github.com/cory-evans/gps-tracker-authentication/pkg/jwtauth"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type PositionService struct {
	position.PositionServiceServer

	DB *mongo.Database
}

func (s *PositionService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	ctx, err := jwtauth.MapJWT(ctx)
	if err != nil {
		log.Println("error mapping JWT to context metadata:", err)

		// all requests are authenticated
		return ctx, status.Errorf(codes.Unauthenticated, "Not authenticated.")
	}

	return ctx, nil
}

func (s *PositionService) getAuthServiceClient() (auth.AuthServiceClient, error) {
	authConn, err := grpc.Dial(os.Getenv("AUTH_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return auth.NewAuthServiceClient(authConn), nil
}
