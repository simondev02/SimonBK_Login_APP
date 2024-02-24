package services

import (
	"SimonBK_Login_APP/domain/models"
	"SimonBK_Login_APP/infra/db"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

func ValidateRefreshToken(param interface{}) (*uint, error) {
	var RefreshToken models.RefreshToken
	var Query *gorm.DB

	switch v := param.(type) {
	case string:
		Query = db.DBConn.Where("token = ?", v).First(&RefreshToken)
	case uint:
		Query = db.DBConn.Where("fk_aplication_user = ?", v).First(&RefreshToken)
	default:
		return nil, errors.New("tipo de dato no válido")
	}
	if Query.Error != nil {
		if errors.Is(Query.Error, gorm.ErrRecordNotFound) {
			log.Print("[ValidateRefreshToken] - Refresh token no encontrado")
			return nil, errors.New("refresh token no encontrado")
		}
		return nil, Query.Error
	}

	curentTime := time.Now()
	if RefreshToken.ExpiryDate.Before(curentTime) {
		log.Print("[ValidaterefreshToken] - Token expirado")
		return nil, errors.New("token expirado")
	}
	return RefreshToken.FkUser, nil
}
