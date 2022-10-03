package request

import "github.com/go-playground/validator"

type Login struct {
	Identity string `validate:"required" json:"identity"`
	Password string `validate:"required" json:"password"`
}

type Register struct {
	Name     string `validate:"required" json:"name"`
	Username string `validate:"required" json:"username"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type RefreshToken struct {
	RefreshToken string `validate:"required" json:"refresh_token"`
}

type AuthErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

var authValidate = validator.New()

func LoginValidation(request Login) []*AuthErrorResponse {
	var errors []*AuthErrorResponse
	err := authValidate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element AuthErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func RegisterValidation(request Register) []*AuthErrorResponse {
	var errors []*AuthErrorResponse
	err := authValidate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element AuthErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
