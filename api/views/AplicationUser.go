package views

type AplicationUser struct {
	Id          *uint
	UserNameApp *string `gorm:"unique;not null" json:"userNameApp"`
	Password    string  `json:"password"`
	FkCompany   *uint   `json:"fkCompany"`
	FkCustomer  *uint
	UpdatedBy   *uint
	CreatedBy   *uint
}
