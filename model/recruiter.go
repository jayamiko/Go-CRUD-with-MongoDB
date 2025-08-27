package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recruiter struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Phone     string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Status    string             `bson:"status" json:"status"` 
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
