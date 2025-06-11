package dto

type CourseTable struct {
	CourseCode    string   `json:"course_code"`
	CourseName    string   `json:"course_name"`
	Credits       int      `json:"credits"`
	Prerequisites []string `json:"prerequisites"`
}

type PlanTables struct {
	PlanID   uint          `json:"plan_id"`
	PlanName string        `json:"plan_name"`
	UserID   uint          `json:"user_id"`
	Courses  []CourseTable `json:"courses"`
}
