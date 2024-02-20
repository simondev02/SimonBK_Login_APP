package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email          *string    `gorm:"column:email"`
	Password       *string    `gorm:"column:password"`
	FkCompany      *uint      `gorm:"column:fk_company"`
	FkCustomer     *uint      `gorm:"column:fk_customer"`
	Company        Company    `gorm:"foreignKey:FkCompany"`
	Customer       Customer   `gorm:"foreignKey:FkCustomer"`
	Contacts       []Contact  `gorm:"many2many:user_contacts;foreignKey:ID;joinForeignKey:FkUser;References:ID;joinReferences:FkContact"`
	UserRoles      []UserRole `gorm:"foreignKey:FkUser"` // Relación con UserRole
	DeletedBy      *uint      `gorm:"column:deleted_by"`
	UpdatedBy      *uint      `gorm:"column:updated_by"`
	Last_login     *time.Time `gorm:"column:last_login"`
	Login_attempts *uint      `gorm:"column:login_attempts"`
	State          *uint      `gorm:"column:state"`
}
