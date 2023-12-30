package auth

type Token struct {
	AuthToken    string       `json:"token"`
	Type         string       `json:"type"`
	Expiration   string       `json:"expiration"`
	RefreshToken RefreshToken `json:"refreshToken"`
}

type TokenWithPrivilegeTitleExposed struct {
	Token
	// Only purpose is to let the frontend code know the privilege.
	// Not a security vulnerability since the info is also embedded in the token itself.
	PrivilegeTitle string
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"`
	Expiration   string `json:"expiration"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}
