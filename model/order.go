package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    ProductID   primitive.ObjectID `bson:"productId" json:"productId"`
    MemberID    primitive.ObjectID `bson:"memberId" json:"memberId"`
    RecruiterID primitive.ObjectID `bson:"recruiterId" json:"recruiterId"`
    Status      string             `bson:"status" json:"status"`
    TotalAmount float64            `bson:"totalAmount,omitempty" json:"totalAmount,omitempty"`
    CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
    UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}
