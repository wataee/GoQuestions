package user

import (
	//"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
	// "github.com/wataee/GoQuestions/config"
)

func AuthHandler(с *gin.Context) {
	fmt.Print(GenerateToken(12,"test123","admin"))
	
}