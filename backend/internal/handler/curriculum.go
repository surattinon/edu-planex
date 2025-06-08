package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surattinon/edu-planex/backend/internal/service"
)

func GetCourses(c *gin.Context) {
	if err := c.Bind(&service.Courses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, service.Courses)
}

func GetCourseByID(c *gin.Context) {
	code := c.Param("code")
	course, err := service.FindCourseByID(&code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, course)
}
