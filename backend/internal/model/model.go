package model

type CurriculumCategory struct {
	CatID     uint     `gorm:"column:cat_id;primaryKey" json:"cat_id"`
	Name      string   `gorm:"column:name" json:"name"`
	CreditReq int      `gorm:"column:credit_req" json:"credit_required"`
	Courses   []Course `gorm:"many2many:curriculum_courses;joinForeignKey:CurriculumCatID;joinReferences:CourseCode" json:"-"`
}

type Course struct {
	CourseCode    string               `gorm:"column:course_code;primaryKey" json:"course_code"`
	CourseName    string               `gorm:"column:course_name" json:"course_name"`
	Credits       int                  `gorm:"column:credits" json:"credits"`
	YearOffered   int                  `gorm:"column:year_offered" json:"year_offered"`
	CourseType    string               `gorm:"column:course_type" json:"course_type"`
	Description   string               `gorm:"column:description" json:"description"`
	Categories    []CurriculumCategory `gorm:"many2many:curriculum_courses;joinForeignKey:CourseCode;joinReferences:CurriculumCatID" json:"categories"`
	Prerequisites []Prerequisite       `gorm:"foreignKey:CourseCode;references:CourseCode" json:"prerequisites"`
}

type Prerequisite struct {
	CourseCode    string `gorm:"column:course_code;primaryKey"`
	PreCourseCode string `gorm:"column:pre_course_code;primaryKey" json:"pre_course_code"`
}

func (CurriculumCategory) TableName() string {
	return "curriculum_categories"
}

func (Course) TableName() string {
	return "courses"
}

func (Prerequisite) TableName() string {
	return "prerequisites"
}
