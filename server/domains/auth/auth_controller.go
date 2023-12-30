package auth

import (
	"app/domains/account"
	"app/util/resp"
	"encoding/json"
	"github.com/go-chi/jwtauth"
	"net/http"
)

type Controller struct {
	tokenAuth *jwtauth.JWTAuth
	serv      *AuthService
	accServ   *account.AccountService
}

func NewAuthController(
	tokenAuth *jwtauth.JWTAuth, hAuth *AuthService, accServ *account.AccountService,
) *Controller {
	return &Controller{
		tokenAuth: tokenAuth,
		serv:      hAuth,
		accServ:   accServ,
	}
}

func (c *Controller) TestGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	resp.OK(w, r, "test")
}

//
//// GetAuthToken
//// @Summary Exchange password and accountID with security token.
//// @Tags Auth
//// @Description depending on whether two-factor auth is enabled, api will return pre-auth token or auth token
//// @Accept  json
//// @Produce  json
//// @Param account body account.AccountCredentials true "AccountCredentials"
//// @Success 200
//// @Failure 400
//// @Router /auth/token [post]
//func (c *Controller) GetAuthToken(w http.ResponseWriter, r *http.Request) {
//	decoder := json.NewDecoder(r.Body)
//	var req account.AccountCredentials
//	err := decoder.Decode(&req)
//	if err == io.EOF {
//		resp.Bad(w, r, errors.New("EOF: unable to parse request body"))
//		return
//	}
//	if err != nil || req.AccountID == "" {
//		resp.Bad(w, r, errors.New("name and password must be passed in:"+err.Error()))
//		return
//	}
//
//	authenticated, err := c.serv.AuthenticateByAccountIDAndPWD(req.AccountID, req.PWD)
//	if err != nil {
//		resp.Bad(w, r, err)
//		return
//	}
//
//	if authenticated == false {
//		resp.Bad(w, r, errors.New("wrong password or account ID, please try again"))
//		return
//	}
//
//	account, err := c.accServ.GetAccountByAccountID(req.AccountID)
//	if err != nil {
//		resp.Bad(w, r, err)
//		return
//	}
//	success, authToken, authTokenExpireTime, refreshToken, refreshTokenExpireTime, err := c.serv.GetTokensForAccount(req.AccountID)
//	if err != nil {
//		resp.Bad(w, r, errors.New("failed to create token,"+err.Error()))
//		return
//	}
//	rToken := RefreshToken{RefreshToken: refreshToken, Expiration: refreshTokenExpireTime}
//	if success {
//		authToken := Token{AuthToken: authToken, Type: "auth", Expiration: authTokenExpireTime, RefreshToken: rToken}
//		data := TokenWithPrivilegeTitleExposed{
//			Token:          authToken,
//			PrivilegeTitle: account.PrivilegeTitle,
//		}
//		resp.Data(w, r, data)
//		return
//	}
//
//	resp.Bad(w, r, nil)
//}

// RefreshWithRefreshToken
// @Summary Refresh tokens with refresh token
// @Tags Auth
// @Description Refresh tokens with refresh token
// @Accept  json
// @Produce  json
// @Param RefreshTokenRequest body RefreshTokenRequest true "RefreshTokenRequest"
// @Success 200
// @Failure 400
// @Router /auth/refresh [post]
func (c *Controller) RefreshWithRefreshToken(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req RefreshToken
	err := decoder.Decode(&req)
	if err != nil {
		resp.Bad(w, r, err)
		return
	}

	_, refreshToken, refreshTokenExpiration := c.serv.ExchangeRefreshToken(req.RefreshToken)
	rToken := RefreshToken{
		RefreshToken: refreshToken,
		Expiration:   refreshTokenExpiration,
	}
	resp.Data(w, r, rToken)
}
