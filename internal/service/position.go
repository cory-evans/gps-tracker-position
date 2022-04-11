package service

import (
	"context"

	"github.com/cory-evans/gps-tracker-position/internal/models"
	"github.com/cory-evans/gps-tracker-position/pkg/position"
)

func (s *PositionService) PostPosition(ctx context.Context, req *position.PostPositionRequest) (*position.PostPositionResponse, error) {
	positionsCol := s.DB.Collection(models.POSITION_COLLECTION)
	return nil, nil
}
