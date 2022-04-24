package service

import (
	"context"
	"log"
	"os"

	"github.com/cory-evans/gps-tracker-authentication/pkg/auth"
	"github.com/cory-evans/gps-tracker-authentication/pkg/jwtauth"
	"github.com/cory-evans/gps-tracker-position/pkg/position"
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

	// check if session still exists
	authClient, err := s.getAuthServiceClient()
	if err != nil {
		log.Println("error getting auth service client:", err)
		return ctx, status.Errorf(codes.Internal, "Internal server error.")
	}
	log.Println("INFO: Authenticated request to", fullMethodName)

	resp, err := authClient.SessionIsStillValid(ctx, &auth.SessionIsStillValidRequest{SessionId: jwtauth.GetSessionIdFromContext(ctx)})
	if err != nil {
		return ctx, status.Errorf(codes.Internal, "Internal Server Error: %s", err.Error())
	}

	if !resp.IsValid {
		return ctx, status.Errorf(codes.Unauthenticated, "Session expired or no longer exists.")
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
