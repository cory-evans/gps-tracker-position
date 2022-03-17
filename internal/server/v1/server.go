package serverv1

import (
	"context"
	"math/rand"

	positionv1 "github.com/cory-evans/gps-tracker-position/pkg/position/v1"
)

type PositionService struct {
	positionv1.PositionServiceServer
}

func (s *PositionService) GetPosition(ctx context.Context, req *positionv1.GetPositionRequest) (*positionv1.GetPositionResponse, error) {
	return &positionv1.GetPositionResponse{
		Position: &positionv1.Position{
			Latitude:  float64(rand.Intn(90) - 45),
			Longitude: float64(rand.Intn(180) - 90),
		},
	}, nil
}
