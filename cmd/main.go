package main

import (
	"github.com/gin-gonic/gin"

	"github.com/wataee/GoQuestions/internal/auth"
	"github.com/wataee/GoQuestions/internal/database"
	"github.com/wataee/GoQuestions/internal/middleware"
)

func main() {
    r := gin.Default()
    database.ConnectDB()


    r.Use(middleware.CORS()) // ВСЕГДА ПЕРВЫЙ
    r.Use(middleware.AuthMiddlware())
    
    r.GET("/auth", auth.AuthHandler)
    r.Run()
}

