package dto

type PCourseDTO struct {
	Code       string `json:"code"`
	Credits    int    `json:"credits"`
	IsEnrolled bool   `json:"is_enrolled"`
}

type PCourseTypeDTO struct {
	ID      int          `json:"id"`
	Name    string       `json:"name"`
	Courses []PCourseDTO `json:"courses"`
}

type PCategoryDTO struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	CourseTypes []PCourseTypeDTO `json:"course_types"`
}

type PersonalCurriculumDTO struct {
	UserID     uint           `json:"user_id"`
	Categories []PCategoryDTO `json:"categories"`
}
