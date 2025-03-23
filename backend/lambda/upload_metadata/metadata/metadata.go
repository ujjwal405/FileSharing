package metadata

type FileInfo struct {
	S3FileName string `json:"s3_filename"`
	Email      string `json:"email"`
	FileName   string `json:"filename"`
	CreatedAt  string `json:"created_at"`
}
