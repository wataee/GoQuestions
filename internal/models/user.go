package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
}

type UserClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken string `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
}
