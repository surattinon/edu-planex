package model

import (
	"time"
)

// Courses Catalog
type Course struct {
	CourseID          string             `gorm:"column:course_id;primaryKey;size:10" json:"course_id"`
	CourseName        string             `gorm:"column:course_name;size:100;not null" json:"course_name"`
	Credits           int                `gorm:"not null" json:"credits"`
	CurriculumCourses []CurriculumCourse `gorm:"foreignKey:CourseID;references:CourseID" json:"curriculum_courses"`
	Sections          []Section          `gorm:"foreignKey:CourseID;references:CourseID" json:"sections"`
}

// Curricula
type Curriculum struct {
	CurriculumID  uint             `gorm:"column:curriculum_id;primaryKey;autoIncrement" json:"curriculum_id"`
	Name          string           `gorm:"size:100;not null" json:"name"`
	Version       string           `gorm:"size:10;not null" json:"version"`
	EffectiveDate time.Time        `gorm:"not null" json:"effective_date"`
	RetiredDate   *time.Time       `json:"retired_date"`
	Categories    []CourseCategory `gorm:"foreignKey:CurriculumID" json:"categories"`
}

// Course Category
type CourseCategory struct {
	CategoryID        uint               `gorm:"column:category_id;primaryKey;autoIncrement" json:"category_id"`
	CurriculumID      uint               `gorm:"not null" json:"curriculum_id"`
	Curriculum        Curriculum         `gorm:"foreignKey:CurriculumID;references:CurriculumID" json:"curriculum"`
	Name              string             `gorm:"size:50;not null" json:"name"`
	CurriculumCourses []CurriculumCourse `gorm:"foreignKey:CategoryID" json:"curriculum_courses"`
}

// Curriculum–Course link
type CurriculumCourse struct {
	ID           uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID   uint           `gorm:"not null" json:"category_id"`
	Category     CourseCategory `gorm:"foreignKey:CategoryID;references:CategoryID" json:"category"`
	CourseID     string         `gorm:"size:10;not null" json:"course_id"`
	Course       Course         `gorm:"foreignKey:CourseID;references:CourseID" json:"course"`
	YearOffered  int            `gorm:"not null" json:"year_offered"`
	Prerequisite *string        `gorm:"size:10" json:"prerequisite"`
}

// Advisors
type Advisor struct {
	AdvisorID string    `gorm:"column:advisor_id;primaryKey" json:"advisor_id"`
	FName     string    `gorm:"column:fname;size:50;not null" json:"fname"`
	LName     string    `gorm:"column:lname;size:50;not null" json:"lname"`
	Password  string    `gorm:"column:pass_hash;size:255;not null" json:"pass_hash"`
	Students  []Student `gorm:"foreignKey:AdvisorID" json:"students"`
}

// Students
type Student struct {
	StudentID  string      `gorm:"column:student_id;primaryKey" json:"student_id"`
	FName      string      `gorm:"column:fname;size:50;not null" json:"fname"`
	LName      string      `gorm:"column:lname;size:50;not null" json:"lname"`
	Password   string      `gorm:"column:pass_hash;size:255;not null" json:"pass_hash"`
	AdvisorID  string      `gorm:"not null" json:"advisor_id"`
	Advisor    Advisor     `gorm:"foreignKey:AdvisorID;references:AdvisorID" json:"advisor"`
	StudyPlans []StudyPlan `gorm:"foreignKey:StudentID" json:"study_plans"`
}

// Terms
type Term struct {
	TermID     uint        `gorm:"column:term_id;primaryKey;autoIncrement" json:"term_id"`
	TermName   string      `gorm:"size:50;not null" json:"term_name"`
	StartDate  time.Time   `gorm:"not null" json:"start_date"`
	EndDate    time.Time   `gorm:"not null" json:"end_date"`
	StudyPlans []StudyPlan `gorm:"foreignKey:TermID" json:"study_plans"`
	Sections   []Section   `gorm:"foreignKey:TermID" json:"sections"`
}

// Study Plans
type StudyPlan struct {
	PlanID      uint         `gorm:"column:plan_id;primaryKey;autoIncrement" json:"plan_id"`
	StudentID   uint         `gorm:"not null" json:"student_id"`
	Student     Student      `gorm:"foreignKey:StudentID;references:StudentID" json:"student"`
	TermID      uint         `gorm:"not null" json:"term_id"`
	Term        Term         `gorm:"foreignKey:TermID;references:TermID" json:"term"`
	Status      string       `gorm:"size:20;not null" json:"status"`
	CreatedAt   time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
	Enrollments []Enrollment `gorm:"foreignKey:PlanID" json:"enrollments"`
}

// Sections
type Section struct {
	SectionID   uint         `gorm:"column:section_id;primaryKey;autoIncrement" json:"section_id"`
	CourseID    string       `gorm:"size:10;not null" json:"course_id"`
	Course      Course       `gorm:"foreignKey:CourseID;references:CourseID" json:"course"`
	TermID      uint         `gorm:"not null" json:"term_id"`
	Term        Term         `gorm:"foreignKey:TermID;references:TermID" json:"term"`
	SectionNo   int          `gorm:"not null" json:"section_no"`
	Schedule    *string      `gorm:"size:100" json:"schedule"`
	Capacity    *int         `json:"capacity"`
	Instructor  *string      `gorm:"size:100" json:"instructor"`
	Enrollments []Enrollment `gorm:"foreignKey:SectionID" json:"enrollments"`
}

// Enrollments
type Enrollment struct {
	EnrollID  uint      `gorm:"column:enroll_id;primaryKey;autoIncrement" json:"enroll_id"`
	PlanID    uint      `gorm:"not null" json:"plan_id"`
	StudyPlan StudyPlan `gorm:"foreignKey:PlanID;references:PlanID" json:"study_plan"`
	SectionID uint      `gorm:"not null" json:"section_id"`
	Section   Section   `gorm:"foreignKey:SectionID;references:SectionID" json:"section"`
	Status    *string   `gorm:"size:20" json:"status"`
	Grade     *string   `gorm:"size:5" json:"grade"`
}
