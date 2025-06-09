package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service AdminService
}

func NewHandler(service AdminService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) UserListHandler(c *gin.Context) {
	users, err := h.service.UserList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"users":users})
}