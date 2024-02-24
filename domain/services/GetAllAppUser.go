package services

import (
	views "SimonBK_Login_APP/api/views/Outputs"
	"SimonBK_Login_APP/domain/models"

	"gorm.io/gorm"
)

func GetAllAppUser(db *gorm.DB) ([]views.AplicationUserOutputs, error) {
	var UserApp []models.AplicationUser
	query := db.
		Preload("Company").
		Find(&UserApp)

	if query.Error != nil {
		return nil, query.Error
	}

	var responses []views.AplicationUserOutputs
	for _, userApp := range UserApp {
		response := views.AplicationUserOutputs{
			Id:          userApp.ID,
			UserNameApp: userApp.UserNameApp,
			FkCompany:   userApp.FkCompany,
			Company:     userApp.Company.Name,
		}
		responses = append(responses, response)
	}

	return responses, nil
}
