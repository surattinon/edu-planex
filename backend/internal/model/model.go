package model

type CurriculumCategories struct {
	ID        uint     `gorm:"primaryKey;autoIncrement"`          // maps to SERIAL PRIMARY KEY
	Name      string   `gorm:"type:varchar(100);not null;unique"` // maps to VARCHAR(100) NOT NULL UNIQUE
	CreditReq int      `gorm:"not null"`                          // maps to INT NOT NULL
	Courses   []Courses `gorm:"many2many:curriculum_courses;constraint:OnDelete:CASCADE"`
}

type Courses struct {
	Code              string               `gorm:"primaryKey;type:varchar(20)"` // maps to VARCHAR(20) PRIMARY KEY :contentReference[oaicite:2]{index=2}
	Name              string               `gorm:"type:varchar(255);not null"`  // maps to VARCHAR(255) NOT NULL
	Credits           int                  `gorm:"not null"`                    // maps to INT NOT NULL
	YearOffered       int                  `gorm:"not null"`                    // maps to INT NOT NULL
	CourseType        string               `gorm:"type:varchar(100);not null"`  // maps to VARCHAR(100) NOT NULL
	Description       string               `gorm:"type:text"`                   // maps to TEXT (nullable)
	Categories        []CurriculumCategories `gorm:"many2many:curriculum_courses;constraint:OnDelete:CASCADE"`
	Prerequisites     []Courses             `gorm:"many2many:prerequisites;joinForeignKey:CourseCode;joinReferences:PreCourseCode;constraint:OnDelete:CASCADE"`
	IsPrerequisiteFor []Courses             `gorm:"many2many:prerequisites;joinForeignKey:PreCourseCode;joinReferences:CourseCode;constraint:OnDelete:CASCADE"`
}

type CurriculumCourses struct {
	CurriculumCatID uint   `gorm:"column:curriculum_cat_id;primaryKey;"`           // fk → CurriculumCategory(ID) :contentReference[oaicite:5]{index=5}
	CourseCode      string `gorm:"column:course_code;type:varchar(20);primaryKey"` // fk → Course(Code)
}

type Prerequisites struct {
	CourseCode    string `gorm:"column:course_code;type:varchar(20);primaryKey"`     // fk → Course(Code)
	PreCourseCode string `gorm:"column:pre_course_code;type:varchar(20);primaryKey"` // fk → Course(Code)
}
