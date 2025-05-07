package router

import (
	"github.com/gin-gonic/gin"


	"github.com/wataee/GoQuestions/internal/middleware"
	"github.com/wataee/GoQuestions/internal/user"
	"github.com/wataee/GoQuestions/internal/questions"
)

func SetupRouter(userHandler *user.Handler, questionsHandler *questions.Handler) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/login", userHandler.Login) // Login/Registration
	r.POST("/refresh", userHandler.RefreshToken)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", userHandler.Profile)
		auth.GET("/questions", questionsHandler.QuestionsHandler)
		auth.POST("/addquestion", )
	}

	admin := r.Group("/admin")
	//admin.Use(middleware.AdminAuthMiddleware())
	{
		admin.GET("/user_list", )
		admin.DELETE("/delete_user/:id", )
	}
	
	return r
}