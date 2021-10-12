package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rol struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Rol string `bson:"rol" json:"rol,omitempty"`
} 