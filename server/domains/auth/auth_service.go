package auth

import (
	"app/domains/account"
	"github.com/go-chi/jwtauth"
	"strconv"
	"time"
)

type IAuthService interface {
	setAuthTokenDuration(duration time.Duration)
	authTokenExpireTime() time.Time
	refreshTokenExpireTime() time.Time
	getJwt(accountID string, level int) Jwt
	//getAuthToken(id string, level string) (string, string)
	//getRefreshToken(id string, level string) (string, string)
	exchangeRefreshToken(tokenString string) (bool, string, string)
}

type AuthService struct {
	tokenAuth            *jwtauth.JWTAuth
	accountServ          account.IAccountService
	authTokenDuration    time.Duration
	refreshTokenDuration time.Duration
}

func NewAuthService(
	tAuth *jwtauth.JWTAuth,
	accountServ account.IAccountService,
) *AuthService {
	return &AuthService{
		tokenAuth:            tAuth,
		accountServ:          accountServ,
		authTokenDuration:    time.Hour * 12,
		refreshTokenDuration: time.Hour * 13,
	}
}

func (s *AuthService) setAuthTokenDuration(duration time.Duration) {
	s.authTokenDuration = duration
}

func (s *AuthService) authTokenExpireTime() time.Time {
	return time.Now().Add(s.authTokenDuration)
}
func (s *AuthService) refreshTokenExpireTime() time.Time {
	return time.Now().Add(s.refreshTokenDuration)
}

func (s *AuthService) getJwt(accountID string, level int) Jwt {
	authToken, expire := s.authToken(accountID, level)
	refToken, refExpire := s.getRefreshToken(accountID, level)
	rToken := RefreshToken{Token: refToken, Expiration: refExpire}
	return Jwt{Token: authToken, Expiration: expire, RefreshToken: rToken}
}

// LoginInfo id is embedded in
func (s *AuthService) authToken(id string, level int) (string, string) {
	aTokenClaims := map[string]interface{}{
		"id":         id,
		"token_type": "auth",
		"level":      strconv.Itoa(level),
	}
	jwtauth.SetExpiry(aTokenClaims, s.authTokenExpireTime())
	_, authToken, _ := s.tokenAuth.Encode(aTokenClaims)
	return authToken, s.authTokenExpireTime().String()
}

func (s *AuthService) getRefreshToken(id string, level int) (string, string) {
	rTokenClaims := map[string]interface{}{
		"id":         id,
		"token_type": "refresh",
		"level":      strconv.Itoa(level),
	}
	jwtauth.SetExpiry(rTokenClaims, s.refreshTokenExpireTime())
	_, refreshToken, _ := s.tokenAuth.Encode(rTokenClaims)
	return refreshToken, s.refreshTokenExpireTime().String()
}

func (s *AuthService) exchangeRefreshToken(tokenString string) (bool, string, string) {
	token, err := s.tokenAuth.Decode(tokenString)
	if err != nil {
		return false, err.Error(), ""
	}
	claims := token.PrivateClaims()
	if claims["token_type"] != "refresh" {
		return false, "This is not s refresh token", ""
	}
	levelStr := claims["level"].(string)
	level, _ := strconv.Atoi(levelStr)
	rToken, expirationTime := s.getRefreshToken(claims["id"].(string), level)
	return true, rToken, expirationTime
}

//
//func (s *AuthService) AuthenticateByAccountIDAndPWD(id string, pwd string) (bool, error) {
//	return s.accountServ.AuthenticateByAccountIDAndPWD(id, pwd)
//}
//
//func (s *AuthService) authenticatePreToken(tokenString string) error {
//	token, err := s.tokenAuth.Decode(tokenString)
//	if token == nil || err != nil {
//		return err
//	}
//	claims := token.PrivateClaims()
//	if claims["token_type"] != "pre-auth" {
//		return errors.New("invalid pre-auth token")
//	}
//
//	id := fmt.Sprintf("%v", claims["id"])
//	pwd := fmt.Sprintf("%v", claims["password"])
//	if len(id) == 0 || len(pwd) == 0 {
//		return errors.New("invalid pre-auth token. id or pwd contains zero-length string")
//	}
//	passed, err := s.accountServ.AuthenticateByAccountIDAndPWD(id, pwd)
//	if !passed || err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *AuthService) getEmailAddress(tokenString string) string {
//	token, err := s.tokenAuth.Decode(tokenString)
//	if token == nil || err != nil {
//		return ""
//	}
//	claims := token.PrivateClaims()
//	email := fmt.Sprintf("%v", claims["email"])
//	return email
//}
//
//func (s *AuthService) getAccountIdFromToken(tokenString string) string {
//	token, err := s.tokenAuth.Decode(tokenString)
//	if token == nil || err != nil {
//		return ""
//	}
//	claims := token.PrivateClaims()
//	id := fmt.Sprintf("%v", claims["id"])
//	return id
//}
//
//func (s *AuthService) getAccount(r *http.Request) (account.LoginInfo, error) {
//	_, claims, err := jwtauth.FromContext(r.Context())
//	if err != nil {
//		return account.LoginInfo{}, err
//	}
//	id := claims["id"].(string)
//	if len(id) == 0 {
//		return account.LoginInfo{}, errors.New("id is blank")
//	}
//	return s.accountServ.GetAccountByAccountID(id)
//}
//
//func (s *AuthService) getAccountFromToken(tokenString string) (account.LoginInfo, error) {
//	token, err := s.tokenAuth.Decode(tokenString)
//	if token == nil || err != nil {
//		return account.LoginInfo{}, nil
//	}
//	claims := token.PrivateClaims()
//	id := fmt.Sprintf("%v", claims["id"])
//	if len(id) == 0 {
//		return account.LoginInfo{}, errors.New("id is blank")
//	}
//	return s.accountServ.GetAccountByAccountID(id)
//}
//
//func (s *AuthService) GetTokensForAccount(accountID string) (bool, string, string, string, string, error) {
//	acc, err := s.accountServ.GetAccountByAccountID(accountID)
//	if err != nil {
//		return false, "", "", "", "", err
//	}
//	return s.getJwt(accountID, acc.level)
//}
