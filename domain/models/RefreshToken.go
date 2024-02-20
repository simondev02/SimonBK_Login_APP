package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	Token            *string
	FkUser           *uint
	FkAplicationUser *uint
	ExpiryDate       *time.Time
	DeletedByUserID  *uint
	UpdatedByUserID  *uint
	User             Users          `gorm:"foreignKey:FkUser"`
	AplicationUser   AplicationUser `gorm:"foreignKey:FkAplicationUser"`
}
