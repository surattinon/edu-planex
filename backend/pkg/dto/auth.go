package dto

// SignupRequest is the payload for /auth/signup
type SignupRequest struct {
	FName     string `json:"fname" validate:"required"`
	LName     string `json:"lname" validate:"required"`
	StudentID string `json:"student_id" validate:"required"`
	Pass      string `json:"password" validate:"required,min=6"`
	AdvisorID string `json:"advisor_id" validate:"required"`
}

type StudentSignupRequest struct {
	FName     string `json:"fname" validate:"required"`
	LName     string `json:"lname" validate:"required"`
	StudentID string `json:"student_id" validate:"required"`
	Pass      string `json:"password" validate:"required,min=6"`
	AdvisorID string `json:"advisor_id" validate:"required"`
}

type AdvisorSignupRequest struct {
	AdvisorID string `json:"advisor_id" validate:"required"`
	FName     string `json:"fname" validate:"required"`
	LName     string `json:"lname" validate:"required"`
	Pass      string `json:"password" validate:"required,min=6"`
}

// LoginRequest is the payload for /auth/login
type LoginRequest struct {
	ID string `json:"id" validate:"required"`
	Pass      string `json:"password" validate:"required,min=6"`
}

// TokenResponse returns a JWT
type TokenResponse struct {
	Token string `json:"token"`
}
