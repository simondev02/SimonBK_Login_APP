package models

import "gorm.io/gorm"

type AplicationUser struct {
	gorm.Model
	UserNameApp *string `gorm:"unique;not null" json:"userNameApp"`
	Password    string  `json:"password"`
	FkCompany   *uint   `json:"fkCompany"`
	FkCustomer  *uint
	DeletedBy   *uint
	UpdatedBy   *uint
	CreatedBy   *uint
	Comnpay     Company `gorm:"foreignKey:FkCompany"`
}
