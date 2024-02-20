package views

import "time"

type Users struct {
	ID             uint       `json:"id"`
	FkCompany      *uint      `json:"fkCompany"`
	FkCustomer     *uint      `json:"fkCustomer"`
	Email          *string    `json:"email"`
	Name           *string    `json:"name"`
	Lastname       *string    `json:"lastname"`
	FkRole         *uint      `json:"fkRole"`
	Role           *string    `json:"role"`
	Login_attempts *uint      `json:"-"`
	Last_login     *time.Time `json:"lastLogin"`
}
