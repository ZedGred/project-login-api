package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User
type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Username      string             `json:"Username" validate:"required,min=2,max=100"`
	Password      string             `json:"Password" validate:"required,min=6"`
	Email         string             `json:"email" validate:"email,required"`
	Phone         string             `json:"phone"`
	Token         *string             `json:"token"`
	RefreshToken  *string             `json:"refresh_token"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
}
