package model

// NewUser is used to store values passed by the user while sign up
type NewUser struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Bio   *string `json:"bio"`
}
