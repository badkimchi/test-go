package account

type Account struct {
	UserID         string `gorm:"unique;not null;type:varchar(100);"` // set member number to unique and not null
	PWD            string
	PrivilegeTitle string `example:"monitor, manager, admin"`
	Email          string
}

type AccountChangePasswordDto struct {
	OldPassword string
	NewPassword string
}

type AccountArrayDto struct {
	Accounts []Account
}

type AccountCredentials struct {
	UserID       string `json:"userID" example:"admin"`
	PWD          string `json:"pwd" example:"admin"`
	Email        string `json:"email"`
	PreAuthToken string `json:"preAuthToken"`
	Otp          string `json:"otp"`
}
