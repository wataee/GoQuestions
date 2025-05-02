package models

type UserInput struct {
	Username string `json:"username" validate:"required,min=3,max=38,alphanum"`
	Password string `json:"password" validate:"required,min=7,max=128"`
	Role string `json:"role" validate:"required,oneof=admin user"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
