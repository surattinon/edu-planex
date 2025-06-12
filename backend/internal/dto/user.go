package dto

type Profile struct {
	UserID int `json:"user_id"`
	StuID int `json:"student_id"`
	Username string `json:"username"`
	YearStart string `json:"year_start"`
}
