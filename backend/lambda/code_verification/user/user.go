package user

type UserConfirmation struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
