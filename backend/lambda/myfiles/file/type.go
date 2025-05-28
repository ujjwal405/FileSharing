package file

type FileItem struct {
	S3Filename string `json:"s3filename"`
	Filename   string `json:"filename"`
	CreatedAt  string `json:"created_at"`
}

// PaginatedFiles represents the response with pagination info
type PaginatedFiles struct {
	Items      []FileItem `json:"items"`
	TotalPages int        `json:"total_pages"`
}

type Token struct {
	AccessToken *string `json:"access_token,omitempty"`
	IDToken     *string `json:"id_token,omitempty"`
}

type Response struct {
	Files     PaginatedFiles `json:"files"`
	UserToken *Token         `json:"user_token,omitempty"`
}
