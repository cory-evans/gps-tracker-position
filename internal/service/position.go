package service

import (
	"context"
	"fmt"

	"github.com/cory-evans/gps-tracker-authentication/pkg/auth"
	"github.com/cory-evans/gps-tracker-authentication/pkg/jwtauth"
	"github.com/cory-evans/gps-tracker-position/internal/models"
	"github.com/cory-evans/gps-tracker-position/pkg/position"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/metadata"
)

func (s *PositionService) PostPosition(ctx context.Context, req *position.PostPositionRequest) (*position.PostPositionResponse, error) {
	positionsCol := s.DB.Collection(models.POSITION_COLLECTION)

	// request has to be authenticated
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("Failed to get metadata from context")
	}

	deviceId := jwtauth.GetUserIdFromMetadata(md)

	// check device exists
	authServiceClient, err := s.getAuthServiceClient()
	if err != nil {
		return nil, err
	}
	getDeviceResp, err := authServiceClient.GetDevice(ctx, &auth.GetDeviceRequest{
		DeviceId: deviceId,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to get device: %v", err)
	}

	if getDeviceResp.Device == nil {
		return nil, fmt.Errorf("Device does not exist")
	}

	_, err = positionsCol.InsertOne(ctx, &models.Position{
		Latitude:   req.GetLongitude(),
		Longitude:  req.GetLatitude(),
		Altitude:   req.GetAltitude(),
		SpeedKnots: req.GetSpeedKnots(),
		Accuracy:   req.GetAccuracy(),
		DeviceId:   getDeviceResp.Device.DeviceId,
	})

	if err != nil {
		return nil, err
	}

	return &position.PostPositionResponse{}, nil
}

func (s *PositionService) GetRecentPosition(ctx context.Context, req *position.GetRecentPositionRequest) (*position.GetRecentPositionResponse, error) {
	positionsCol := s.DB.Collection(models.POSITION_COLLECTION)

	// request has to be authenticated
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("Failed to get metadata from context")
	}

	userId := jwtauth.GetUserIdFromMetadata(md)
	if userId == "" {
		return nil, fmt.Errorf("Failed to get user id from metadata")
	}

	// check device exists
	authServiceClient, err := s.getAuthServiceClient()
	if err != nil {
		return nil, err
	}
	getDeviceResp, err := authServiceClient.GetDevice(ctx, &auth.GetDeviceRequest{
		DeviceId: req.GetDeviceId(),
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to get device: %v", err)
	}

	if getDeviceResp.Device == nil {
		return nil, fmt.Errorf("Device does not exist")
	}

	// get the most recent position
	var elements = make([]*position.Position, 0)
	cur, err := positionsCol.Find(
		ctx,
		bson.M{
			"device_id": getDeviceResp.Device.DeviceId,
		},
		options.Find().SetLimit(25),
	)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var elem models.Position
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		elements = append(elements, elem.AsProtoBuf())
	}

	return &position.GetRecentPositionResponse{
		Positions: elements,
	}, nil
}
