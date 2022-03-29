package servicev1

import (
	"context"
	"fmt"
	"math/rand"

	jwtauthv1 "github.com/cory-evans/gps-tracker-authentication/pkg/jwtauth/v1"
	positionv1 "github.com/cory-evans/gps-tracker-position/pkg/position/v1"
	"google.golang.org/grpc/metadata"
)

type PositionService struct {
	positionv1.PositionServiceServer
}

func (s *PositionService) GetPosition(ctx context.Context, req *positionv1.GetPositionRequest) (*positionv1.GetPositionResponse, error) {
	// request has to be authenticated
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("Failed to get metadata from context")
	}
	userId := jwtauthv1.GetUserIdFromMetadata(md)
	if userId == "" {
		return nil, fmt.Errorf("Not authenticated.")
	}
	return &positionv1.GetPositionResponse{
		Position: &positionv1.Position{
			Latitude:  float64(rand.Intn(90) - 45),
			Longitude: float64(rand.Intn(180) - 90),
		},
	}, nil
}
