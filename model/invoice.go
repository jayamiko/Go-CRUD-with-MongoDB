package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	OrderID       primitive.ObjectID `bson:"orderId" json:"orderId"`
	Amount        float64            `bson:"amount,omitempty" json:"amount,omitempty"`
	Status        string             `bson:"status" json:"status"` // DRAFT, FINAL, CANCELLED
	PaymentStatus string             `bson:"paymentStatus" json:"paymentStatus"` // UNPAID, PAID, PARTIAL
	IssuedAt      *time.Time         `bson:"issuedAt,omitempty" json:"issuedAt,omitempty"`
	PaidAt        *time.Time         `bson:"paidAt,omitempty" json:"paidAt,omitempty"`
	CreatedAt     time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt" json:"updatedAt"`
}
