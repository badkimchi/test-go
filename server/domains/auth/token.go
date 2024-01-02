package auth

type LoginRequest struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
}

type Token struct {
	AuthToken    string       `json:"token"`
	Expiration   string       `json:"expiration"` // expiration is already in the token but useful for frontend
	RefreshToken RefreshToken `json:"refreshToken"`
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"`
	Expiration   string `json:"expiration"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}
