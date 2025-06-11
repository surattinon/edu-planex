package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surattinon/edu-planex/backend/internal/service"
)

type CourseHandler struct {
	svc *service.CourseService
}

func NewCourseHandler(svc *service.CourseService) *CourseHandler {
	return &CourseHandler{svc: svc}
}

func (h *CourseHandler) GetCourseList(c *gin.Context) {
	cs, err := h.svc.ListAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cs)
}


func (h *CourseHandler) CourseTable(c *gin.Context) {
	cs, err := h.svc.ListSummariesWithPrereqs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cs)
}

func (h *CourseHandler) GetCourseByCode(c *gin.Context) {
	code := c.Param("code")
	course, err := h.svc.FindCourseByCode(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) GetCatByCode(c *gin.Context) {
    course, err := h.svc.FindCourseByCode(c.Param("code"))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error":"not found"})
        return
    }
    c.JSON(http.StatusOK, course.Categories)
}
