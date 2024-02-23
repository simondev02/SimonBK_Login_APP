package views

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	jwt.StandardClaims
	UserId     *uint `json:"userId"`
	RoleId     *uint `json:"roleId"`
	FkCompany  *uint `json:"fk_company"`
	FkCustomer *uint `json:"fk_customer"`
}
