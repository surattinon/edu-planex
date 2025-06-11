package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surattinon/edu-planex/backend/internal/model"
	"github.com/surattinon/edu-planex/backend/internal/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	u, err := h.svc.GetProfile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var input model.UserProfile
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	// Forced to 1 because have only 1 user
	input.UserID = 1

	updated, err := h.svc.UpdateProfile(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}
