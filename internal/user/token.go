package user

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wataee/GoQuestions/config"
	"github.com/wataee/GoQuestions/internal/models"
)

func GenerateTokenPair(userID int, username string, role string) models.TokenPair {
	accessToken := generateToken(userID, username, role, 2 * time.Hour)
	refreshToken := generateToken(userID, username, role, 7 * 24 * time.Hour)

	return models.TokenPair{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
}

func generateToken(userID int, username string, role string, ttl time.Duration) string {
	claims := models.UserClaims{
		UserID: userID,
		Username: username,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "goquestions",
		},
	}
	
	createToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := createToken.SignedString(config.JwtKey)
	if err != nil {
		log.Printf("Не удалось создать и подписать токен: %v",err)
	}
	
	return tokenString
}

func RefreshToken(refreshToken string) (models.TokenPair, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JwtKey, nil
	})

	if err != nil || !token.Valid {
		log.Printf("Ошибка при парсинге токена: %v", err)
		return models.TokenPair{}, nil
	}

	claims, ok := token.Claims.(*models.UserClaims)
	if !ok {
		log.Printf("Type asserion ошибка: %v", ok)
		return models.TokenPair{}, nil
	}

	return GenerateTokenPair(claims.UserID, claims.Username, claims.Role), nil

}
 
