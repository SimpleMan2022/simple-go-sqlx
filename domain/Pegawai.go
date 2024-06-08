package domain

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Pegawai struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"Name"`
	Address string `json:"address" db:"Address"`
	Age     int    `json:"age" db:"Age"`
}

type PegawaiRequest struct {
	Name    string `json:"name" db:"Name" validate:"required"`
	Address string `json:"address" db:"Address" validate:"required"`
	Age     int    `json:"age" db:"Age" validate:"required,number"`
}

func (r *PegawaiRequest) ValidatePegawai() []string {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		var ErrorMessages []string
		for _, fieldError := range err.(validator.ValidationErrors) {
			switch fieldError.Tag() {
			case "required":
				ErrorMessages = append(ErrorMessages,
					fmt.Sprintf("Field %s must be required"), fieldError.Tag())
			case "number":
				ErrorMessages = append(ErrorMessages,
					fmt.Sprintf("Field %s must be a number"), fieldError.Tag())
			}

		}
		if len(ErrorMessages) > 0 {
			return ErrorMessages
		}
	}
	return nil
}
