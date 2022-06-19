package service

import (
	"context"
	"fmt"

	"github.com/cory-evans/gps-tracker-authentication/pkg/jwtauth"
	"github.com/cory-evans/gps-tracker-position/internal/models"

	auth "go.buf.build/grpc/go/corux/gps-tracker-auth/auth/v1"
	position "go.buf.build/grpc/go/corux/gps-tracker-position/position/v1"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/metadata"
)

func (s *PositionService) PostPosition(ctx context.Context, req *position.PostPositionRequest) (*position.PostPositionResponse, error) {
	positionsCol := s.DB.Collection(models.POSITION_COLLECTION)

	deviceId := jwtauth.GetUserIdFromContext(ctx)

	// check device exists
	authServiceClient, err := s.getAuthServiceClient()
	if err != nil {
		return nil, err
	}

	md, _ := metadata.FromIncomingContext(ctx)
	authCtx := metadata.NewOutgoingContext(ctx, md)
	getDeviceResp, err := authServiceClient.GetDevice(authCtx, &auth.GetDeviceRequest{
		DeviceId: deviceId,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to get device: %v", err)
	}

	if getDeviceResp.Device == nil {
		return nil, fmt.Errorf("Device does not exist")
	}

	_, err = positionsCol.InsertOne(ctx, &models.Position{
		Latitude:   req.GetLatitude(),
		Longitude:  req.GetLongitude(),
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

func (s *PositionService) GetPosition(ctx context.Context, req *position.GetPositionRequest) (*position.GetPositionResponse, error) {
	positionCol := s.DB.Collection(models.POSITION_COLLECTION)

	var p models.Position
	opts := options.FindOne().SetSort(bson.D{{"created_at", -1}})
	err := positionCol.FindOne(ctx, bson.M{"device_id": req.GetDeviceId()}, opts).Decode(&p)
	if err != nil {
		return nil, err
	}

	return &position.GetPositionResponse{
		Position: p.AsProtoBuf(),
	}, nil
}

func (s *PositionService) GetOwnedDevicesPositions(ctx context.Context, req *position.GetOwnedDevicesPositionsRequest) (*position.GetOwnedDevicesPositionsResponse, error) {
	authServiceClient, err := s.getAuthServiceClient()
	if err != nil {
		return nil, err
	}

	md, _ := metadata.FromIncomingContext(ctx)
	authCtx := metadata.NewOutgoingContext(ctx, md)
	devices, err := authServiceClient.GetOwnedDevices(authCtx, &auth.GetOwnedDevicesRequest{})
	if err != nil {
		return nil, err
	}

	positionCol := s.DB.Collection(models.POSITION_COLLECTION)

	positions := make([]*position.Position, len(devices.Devices))

	opts := options.FindOne().SetSort(bson.D{{"created_at", -1}})

	for i, device := range devices.Devices {
		var p models.Position
		err := positionCol.FindOne(ctx, bson.M{"device_id": device.DeviceId}, opts).Decode(&p)
		if err != nil {
			continue
		}

		positions[i] = p.AsProtoBuf()
	}

	return &position.GetOwnedDevicesPositionsResponse{
		Positions: positions,
	}, nil
}

func (s *PositionService) GetRecentPosition(ctx context.Context, req *position.GetRecentPositionRequest) (*position.GetRecentPositionResponse, error) {
	positionsCol := s.DB.Collection(models.POSITION_COLLECTION)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("Failed to get metadata from context")
	}

	// check device exists
	authServiceClient, err := s.getAuthServiceClient()
	if err != nil {
		return nil, err
	}

	authCtx := metadata.NewOutgoingContext(ctx, md)
	getDeviceResp, err := authServiceClient.GetDevice(authCtx, &auth.GetDeviceRequest{
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
