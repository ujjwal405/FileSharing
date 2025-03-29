package helper

import (
	"errors"

	"github.com/go-passwd/validator"
)



var(
	errMinLength= errors.New("password must be atleast of length 8")
)




type Validator struct{
passwordValidator *validator.Validator
}


func NewValidator()*Validator{
	passwordValidator := validator.New(validator.MinLength(8,errMinLength))
	return &Validator{
		passwordValidator: passwordValidator,
	}
}


func (v *Validator) ValidatePassword(password string)error{
	if err := v.passwordValidator.Validate(password);err!=nil{
		return err
	}
	return nil

}
