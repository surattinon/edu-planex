package dto

type CatCredits struct {
	GeneralEd	int `json:"general_education"`
	Professional int `json:"professional"`
	FreeElec	int `json:"free_elective"`
	Intern int `json:"internship"`
}

type CreditResult struct {
	UserID	int `json:"user_id"`
	Credits []CatCredits `json:"credits"`
}
