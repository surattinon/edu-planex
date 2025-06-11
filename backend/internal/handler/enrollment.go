package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/surattinon/edu-planex/backend/internal/service"
)

type EnrollHandler struct {
	svc *service.EnrollService
}

func NewEnrollHandler(svc *service.EnrollService) *EnrollHandler {
	return &EnrollHandler{svc: svc}
}

func (h *EnrollHandler) GetEnrollList(c *gin.Context) {
	enrollList, err := h.svc.ListEnroll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enrollList)
}

func (h *EnrollHandler) GetEnrollBySemester(c *gin.Context) {
	// read query params: ?year=2025&semester=2
	yearStr := c.Query("year")
	semStr := c.Query("semester")
	year, err1 := strconv.Atoi(yearStr)
	sem, err2 := strconv.Atoi(semStr)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid year or semester"})
		return
	}

	enrolls, err := h.svc.ListBySemester(year, sem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enrolls)
}

func (h *EnrollHandler) GetEnrollByYear(c *gin.Context) {
	// read query params: ?year=2025
	yearStr := c.Query("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid year"})
		return
	}

	enrolls, err := h.svc.ListByYear(year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enrolls)
}

func (h *PlanHandler) Create(c *gin.Context) {
	var req struct {
		Name        string   `json:"name" binding:"required"`
		CourseCodes []string `json:"course_codes" binding:"required,min=1"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hardcode userID because have only 1 user
	userID := uint(1)

	plan, err := h.svc.CreatePlan(userID, req.Name, req.CourseCodes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, plan)
}
