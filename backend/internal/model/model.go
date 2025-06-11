package model

import "time"

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

type UserProfile struct {
	UserID      uint      `gorm:"column:user_id;primaryKey" json:"user_id"`
	DisplayName string    `gorm:"column:display_name;not null" json:"display_name"`
	Email       string    `gorm:"column:email" json:"email"`
	AvatarURL   string    `gorm:"column:avatar_url" json:"avatar_url"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

type Plan struct {
	PlanID    uint      `gorm:"column:plan_id;primaryKey" json:"plan_id"`
	UserID    uint      `gorm:"column:user_id;not null" json:"user_id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`

	// preloadable relations
	User    UserProfile `gorm:"foreignKey:UserID;references:UserID" json:"user"`
	Courses []Course    `gorm:"many2many:plan_courses;joinForeignKey:PlanID;joinReferences:CourseCode" json:"courses"`
}

type PlanCourse struct {
	PlanID     uint   `gorm:"column:plan_id;primaryKey" json:"plan_id"`
	CourseCode string `gorm:"column:course_code;primaryKey" json:"course_code"`
}

type Semester struct {
	SemesterID uint `gorm:"column:semester_id;primaryKey" json:"semester_id"`
	Year       int  `gorm:"column:year" json:"year"`
	SemesterNo int  `gorm:"column:semester_number" json:"semester_number"`
}

type Enrollment struct {
	EnrollmentID uint      `gorm:"column:enrollment_id;primaryKey" json:"enrollment_id"`
	UserID       uint      `gorm:"column:user_id" json:"user_id"`
	SemesterID   uint      `gorm:"column:semester_id" json:"semester_id"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`

	// preloadable relations
	Semester Semester `gorm:"foreignKey:SemesterID;references:SemesterID" json:"semester"`
	Courses  []Course `gorm:"many2many:enrollment_courses;joinForeignKey:EnrollmentID;joinReferences:CourseCode" json:"courses"`
}

type EnrollmentCourse struct {
	EnrollmentID uint   `gorm:"column:enrollment_id;primaryKey"`
	CourseCode   string `gorm:"column:course_code;primaryKey"`
}

// Force GORM to use these table name
func (CurriculumCategory) TableName() string { return "curriculum_categories" }
func (Course) TableName() string             { return "courses" }
func (Prerequisite) TableName() string       { return "prerequisites" }
func (UserProfile) TableName() string        { return "user_profile" }
func (Plan) TableName() string               { return "plans" }
func (PlanCourse) TableName() string         { return "plan_courses" }
func (Semester) TableName() string           { return "semesters" }
func (Enrollment) TableName() string         { return "enrollments" }
func (EnrollmentCourse) TableName() string   { return "enrollment_courses" }
