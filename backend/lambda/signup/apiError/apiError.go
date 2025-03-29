package apiError

import "fmt"

type APIError struct {
	StatusCode int `json:"statusCode"`
	Msg        any `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error %v", e.Msg)
}

func UserAlreadyExistsError() APIError {
	return APIError{
		StatusCode: 400,
		Msg:        "User already exists with this email",
	}
}

func MinimumLength()APIError{
	return APIError{
		StatusCode: 400,
		Msg: "password must be atleast of length 8",
	}
}