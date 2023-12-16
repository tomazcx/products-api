package utils

import (
	"time"

	"github.com/go-chi/jwtauth"
)

func GenerateJWT(userEmail string, exp int, jwt *jwtauth.JWTAuth) (string, error) {
	_, tokenString, err := jwt.Encode(map[string]interface{}{
		"sub": userEmail,
		"exp": int(time.Second) * exp,
	})
	return tokenString, err
}
