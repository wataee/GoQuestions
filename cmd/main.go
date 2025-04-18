package main

import (
    "github.com/wataee/GoQuestions/internal/router"
	"github.com/wataee/GoQuestions/internal/database"
)

func main() {
    database.ConnectDB()
    r := router.SetupRouter()
    r.Run()
}

