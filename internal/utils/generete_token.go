package utils

import (
	"project/internal/entity"
	"project/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenereteAllToken(user *entity.User, jwtKey []byte) (string, string, error) {
	claims := model.SignedDetails{
		Username: user.Username,
		ID:       user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(24 * time.Hour)),
		},
	}

	refreshClaims := model.SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(7 * 24 * time.Hour)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		return "", "", nil
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(jwtKey)
	if err != nil {
		return "", "", nil
	}

	return token, refreshToken, nil
}
