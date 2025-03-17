package apierror

import "errors"

var (
	ErrTokenExpired = errors.New("token expired")
)
