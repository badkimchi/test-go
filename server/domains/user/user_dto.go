package user

type User struct {
	UserID         string `gorm:"unique;not null;type:varchar(100);"` // set member number to unique and not null
	PWD            string
	PrivilegeTitle string `example:"monitor, manager, admin"`
	Email          string
}

type UserChangePasswordDto struct {
	OldPassword string
	NewPassword string
}

type UserArrayDto struct {
	Users []User
}

type UserCredentials struct {
	UserID       string `json:"userID" example:"admin"`
	PWD          string `json:"pwd" example:"admin"`
	Email        string `json:"email"`
	PreAuthToken string `json:"preAuthToken"`
	Otp          string `json:"otp"`
}
