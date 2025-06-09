package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserInputDTO struct {
	Username string `json:"username" validate:"required,min=3,max=38,alphanum"`
	Password string `json:"password" validate:"required,min=7,max=128"`
	Role     string `json:"role" validate:"required,oneof=admin user"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type ProfileDTO struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Answers   int    `json:"answers"`
	CreatedAt time.Time `json:"created_at"`
}

type TokenPair struct {
	AccessToken string `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
}

type UserClaims struct {
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}
