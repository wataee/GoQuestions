package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wataee/GoQuestions/internal/middleware"
	"github.com/wataee/GoQuestions/internal/user"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS()) // ВСЕГДА ПЕРВЫЙ

	userGroup := r.Group("/user")
	{
		userGroup.GET("/login", user.LoginHandler)
		userGroup.POST("/refresh", user.RefreshTokenHandler)
	}

	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddlware())
	{
		protected.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"test": "super",
			})
		})
	}

	return r
}