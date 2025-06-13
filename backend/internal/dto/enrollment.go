package dto

// CourseInfo holds only the fields we need in the response.
type CourseInfo struct {
    CourseCode string `json:"course_code"`
    CourseName string `json:"course_name"`
    Credits    int    `json:"credits"`
}

// SemesterInfo is embedded in the enrollment DTO.
type SemesterInfo struct {
    Year   int `json:"year"`
    Number int `json:"number"`
}

// EnrollmentHistoryDTO is the API response for each enrollment.
type EnrollmentHistoryDTO struct {
    EnrollmentID uint          `json:"enrollment_id"`
    UserID       uint          `json:"user_id"`
    Semester     SemesterInfo  `json:"semester"`
    Courses      []CourseInfo  `json:"courses"`
}
