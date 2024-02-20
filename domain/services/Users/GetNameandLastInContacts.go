package services

import "gorm.io/gorm"

type ContactName struct {
	Name     *string `json:"name"`
	LastName *string `json:"lastname"`
}

func GetContactNameByID(db *gorm.DB, id uint) (*ContactName, error) {
	var contactName ContactName
	err := db.Raw(`SELECT contacts.name AS Name, contacts.lastname AS LastName 
	FROM users 
	INNER JOIN user_contacts ON user_contacts.fk_user = users.id 
	INNER JOIN contacts ON contacts.id = user_contacts.fk_contact 
	WHERE users.id = ? LIMIT 1`, id).Scan(&contactName).Error
	if err != nil {
		return nil, err
	}
	return &contactName, nil
}
