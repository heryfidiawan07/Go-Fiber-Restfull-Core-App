package request

import "github.com/go-playground/validator"

type RoleStore struct {
	Name        string   `validate:"required" json:"name"`
	Permissions []string `validate:"required" json:"permissions"`
}

type RoleUpdate struct {
	Name        string   `validate:"required" json:"name"`
	Permissions []string `validate:"required" json:"permissions"`
}

type RoleErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

var roleValidate = validator.New()

func RoleStoreValidation(request RoleStore) []*RoleErrorResponse {
	var errors []*RoleErrorResponse
	err := roleValidate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element RoleErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func RoleUpdateValidation(request RoleUpdate) []*RoleErrorResponse {
	var errors []*RoleErrorResponse
	err := roleValidate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element RoleErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
