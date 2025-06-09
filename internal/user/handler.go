package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/wataee/GoQuestions/internal/models"
)

type Handler struct {
	service UserService
}

func NewHandler(service UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Login(c *gin.Context) {
	validate := validator.New()
	var errorMessages []string
	var input models.UserInputDTO

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(input); err != nil {
		for _,e := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, fmt.Sprintf("Поле: %s | Тег валидации: %s", e.Field(), e.Tag()))
		}
	}
	if len(errorMessages) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessages})
		return
	}

	tokens, err := h.service.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var RefreshTokenRequest models.RefreshTokenRequest
	if err := c.ShouldBindJSON(&RefreshTokenRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenPair, err := h.service.RefreshToken(RefreshTokenRequest.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokenPair)
}

func (h *Handler) Profile(c *gin.Context) {
	UserID, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось взять UserID из userclaims"})
	}
	profile,err := h.service.GetProfile(UserID.(int))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, profile)
}