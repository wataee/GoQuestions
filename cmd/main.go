package main

import (
    "github.com/gin-gonic/gin"

    "goquestions/internal/middleware"
    "goquestions/internal/auth"
)

func main() {
    r := gin.Default()

    r.Use(middleware.CORS()) // ВСЕГДА ПЕРВЫЙ

    r.GET("/auth", )

    r.Run()
}

