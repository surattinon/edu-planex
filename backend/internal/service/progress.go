package service

import (
	"github.com/surattinon/edu-planex/backend/internal/dto"
	"gorm.io/gorm"
)

type ProgressService struct {
	db *gorm.DB
}

func NewProgressService(db *gorm.DB) *ProgressService {
	return &ProgressService{db: db}
}

// GetProgress fetches required vs earned credits per category for userID.
func (s *ProgressService) GetProgress(userID uint) (*dto.ProgressResponse, error) {
	const sql = `
WITH user_courses AS (
  SELECT DISTINCT ec.course_code
  FROM enrollments e
  JOIN enrollment_courses ec
    ON e.enrollment_id = ec.enrollment_id
  WHERE e.user_id = ?
)
SELECT
  ccat.cat_id,
  ccat.name,
  ccat.credit_req,
  COALESCE(SUM(c.credits), 0) AS earned
FROM curriculum_categories AS ccat
LEFT JOIN curriculum_courses AS cc
  ON ccat.cat_id = cc.curriculum_cat_id
LEFT JOIN user_courses AS uc
  ON cc.course_code = uc.course_code
LEFT JOIN courses AS c
  ON c.course_code = uc.course_code
GROUP BY
  ccat.cat_id, ccat.name, ccat.credit_req
ORDER BY
  ccat.cat_id;
`
	rows, err := s.db.Raw(sql, userID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resp := &dto.ProgressResponse{UserID: userID}
	for rows.Next() {
		var catID int
		var name string
		var required, earned int

		if err := rows.Scan(&catID, &name, &required, &earned); err != nil {
			return nil, err
		}

		// Map the human name to your JSON key
		key := ""
		switch name {
		case "General Education Courses":
			key = "general_education"
		case "Professional Courses":
			key = "professional"
		case "Free Electives":
			key = "free_elective"
		case "Internship":
			key = "internship"
		default:
			key = name
		}

		resp.Courses = append(resp.Courses, dto.ProgressItem{
			CategoryKey: key,
			Required:    required,
			Earned:      earned,
		})
	}

	return resp, nil
}
