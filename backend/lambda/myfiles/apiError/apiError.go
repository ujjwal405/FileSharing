package apiError

import "fmt"

type APIError struct {
	StatusCode int `json:"statusCode"`
	Msg        any `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error %v", e.Msg)
}

func FileNotFoundError() APIError {
	return APIError{
		StatusCode: 404,
		Msg:        "no file exists",
	}
}
