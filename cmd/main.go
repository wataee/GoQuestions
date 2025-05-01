package main

import (
	"github.com/wataee/GoQuestions/internal/database"
	"github.com/wataee/GoQuestions/internal/database/repository"
	"github.com/wataee/GoQuestions/internal/router"
	"github.com/wataee/GoQuestions/internal/user"
)


func main() {
    db,_ := database.ConnectDB()
    userRepo := repository.NewUserRepository(db)
    userService := user.NewUserService(userRepo)
    userHandler := user.NewHandler(userService)
    r := router.SetupRouter(userHandler)
    r.Run()
}

