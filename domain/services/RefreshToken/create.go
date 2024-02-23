package services

import (
	"SimonBK_Login_APP/domain/models"
	"SimonBK_Login_APP/infra/db"
	"time"

	"github.com/google/uuid"
)

func CreateRefreshToken(userId uint) (*string, error) {

	refreshToken := models.RefreshToken{
		FkAplicationUser: userId,
		Token:            uuid.New().String(),
		ExpiryDate:       time.Now().Add(time.Hour * 24 * 7),
	}

	err := db.DBConn.Create(&refreshToken).Error
	if err != nil {
		return nil, err
	}

	return &refreshToken.Token, nil
}
