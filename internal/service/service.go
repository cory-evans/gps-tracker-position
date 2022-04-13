package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/cory-evans/gps-tracker-authentication/pkg/auth"
	"github.com/cory-evans/gps-tracker-authentication/pkg/jwtauth"
	"github.com/cory-evans/gps-tracker-position/pkg/position"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/metadata"
)

type PositionService struct {
	position.PositionServiceServer

	DB                *mongo.Database
	AuthServiceClient auth.AuthServiceClient
}

func (s *PositionService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	// map incomming JWT to context metadata
	ctx, err := jwtauth.MapJWT(ctx)
	if err != nil {
		log.Println("error mapping JWT to context metadata:", err)
	}

	return ctx, nil
}

func (s *PositionService) GetPosition(ctx context.Context, req *position.GetPositionRequest) (*position.GetPositionResponse, error) {
	// request has to be authenticated
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("Failed to get metadata from context")
	}
	userId := jwtauth.GetUserIdFromMetadata(md)
	if userId == "" {
		return nil, fmt.Errorf("Not authenticated.")
	}
	return &position.GetPositionResponse{
		Position: &position.Position{
			Latitude:  float64(rand.Intn(90) - 45),
			Longitude: float64(rand.Intn(180) - 90),
		},
	}, nil
}

func (s *PositionService) GetOwnedDevicesPosition(ctx context.Context, req *position.GetOwnedDevicesPositionRequest) (*position.GetOwnedDevicesPositionResponse, error) {
	// request has to be authenticated
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("Failed to get metadata from context")
	}
	userId := jwtauth.GetUserIdFromMetadata(md)
	if userId == "" {
		return nil, fmt.Errorf("Not authenticated.")
	}
	return &position.GetOwnedDevicesPositionResponse{
		Positions: []*position.Position{
			{
				Latitude:  float64(rand.Intn(90) - 45),
				Longitude: float64(rand.Intn(180) - 90),
			},
		},
	}, nil
}
