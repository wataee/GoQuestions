package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wataee/GoQuestions/internal/user"
	"github.com/wataee/GoQuestions/internal/database"
	"github.com/wataee/GoQuestions/internal/middleware"
)

func main() {
    r := gin.Default()
    database.ConnectDB()
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

    r.Run()
}

