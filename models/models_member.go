package models

type Member struct {
	Email        string `json:"email" validate:"required,email"`
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	LineID       string `json:"lineId"`
	Tel          string `json:"tel" validate:"required"`
	BusinessType string `json:"businessType" validate:"required"`
	WebsiteName  string `json:"websiteName" validate:"required"`
}

// func validUsername(fl validator.FieldLevel) bool {
// 	username := fl.Field().String()
// 	match, _ := regexp.MatchString("^[a-zA-Z0-9_.]+$", username)

// 	return match
// }
