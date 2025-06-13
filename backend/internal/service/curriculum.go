package service

import (
	"sort"

	"github.com/surattinon/edu-planex/backend/internal/dto"
	"github.com/surattinon/edu-planex/backend/internal/model"
	"gorm.io/gorm"
)

type CurriculumService struct {
	db *gorm.DB
}

func NewCurriculumService(db *gorm.DB) *CurriculumService {
	return &CurriculumService{db}
}

// GetCurriculum loads categories + courses + prereqs, then groups by CourseType.
func (s *CurriculumService) GetCurriculum() ([]dto.CurriculumDTO, error) {
	// 1) Load everything
	var cats []model.CurriculumCategory
	if err := s.db.
		Preload("Courses.Prerequisites").
		Find(&cats).Error; err != nil {
		return nil, err
	}

	// 2) Map into DTOs
	var out []dto.CurriculumDTO
	// We're assuming a single “curriculum” — adjust if you have many
	cur := dto.CurriculumDTO{
		ID:         1,
		Name:       "Bsc IT 2021",
		Categories: make([]dto.CategoryDTO, len(cats)),
	}

	for i, c := range cats {
		catDTO := dto.CategoryDTO{
			ID:             int(c.CatID),
			Name:           c.Name,
			CreditRequired: c.CreditReq,
			CourseTypes:    []dto.CourseTypeDTO{},
		}

		// group c.Courses by CourseType
		typeGroup := map[string][]model.Course{}
		for _, course := range c.Courses {
			typeGroup[course.CourseType] = append(typeGroup[course.CourseType], course)
		}

		keys := make([]string, 0, len(typeGroup))
		for courseType := range typeGroup {
			keys = append(keys, courseType)
		}
		sort.Strings(keys)

		for idx, courseType := range keys {
			courses := typeGroup[courseType]
			ctDTO := dto.CourseTypeDTO{
				ID:      idx,        // or derive your own ID
				Name:    courseType, // use course.CourseType directly
				Courses: make([]dto.CourseDTO, len(courses)),
			}
			for j, course := range courses {
				prereqs := make([]string, len(course.Prerequisites))
				for k, pr := range course.Prerequisites {
					prereqs[k] = pr.PreCourseCode
				}
				ctDTO.Courses[j] = dto.CourseDTO{
					Code:          course.CourseCode,
					Name:          course.CourseName,
					Desc:          course.Description,
					Credits:       course.Credits,
					Prerequisites: prereqs,
				}
			}

			catDTO.CourseTypes = append(catDTO.CourseTypes, ctDTO)
		}

		cur.Categories[i] = catDTO
	}
	out = append(out, cur)
	return out, nil
}
