package main

import (
    "github.com/wataee/GoQuestions/internal/router"
	"github.com/wataee/GoQuestions/internal/database/repository"
    "github.com/wataee/GoQuestions/internal/database"
)

func main() {
    db,_ := database.ConnectDB()
    repository.NewUserRepository(db)
    r := router.SetupRouter()
    r.Run()
}

