package middleware

import (
	"net/http"
	"strings"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"


	"github.com/wataee/GoQuestions/config"
)

func AuthMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Токен не найден в Header/Authorization"})
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Fatal("Ошибка при подписи токена")
			}
			return config.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Токен недействителен"})
		}

		c.Next()
	}
}
