package dto

type Plan struct {
	PlanID        string   `json:"plan_id" binding:"required"`
}
