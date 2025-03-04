package confirmation

type Confirmation struct {
	Email            string `json:"email"`
	ConfirmationCode string `json:"confirmation_code"`
	Password         string `json:"password"`
}
