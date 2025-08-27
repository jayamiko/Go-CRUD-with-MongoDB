package model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrMemberNotFound = errors.New("member not found")

type Member struct {
	ID    		   primitive.ObjectID `bson:"_id" json:"_id"`
	RecruiterID    primitive.ObjectID `bson:"recruiterId" json:"recruiterId"`
	StatusAktivasi string             `bson:"statusAktivasi" json:"statusAktivasi"` // ACTIVE, INACTIVE, PENDING
	Email          string             `bson:"email" json:"email"`
	Password       string             `bson:"password" json:"password,omitempty"`             // password hash
	CreatedAt      time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt" json:"updatedAt"`
}