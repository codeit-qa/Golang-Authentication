package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    string             `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     string             `json:"last_name" validate:"required,min=2,max=100"`
	Password      string             `json:"Password" validate:"required,min=6,max=20"`
	Email         string             `json:"email" validate:"email,required"`
	Phone         int                `json:"phone" validate:"required"`
	Token         string             `json:"token"`
	Refresh_token string             `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}

type AuthenticationModel struct {
	Email    string `json:"email" validate:"required" bson:"email"`
	Password string `json:"password" validate:"required" bson:"password"`
}
