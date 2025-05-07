package questions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Handler struct {
	service QuestionsService
}

func NewHandler(service QuestionsService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) QuestionsHandler(c *gin.Context) {
	questions, err := h.service.QuestionsList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"questions": questions})
}