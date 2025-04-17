package middleware

import (
	"time"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	log.Printf("Применение CORS")
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
