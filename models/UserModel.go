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
	First_name string    `json:"first_name" validate:"required"`
	Last_name  string    `json:"last_name" validate:"required"`
	Email      string    `json:"email" validate:"required" bson:"email"`
	Password   string    `json:"password" validate:"required" bson:"password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	User_id    string    `json:"user_id"`
}

type Error struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type ResponseModel struct {
	Token      string `json:"token"`
	Expires_in int64  `json:"expires_in"`
}

type Code struct {
	Code int `json:"code"`
}

type ForgotPass struct {
	Email string `json:"email"`
}

type NewPassword struct {
	NewPassword     string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserAgent struct {
	UserAgent string `json:"user_agent"`
	Browser   struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"browser"`
	OS struct {
		Platform string `json:"platform"`
		Name     string `json:"name"`
		Version  string `json:"version"`
	} `json:"os"`
	DeviceType string `json:"device_type"`
}

type GoogleAuthenticate struct {
	Email string `json:"email"`
	Code  string `json:"code"`
	URL   string `json:"url"`
}
