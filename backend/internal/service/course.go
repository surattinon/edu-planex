package service

import (
	"errors"

	"github.com/surattinon/edu-planex/backend/internal/model"
)

func FindCourseByID(code *string) (*model.Courses, error) {
	for i, c := range Courses {
		if c.Code == *code {
			return &Courses[i], nil
		}
	}

	return nil, errors.New("Course not found")
}
