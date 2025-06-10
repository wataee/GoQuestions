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

// @Summary Add question
// @Description Adds a new question to the system
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body models.AddQuestionDTO true "Question data"
// @Success 200 "Question added successfully"
// @Failure 400 {object} map[string]string
// @Router /admin/addquestion [post]
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

// @Summary Get user list
// @Description Returns a list of all users
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string][]models.ProfileDTO
// @Failure 400 {object} map[string]string
// @Router /admin/user_list [get]
func (h *Handler) UserListHandler(c *gin.Context) {
	users, err := h.service.UserList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"users":users})
}

// @Summary Delete user
// @Description Deletes user by ID
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 "User deleted successfully"
// @Failure 400 {object} map[string]string
// @Router /admin/delete_user/{id} [delete]
func (h *Handler) DeleteUserHandler(c *gin.Context) {
	userID := c.Param("id")
	if err := h.service.DeleteUser(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}