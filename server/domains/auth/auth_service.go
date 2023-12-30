package auth

import (
	"app/domains/account"
	"github.com/go-chi/jwtauth"
	"time"
)

type AuthService struct {
	tokenAuth            *jwtauth.JWTAuth
	accountServ          *account.AccountService
	authTokenDuration    time.Duration
	refreshTokenDuration time.Duration
}

func NewAuthService(
	tAuth *jwtauth.JWTAuth,
	accountServ *account.AccountService,
) *AuthService {
	return &AuthService{
		tokenAuth:            tAuth,
		accountServ:          accountServ,
		authTokenDuration:    time.Hour * 12,
		refreshTokenDuration: time.Hour * 13,
	}
}

func (s *AuthService) SetAuthTokenDuration(duration time.Duration) {
	s.authTokenDuration = duration
}

func (s *AuthService) authTokenExpireTime() time.Time {
	return time.Now().Add(s.authTokenDuration)
}
func (s *AuthService) refreshTokenExpireTime() time.Time {
	return time.Now().Add(s.refreshTokenDuration)
}

//
//func (s *AuthService) AuthenticateByUserIDAndPWD(id string, pwd string) (bool, error) {
//	return s.accountServ.AuthenticateByUserIDAndPWD(id, pwd)
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
//	passed, err := s.accountServ.AuthenticateByUserIDAndPWD(id, pwd)
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
//func (s *AuthService) getUserIdFromToken(tokenString string) string {
//	token, err := s.tokenAuth.Decode(tokenString)
//	if token == nil || err != nil {
//		return ""
//	}
//	claims := token.PrivateClaims()
//	id := fmt.Sprintf("%v", claims["id"])
//	return id
//}
//
//func (s *AuthService) getAccount(r *http.Request) (account.Account, error) {
//	_, claims, err := jwtauth.FromContext(r.Context())
//	if err != nil {
//		return account.Account{}, err
//	}
//	id := claims["id"].(string)
//	if len(id) == 0 {
//		return account.Account{}, errors.New("id is blank")
//	}
//	return s.accountServ.GetAccountByUserID(id)
//}
//
//func (s *AuthService) getAccountFromToken(tokenString string) (account.Account, error) {
//	token, err := s.tokenAuth.Decode(tokenString)
//	if token == nil || err != nil {
//		return account.Account{}, nil
//	}
//	claims := token.PrivateClaims()
//	id := fmt.Sprintf("%v", claims["id"])
//	if len(id) == 0 {
//		return account.Account{}, errors.New("id is blank")
//	}
//	return s.accountServ.GetAccountByUserID(id)
//}
//
//func (s *AuthService) GetTokensForUser(userID string) (bool, string, string, string, string, error) {
//	acc, err := s.accountServ.GetAccountByUserID(userID)
//	if err != nil {
//		return false, "", "", "", "", err
//	}
//	return s.GetTokensForUserWithPrivilegeTitle(userID, acc.PrivilegeTitle)
//}

func (s *AuthService) GetTokensForUserWithPrivilegeTitle(userID string, privilegeTitle string) (
	bool, string, string, string, string, error,
) {
	authToken, authExpireTime := s.getAuthToken(userID, privilegeTitle)
	refToken, refTokenExpireTime := s.getRefreshToken(userID, privilegeTitle)
	return true, authToken, authExpireTime, refToken, refTokenExpireTime, nil
}

// User id is embedded in
func (s *AuthService) getAuthToken(id string, privilegeTitle string) (string, string) {
	aTokenClaims := map[string]interface{}{
		"id": id, "token_type": "auth", "privilege_title": privilegeTitle,
	}
	jwtauth.SetExpiry(aTokenClaims, s.authTokenExpireTime())
	_, authToken, _ := s.tokenAuth.Encode(aTokenClaims)
	return authToken, s.authTokenExpireTime().String()
}

func (s *AuthService) getRefreshToken(id string, privilegeTitle string) (string, string) {
	rTokenClaims := map[string]interface{}{"id": id, "token_type": "refresh", "privilege_title": privilegeTitle}
	jwtauth.SetExpiry(rTokenClaims, s.refreshTokenExpireTime())
	_, refreshToken, _ := s.tokenAuth.Encode(rTokenClaims)
	return refreshToken, s.refreshTokenExpireTime().String()
}

func (s *AuthService) ExchangeRefreshToken(tokenString string) (bool, string, string) {
	token, err := s.tokenAuth.Decode(tokenString)
	if err != nil {
		return false, err.Error(), ""
	}

	claims := token.PrivateClaims()

	if claims["token_type"] != "refresh" {
		return false, "This is not s refresh token", ""
	}

	rToken, expirationTime := s.getRefreshToken(claims["id"].(string), claims["privilege_title"].(string))

	return true, rToken, expirationTime
}
