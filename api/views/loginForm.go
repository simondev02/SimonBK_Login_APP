package views

type LoginForm struct {
	Email       *string `json:"email"`
	UserNameApp *string `json:"userNameApp"`
	Password    *string `json:"password"`
}
