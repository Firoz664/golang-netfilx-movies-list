package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie       string             `json:"movie,omitempty"`
	Watched     bool               `json:"watched,omitempty"`
	Genre       string             `json:"genre,omitempty"`
	ReleaseYear int                `json:"releaseYear,omitempty"`
	Rating      *float64           `json:"rating,omitempty"`
	Duration    int                `json:"duration,omitempty"` // Duration in minutes
}
