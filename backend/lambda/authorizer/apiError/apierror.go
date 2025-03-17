package apierror

import "fmt"

type APIError struct {
	StatusCode int `json:"statusCode"`
	Msg        any `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error %v", e.Msg)
}

func NewAPIError(statusCode int, msg any) APIError {
	return APIError{
		StatusCode: statusCode,
		Msg:        msg,
	}
}
func InvalidSignRequest() APIError {
	return APIError{
		StatusCode: 400,
		Msg:        "invalid signing request",
	}
}

func TokenExpired() APIError {
	return APIError{
		StatusCode: 400,
		Msg:        "token has expired",
	}
}

func InvalidToken() APIError {
	return APIError{
		StatusCode: 400,
		Msg:        "invalid  token",
	}
}
