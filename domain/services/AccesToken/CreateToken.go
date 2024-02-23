package services

import (
	"SimonBK_Login_APP/api/views"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateAccessToken(user *views.Users) (*string, error) {
	jwtKey := os.Getenv("JWT_KEY")
	expirationTime := time.Now().Add(24 * time.Hour)
	idStr := strconv.FormatUint(uint64(user.ID), 10)

	claims := views.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   idStr,
			ExpiresAt: expirationTime.Unix(),
		},
		FkCompany:  user.FkCompany,
		FkCustomer: user.FkCustomer,
		UserId:     &user.ID,
		RoleId:     user.FkRole,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
