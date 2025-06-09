package questions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wataee/GoQuestions/internal/models"
)


type Handler struct {
	service QuestionsService
}

func NewHandler(service QuestionsService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) QuestionsListHandler(c *gin.Context) {
	questions, err := h.service.QuestionsList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"questions": questions})
}

func (h *Handler) QuestionAddHandler(c *gin.Context) {
	var input models.AddQuestionDTO

	if err := c.BindJSON(&input); err != nil {
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