package views

type AplicationUserOutputs struct {
	Id          uint    `json:"id"`
	UserNameApp *string `gorm:"unique;not null" json:"userNameApp"`
	FkCompany   *uint   `json:"fkCompany"`
	Company     *string `json:"company"`
}
