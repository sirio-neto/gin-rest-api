package studentmodel

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]{11}$"`
	RG   string `json:"rg" validate:"len=9, regexp=^[0-9]{9}$"`
}

func (student *Student) Validate() error {
	if err := validator.Validate(student); err != nil {
		return err
	}

	return nil
}

var Students []Student
