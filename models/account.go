package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AgencyNumber  string `json:"AgencyNumber" validate:"nonzero,regexp=^[0-9]*$"`
	AccountNumber string `json:"AccountNumber" validate:"nonzero,regexp=^[0-9]*$"`
}

func Validate(account *Account) error {
	if err := validator.Validate(account); err != nil {
		return err
	}

	return nil
}
