package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(jwtKey []byte) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenstring := ctx.Get("Authorization")

		if tokenstring == "" {
			return fiber.ErrUnauthorized
		}

		tokenstring = strings.TrimPrefix(tokenstring,"Bearer ")
		claims:= &jwt.RegisteredClaims{}
		
		token, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			return fiber.ErrUnauthorized
		}

		ctx.Locals("username",claims.Subject)
		return ctx.Next()
	}
}

