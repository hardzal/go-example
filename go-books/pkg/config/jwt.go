package config

import (
	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("abcdefghijklmnopqrstuvwxyz0987654321")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
