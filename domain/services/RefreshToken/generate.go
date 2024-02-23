package services

func GenerateRefreshToken(userId uint) (*string, error) {
	var Token *string
	_, err := ValidateRefreshToken(userId)
	if err != nil {
		if err.Error() == "refresh token no encontrado" {
			// Si el token no existe, crea uno nuevo
			Token, err = CreateRefreshToken(userId)
			if err != nil {
				return nil, err
			}
		} else {
			// Si hay un error diferente al no encontrar el token
			return nil, err
		}
	} else {
		// Si el token existe, actualízalo
		Token, err = UpdateRefreshToken(userId)
		if err != nil {
			return nil, err
		}
	}
	return Token, nil
}
