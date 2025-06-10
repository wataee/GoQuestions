package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wataee/GoQuestions/internal/models"
)

type Handler struct {
	service AdminService
}

func NewHandler(service AdminService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) QuestionAddHandler(c *gin.Context) {
	var input models.AddQuestionDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.QuestionText == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пустой вопрос"})
		return
	}

	err := h.service.QuestionsAdd(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) UserListHandler(c *gin.Context) {
	users, err := h.service.UserList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"users":users})
}

func (h *Handler) DeleteUserHandler(c *gin.Context) {
	userID := c.Param("id")
	if err := h.service.DeleteUser(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}