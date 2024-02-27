package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Session struct {
	RefreshToken string    `json:"refresh_token" bson:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at" bson:"expires_at"`
}

type Verification struct {
	Code     primitive.ObjectID `json:"code" bson:"code"`
	Verified bool               `json:"verified" bson:"verified"`
}
