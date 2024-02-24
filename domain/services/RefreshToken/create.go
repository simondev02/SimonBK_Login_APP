package services

import (
	"SimonBK_Login_APP/domain/models"
	"SimonBK_Login_APP/infra/db"
	"errors"
	"log"
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
		log.Print("[CreaterefreshToken] - error al crear refreshtoken")
		return nil, errors.New("error al crear refreshtoken")
	}

	return &refreshToken.Token, nil
}
