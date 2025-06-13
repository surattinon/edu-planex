package dto
type ProgressItem struct {
	CategoryKey string `json:"key"`
	Required    int    `json:"required"`
	Earned      int    `json:"earned"`
}

type ProgressResponse struct {
	UserID  uint           `json:"user_id"`
	Courses []ProgressItem `json:"courses"`
}
