package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wataee/GoQuestions/internal/models"
)

type Handler struct {
	service UserService
}

func NewHandler(service UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Login(c *gin.Context) {
	var input models.UserInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error":"Не удалось распарсить json"})
		return
	}
	tokens, err := h.service.Login(input)
	fmt.Println(tokens, err)
}






func (h *Handler) RefreshToken(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось взять RefreshToken из JSON"})
		return
	}

	tokenPair, err := h.service.RefreshToken(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка создания пары токенов"})
		return
	}

	c.JSON(http.StatusOK, tokenPair)
}