package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VerifyTokenRequest struct {
	Token string
}

type Token struct {
	Token   string
	Expires time.Time
}

type SignedDetails struct{
	Email string
	Username string
	ID primitive.ObjectID
	jwt.RegisteredClaims
}
