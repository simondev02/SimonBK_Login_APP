package services

import (
	"SimonBK_Login_APP/domain/models"
	"SimonBK_Login_APP/infra/db"
)

func CreateLogin(userId *uint, ip *string, agent *string, success *bool) error {
	var Login models.Login

	Login.FkAplicationUser = userId
	Login.IpAddress = ip
	Login.UserAgent = agent
	Login.Success = success

	err := db.DBConn.Create(&Login)
	if err.Error != nil {
		return err.Error
	}

	return nil

}
