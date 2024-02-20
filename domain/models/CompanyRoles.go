package models

import "gorm.io/gorm"

type CompanyRole struct {
	gorm.Model
	FkCompany *uint `gorm:"column:fk_company"`
	FkRole    *uint `gorm:"column:fk_role"` // Relación hacia tabla roles
	Role      Role  `gorm:"foreignKey:FkRole"`
}
