package dto

type Plan struct {
	PlanID        string   `json:"plan_id" binding:"required"`
}

///////////

type PlanListDTO struct {
    PlanID    uint          `json:"plan_id"`
    Name      string        `json:"name"`
    UserID    uint          `json:"user_id"`
    IsApply   bool          `json:"is_apply"`
    Courses   []CourseInfo  `json:"courses"`
}
