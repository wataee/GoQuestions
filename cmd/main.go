package main

import (
    "github.com/gin-gonic/gin"

    "goquestions/internal/middleware"
)

func main() {
    r := gin.Default()

    r.Use(middleware.CORS()) // ВСЕГДА ПЕРВЫЙ
}

