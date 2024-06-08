package domain

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id       string `json:"id" db:"id"`
	Name     string `json:"name" db:"Name"`
	Email    string `json:"email" db:"Email	"`
	Password string `json:"password" db:"Password"`
}

type UserResponse struct {
	Name  string `json:"name" db:"Name"`
	Email string `json:"email" db:"Email"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginRequest) ValidateUser() []string {
	validate := validator.New()
	err := validate.Struct(r)

	if err != nil {
		var validationErrorMessages []string
		for _, fieldError := range err.(validator.ValidationErrors) {
			switch fieldError.Tag() {
			case "required":
				validationErrorMessages = append(validationErrorMessages,
					fmt.Sprintf("Field %s must be required", fieldError.Field()))
			case "email":
				validationErrorMessages = append(validationErrorMessages,
					fmt.Sprintf("Field %s must be valid email", fieldError.Field()))
			}
		}
		if len(validationErrorMessages) > 0 {
			return validationErrorMessages
		}
	}
	return nil
}
