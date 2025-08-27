package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
    ID    		 primitive.ObjectID `bson:"_id" json:"_id"`
    Username     string    `bson:"username" json:"username"`
    Email        string    `bson:"email" json:"email"`
    Password 	 string    `bson:"password" json:"password,omitempty"` 
    Role         string    `bson:"role" json:"role"`
    Status       string    `bson:"status" json:"status"`
    CreatedAt    time.Time `bson:"createdAt" json:"createdAt"`
    UpdatedAt    time.Time `bson:"updatedAt" json:"updatedAt"`
}
