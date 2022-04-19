package service

import (
	"context"
	"log"
	"math/rand"
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

	return ctx, nil
}

func (s *PositionService) getAuthServiceClient() (auth.AuthServiceClient, error) {
	authConn, err := grpc.Dial(os.Getenv("AUTH_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return auth.NewAuthServiceClient(authConn), nil
}

func (s *PositionService) GetPosition(ctx context.Context, req *position.GetPositionRequest) (*position.GetPositionResponse, error) {
	// TODO: fetch position from DB
	return &position.GetPositionResponse{
		Position: &position.Position{
			Latitude:  float64(rand.Intn(90) - 45),
			Longitude: float64(rand.Intn(180) - 90),
		},
	}, nil
}

func (s *PositionService) GetOwnedDevicesPosition(ctx context.Context, req *position.GetOwnedDevicesPositionRequest) (*position.GetOwnedDevicesPositionResponse, error) {
	userId := jwtauth.GetUserIdFromContext(ctx)

	authServiceClient, err := s.getAuthServiceClient()
	if err != nil {
		return nil, err
	}

	devices, err := authServiceClient.GetOwnedDevices(ctx, &auth.GetOwnedDevicesRequest{})
	if err != nil {
		return nil, err
	}

	positions := make([]*position.Position, len(devices.Devices))

	// TODO: fetch positions from DB
	for i, device := range devices.Devices {
		positions[i] = &position.Position{
			Latitude:  float64(rand.Intn(90) - 45),
			Longitude: float64(rand.Intn(180) - 90),
		}
	}

	return &position.GetOwnedDevicesPositionResponse{
		Positions: positions,
	}, nil
}
