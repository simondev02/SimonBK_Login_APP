package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name    *string
	Address *string
	Phone   *string
	Email   *string
}
