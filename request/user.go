package request

import "github.com/go-playground/validator"

type UserStore struct {
	Name     string `validate:"required" json:"name"`
	Username string `validate:"required" json:"username"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
	RoleId   string `validate:"required" json:"role_id"`
}

type UserUpdate struct {
	Name     string `validate:"required" json:"name"`
	Username string `validate:"required" json:"username"`
	Email    string `validate:"required" json:"email"`
	RoleId   string `validate:"required" json:"role_id"`
}

type UserErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

var userValidate = validator.New()

func UserStoreValidation(request UserStore) []*UserErrorResponse {
	var errors []*UserErrorResponse
	err := userValidate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element UserErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func UserUpdateValidation(request UserUpdate) []*UserErrorResponse {
	var errors []*UserErrorResponse
	err := userValidate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element UserErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
