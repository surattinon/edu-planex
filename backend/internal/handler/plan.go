package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/surattinon/edu-planex/backend/internal/service"
)

type PlanHandler struct {
	svc *service.PlanService
}

func NewPlanHandler(svc *service.PlanService) *PlanHandler {
	return &PlanHandler{svc: svc}
}

func (h *PlanHandler) GetPlanList(c *gin.Context) {
	plans, err := h.svc.ListPlans()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plans)
}

func (h *PlanHandler) AllPlanTable(c *gin.Context) {
	plans, err := h.svc.GetAllPlanTables()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plans)
}


func (h *PlanHandler) PlanTable(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	plans, err := h.svc.GetPlanTableByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plans)
}

func (h *PlanHandler) GetPlanByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid plan ID"})
		return
	}
	plan, err := h.svc.GetPlan(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "plan not found"})
		return
	}
	c.JSON(http.StatusOK, plan)
}

func (h *PlanHandler) Apply(c *gin.Context) {
	// Apply wtih year and semester number
	var req struct {
		Year       int `json:"year" binding:"required"`
		SemesterNo int `json:"semester_no"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid body"})
		return
	}
	planID, _ := strconv.Atoi(c.Param("id"))

	enroll, err := h.svc.ApplyPlan(uint(planID), req.Year, req.SemesterNo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, enroll)
}

/////

func (h *PlanHandler) List(c *gin.Context) {
  userID := uint(1)
  plans, err := h.svc.ListPlansWithApply(userID)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, plans)
}
