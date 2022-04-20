package models

import (
	"time"

	"github.com/cory-evans/gps-tracker-position/pkg/position"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	POSITION_COLLECTION = "device_positions"
)

type Position struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	Latitude   float64 `json:"latitude" bson:"latitude"`
	Longitude  float64 `json:"longitude" bson:"longitude"`
	Altitude   float64 `json:"altitude" bson:"altitude"`
	SpeedKnots float64 `json:"speed_knots" bson:"speed_knots"`
	Accuracy   float64 `json:"accuracy" bson:"accuracy"`
	DeviceId   string  `json:"device_id" bson:"device_id"`
}

func (p *Position) MarshalBSON() ([]byte, error) {
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now().UTC()
	}

	type pos Position
	return bson.Marshal((*pos)(p))
}

func (p *Position) AsProtoBuf() *position.Position {
	return &position.Position{
		CreatedAt:  p.CreatedAt.Format(time.RFC3339),
		Latitude:   p.Latitude,
		Longitude:  p.Longitude,
		Altitude:   p.Altitude,
		SpeedKnots: p.SpeedKnots,
		Accuracy:   p.Accuracy,
		DeviceId:   p.DeviceId,
	}
}
