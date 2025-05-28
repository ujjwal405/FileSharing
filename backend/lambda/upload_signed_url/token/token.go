package token

type Token struct {
	AccessToken *string `json:"access_token,omitempty"`
	IDToken     *string `json:"id_token,omitempty"`
}

type UploadSignedURLResponse struct {
	URL       string `json:"url"`
	FileID    string `json:"file_id"`
	UserToken *Token `json:"user_token,omitempty"`
}
