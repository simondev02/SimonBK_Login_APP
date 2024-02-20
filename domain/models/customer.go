package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name                 *string
	Address              *string
	PhoneNumber          *string
	Email                *string
	IdentificationNumber *string
	FkCompany            *uint
	Company              Company `gorm:"foreignKey:FkCompany"`
}
