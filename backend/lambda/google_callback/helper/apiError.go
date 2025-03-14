package helper

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

func UnexpectedSigningMethod() APIError {
	return APIError{
		StatusCode: 500,
		Msg:        "Unexpected signing method",
	}
}

func InvalidStateToken() APIError {
	return APIError{
		StatusCode: 400,
		Msg:        "Invalid state token",
	}
}

func TokenExpired() APIError {
	return APIError{
		StatusCode: 400,
		Msg:        "Token expired",
	}
}

func InvalidExpirationTime() APIError {
	return APIError{
		StatusCode: 500,
		Msg:        "Invalid expiration time",
	}
}
