package models

import "gorm.io/gorm"

type AplicationUser struct {
	gorm.Model
	UserNameApp *string
	Password    string
	FkCompany   *uint
	FkCustomer  *uint
	DeletedBy   *uint
	UpdatedBy   *uint
	CreatedBy   *uint
	Company     Company `gorm:"foreignKey:FkCompany"`
}
