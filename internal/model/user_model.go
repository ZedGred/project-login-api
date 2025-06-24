package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `json:"name"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Token    string             `json:"token"`
	CreateAt time.Time          `json:"create_at"`
	UpdateAt time.Time          `json:"update_at"`
}

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
