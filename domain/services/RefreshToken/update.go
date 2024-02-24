package services

import (
	"SimonBK_Login_APP/domain/models"
	"SimonBK_Login_APP/infra/db"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

func UpdateRefreshToken(userId uint) (*string, error) {

	// Genera un nuevo token
	newToken := uuid.New().String()
	ExpiryDate := time.Now().Add(time.Hour * 24 * 7)

	result := db.DBConn.Model(&models.RefreshToken{}).Where("fk_aplication_user = ?", userId).Update("token", newToken).Update("expiry_date", ExpiryDate)

	if result.Error != nil {
		log.Print("[UpdateRefresToken] - Actualizacion erronea")
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		log.Print("[UpdateRefreshToken] - Ningún token de refresco fue actualizado")
		return nil, errors.New("ningún token de refresco fue actualizado")
	}

	return &newToken, nil
}
