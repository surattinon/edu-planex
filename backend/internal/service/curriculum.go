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

func (s *CurriculumService) GetPersonalCurriculum(userID uint) (*dto.PersonalCurriculumDTO, error) {
	// 1) Load all curriculum categories and their courses
	var cats []model.CurriculumCategory
	if err := s.db.
		Preload("Courses", func(db *gorm.DB) *gorm.DB {
			// only need course_code, course_name, credits, course_type
			return db.Select("course_code", "course_name", "credits", "course_type")
		}).
		Find(&cats).Error; err != nil {
		return nil, err
	}

	// 2) Load all of the user’s enrolled courses (distinct codes)
	var enrolls []model.Enrollment
	if err := s.db.
		Where("user_id = ?", userID).
		Preload("Courses", func(db *gorm.DB) *gorm.DB {
			return db.Select("course_code")
		}).
		Find(&enrolls).Error; err != nil {
		return nil, err
	}
	enrolledSet := make(map[string]struct{})
	for _, e := range enrolls {
		for _, c := range e.Courses {
			enrolledSet[c.CourseCode] = struct{}{}
		}
	}

	// 3) Build the DTO
	out := &dto.PersonalCurriculumDTO{
		UserID:     userID,
		Categories: make([]dto.PCategoryDTO, len(cats)),
	}

	for i, cat := range cats {
		catDTO := dto.PCategoryDTO{
			ID:          int(cat.CatID),
			Name:        cat.Name,
			CourseTypes: nil,
		}

		// group courses by course_type:
		byType := map[string][]model.Course{}
		for _, crs := range cat.Courses {
			byType[crs.CourseType] = append(byType[crs.CourseType], crs)
		}

		// sort keys for stable output:
		keys := make([]string, 0, len(byType))
		for t := range byType {
			keys = append(keys, t)
		}
		sort.Strings(keys)

		// for each course_type build a CourseTypeDTO
		for idx, tname := range keys {
			courses := byType[tname]
			ct := dto.PCourseTypeDTO{
				ID:      idx,
				Name:    tname,
				Courses: make([]dto.PCourseDTO, len(courses)),
			}
			for j, c := range courses {
				_, isEnrolled := enrolledSet[c.CourseCode]
				ct.Courses[j] = dto.PCourseDTO{
					Code:       c.CourseCode,
					Credits:    c.Credits,
					IsEnrolled: isEnrolled,
				}
			}
			catDTO.CourseTypes = append(catDTO.CourseTypes, ct)
		}

		out.Categories[i] = catDTO
	}

	return out, nil
}
