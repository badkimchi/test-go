package auth

type LoginRequest struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
}

type OAuthRequest struct {
	Token string `json:"token"`
}

type Token struct {
	Token        string       `json:"token"`
	Expiration   string       `json:"expiration"` // expiration is already in the token but useful for frontend
	RefreshToken RefreshToken `json:"refreshToken"`
}

type RefreshToken struct {
	Token      string `json:"token"`
	Expiration string `json:"expiration"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}
