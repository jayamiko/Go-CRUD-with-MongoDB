package model

import "time"

type Admin struct {
    Username     string    `bson:"username" json:"username"`
    Email        string    `bson:"email" json:"email"`
    Password 	 string    `bson:"password" json:"password,omitempty"` 
    Role         string    `bson:"role" json:"role"`
    Status       string    `bson:"status" json:"status"`
    CreatedAt    time.Time `bson:"createdAt" json:"createdAt"`
    UpdatedAt    time.Time `bson:"updatedAt" json:"updatedAt"`
}
