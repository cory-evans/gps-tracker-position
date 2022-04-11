package models

import "github.com/cory-evans/gps-tracker-position/pkg/position"

const (
	POSITION_COLLECTION = "device_positions"
)

type Position struct {
	Latitude   float64 `json:"latitude" bson:"latitude"`
	Longitude  float64 `json:"longitude" bson:"longitude"`
	Altitude   float64 `json:"altitude" bson:"altitude"`
	SpeedKnots float64 `json:"speed_knots" bson:"speed_knots"`
	Accuracy   float64 `json:"accuracy" bson:"accuracy"`
	DeviceId   string  `json:"device_id" bson:"device_id"`
}

func (p *Position) AsProtoBuf() *position.Position {
	return &position.Position{
		Latitude:   p.Latitude,
		Longitude:  p.Longitude,
		Altitude:   p.Altitude,
		SpeedKnots: p.SpeedKnots,
		Accuracy:   p.Accuracy,
		DeviceId:   p.DeviceId,
	}
}
