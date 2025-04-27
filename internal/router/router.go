package router

import (
	"github.com/gin-gonic/gin"


	"github.com/wataee/GoQuestions/internal/middleware"
	"github.com/wataee/GoQuestions/internal/user"
)

func SetupRouter(userService user.UserService) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/login", user.LoginHandler) // Login/Registration
	r.POST("/refresh", user.RefreshTokenHandler)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddlware())
	{
		auth.GET("/profile", )
		auth.GET("/questions", )
		auth.POST("/addquestion", )
	}

	admin := r.Group("/admin")
	// admin.Use(middleware.AdminAuthMiddleware())
	{
		admin.GET("/user_list", )
		admin.DELETE("/delete_user/:id", )
	}
	
	return r
}