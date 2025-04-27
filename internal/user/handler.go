package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func LoginHandler(с *gin.Context) {
	fmt.Println(GenerateTokenPair(13,"lol","admin"))
	
}

func RefreshTokenHandler(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось взять RefreshToken из JSON"})
		return
	}

	tokenPair, err := RefreshToken(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка создания пары токенов"})
		return
	}

	c.JSON(http.StatusOK, tokenPair)
}