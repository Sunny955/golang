package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// json:"_id,omitempty

type Netflix struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
