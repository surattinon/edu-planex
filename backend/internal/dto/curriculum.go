package dto

type CourseDTO struct {
	Code          string   `json:"code"`
	Name          string   `json:"name"`
	Desc          string   `json:"desc"`
	Credits       int      `json:"credits"`
	Prerequisites []string `json:"prerequisites"`
}

type CourseTypeDTO struct {
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	Courses []CourseDTO `json:"courses"`
}

type CategoryDTO struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	CreditRequired int             `json:"credit_required"`
	CourseTypes    []CourseTypeDTO `json:"course_types"`
}

type CurriculumDTO struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Categories []CategoryDTO `json:"categories"`
}
