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

// @Summary Login or Registration
// @Description The handler registers or logs in the user, returning a refresh and access token.
// @Tags user
// @Accept json
// @Produce json
// @Param input body models.UserInputDTO true "Input values"
// @Success 200 {object} models.TokenPair
// @Failure 400 {object} map[string]any
// @Failure 400 {object} map[string]any
// @Router /login [post]
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

// @Summary Refresh Token
// @Description Refreshes access and refresh tokens using a valid refresh token
// @Tags user
// @Accept json
// @Produce json
// @Param input body models.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} models.TokenPair
// @Failure 400 {object} map[string]any
// @Failure 401 {object} map[string]any
// @Router /refresh [post]
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

// @Summary Get user profile
// @Description Returns the profile of the authenticated user
// @Tags user
// @Security BearerAuth
// @Produce json
// @Success 200 {object} models.ProfileDTO
// @Failure 401 {object} map[string]any
// @Security BearerAuth
// @Router /profile [get]
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