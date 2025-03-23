package file

type FileItem struct {
	S3Filename string
	Filename   string
	CreatedAt  string
}

// PaginatedFiles represents the response with pagination info
type PaginatedFiles struct {
	Items      []FileItem
	TotalPages int
}

type Token struct {
	AccessToken *string `json:"access_token,omitempty"`
	IDToken     *string `json:"id_token,omitempty"`
}

type Response struct {
	Files     PaginatedFiles `json:"files"`
	UserToken *Token         `json:"user_token,omitempty"`
}
