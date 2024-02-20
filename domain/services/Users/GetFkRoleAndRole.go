package services

import "gorm.io/gorm"

type UserRole struct {
	FkRole *uint   `json:"fkRole"`
	Role   *string `json:"role"`
}

func GetUserRoleByID(db *gorm.DB, id uint) (*UserRole, error) {
	var userRole UserRole
	err := db.Raw(`SELECT roles.id AS FkRole, roles.role_description AS Role 
	FROM users 
	INNER JOIN user_roles ON user_roles.fk_user = users.id 
	INNER JOIN roles ON roles.id = user_roles.fk_role 
	WHERE users.id = ? LIMIT 1`, id).Scan(&userRole).Error
	if err != nil {
		return nil, err
	}
	return &userRole, nil
}
