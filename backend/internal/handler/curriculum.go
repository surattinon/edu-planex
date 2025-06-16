package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surattinon/edu-planex/backend/internal/service"
)

type CurriculumHandler struct {
	svc *service.CurriculumService
}

func NewCurriculumHandler(svc *service.CurriculumService) *CurriculumHandler {
	return &CurriculumHandler{svc}
}

func (h *CurriculumHandler) List(c *gin.Context) {
	cur, err := h.svc.GetCurriculum()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cur)
}

func (h *CurriculumHandler) Personal(c *gin.Context) {
  uid := uint(1) // or parse from c.Query("user_id") / auth
  dto, err := h.svc.GetPersonalCurriculum(uid)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, dto)
}
