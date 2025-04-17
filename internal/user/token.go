package user

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wataee/GoQuestions/config"
)



func GenerateToken(userID int, username string, role string) string {
	claims := UserClaims{
		UserID: userID,
		Username: username,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "goquestions",
		},
	}
	
	createToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := createToken.SignedString(config.JwtKey)
	if err != nil {
		log.Fatalf("Не удалось создать и подписать токен: %v",err)
	}
	
	return tokenString
}

