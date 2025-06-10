package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/wataee/GoQuestions/internal/admin"
	"github.com/wataee/GoQuestions/internal/database"
	"github.com/wataee/GoQuestions/internal/database/repository"
	"github.com/wataee/GoQuestions/internal/questions"
	"github.com/wataee/GoQuestions/internal/router"
	"github.com/wataee/GoQuestions/internal/user"
)

// @title GoQuestions Restful API
// @version 1.0
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal(err)
    }
    db,err := database.ConnectDB()
    if err != nil {
        log.Fatal(err)
    }
    userRepo := repository.NewUserRepository(db)
    userService := user.NewUserService(userRepo)
    userHandler := user.NewHandler(userService)

    questionsRepo := repository.NewQuestionsRepository(db)
    questionsService := questions.NewQuestionsService(questionsRepo)
    questionsHandler := questions.NewHandler(questionsService)

    adminService := admin.NewAdminService(userRepo, questionsRepo)
    adminHandler := admin.NewHandler(adminService)

    r := router.SetupRouter(userHandler, questionsHandler, adminHandler)
    r.Run()
}