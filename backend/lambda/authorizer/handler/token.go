package handler

type NewToken struct {
	AccessToken string `json:"access_token"`
	IdToken     string `json:"id_token"`
}
