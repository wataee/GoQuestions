package router

import (
	"github.com/gin-gonic/gin"

	"github.com/wataee/GoQuestions/internal/admin"
	"github.com/wataee/GoQuestions/internal/middleware"
	"github.com/wataee/GoQuestions/internal/questions"
	"github.com/wataee/GoQuestions/internal/user"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "github.com/wataee/GoQuestions/docs"
)

func SetupRouter(userHandler *user.Handler, questionsHandler *questions.Handler, adminHandler *admin.Handler) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/login", userHandler.Login) // Login/Registration
	r.POST("/refresh", userHandler.RefreshToken)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", userHandler.Profile)
		auth.GET("/questions", questionsHandler.QuestionsListHandler)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminAuthMiddleware())
	{
		admin.GET("/user_list", adminHandler.UserListHandler)
		admin.DELETE("/delete_user/:id", adminHandler.DeleteUserHandler)
		admin.POST("/addquestion", adminHandler.QuestionAddHandler)
	}
	
	return r
}