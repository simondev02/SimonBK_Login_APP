package services

import (
	"SimonBK_Login_APP/api/views"
	Token "SimonBK_Login_APP/domain/services/AccesToken"
	Login "SimonBK_Login_APP/domain/services/Login"
	RefreshToken "SimonBK_Login_APP/domain/services/RefreshToken"
	UserApp "SimonBK_Login_APP/domain/services/UserApp"
	Users "SimonBK_Login_APP/domain/services/Users"

	"gorm.io/gorm"
)

func LoginUserApp(db *gorm.DB, userNameApp *string, username *string, password string) (*views.Response, error) {

	// 1. Validamos que el usuario de aplicación exista y retorna password cifrado.
	hashPassword, idUser, FkCompanyApp, err := UserApp.GetUserApp(db, userNameApp)
	if err != nil {
		return nil, err
	}

	// 2. Comparamos la contraseña ingresada con la contraseña cifrada almacenada.
	match, err := Login.ComparePasswords(hashPassword, password)
	if err != nil {
		return nil, err
	}

	// 3. Obtenemos los datos del usuario por su email.
	user, err := Users.GetUserByEmail(db, username)
	if err != nil {
		return nil, err
	}
	// 4. Validamos jerarquia entre UserApp y useName
	err = Login.ValidateHierarchyCompany(user.FkCompany, FkCompanyApp)
	if err != nil {
		return nil, err
	}

	// 5. Generamos el token de acceso.
	token, err := Token.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	// 6. Generamos el tokenrefresh
	refreshToken, err := RefreshToken.GenerateRefreshToken(idUser)
	if err != nil {
		return nil, err
	}

	response := views.Response{
		Success:        match,
		FailedAttempts: nil,
		AccessToken:    token,
		RefreshToken:   refreshToken,
		Message:        "Inicio de sesión exitoso",
		Users:          user,
	}

	return &response, err
}
