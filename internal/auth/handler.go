package auth

import (
	//"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wataee/GoQuestions/config"
)

func AuthHandler(с *gin.Context) {
	rff, _ := GenerateToken(12,"test123","admin")
	token, _ := jwt.Parse(rff, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Еррор 14 строка")
		}
		return config.JwtKey, nil

	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user_id"])
		fmt.Println(claims["username"])
		fmt.Println(claims["role"])
	} else {
		fmt.Println("Невалидный токен")
	}
	
	
}