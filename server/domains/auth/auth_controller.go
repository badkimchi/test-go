package auth

import (
	"app/domains/account"
	"app/util/resp"
	"encoding/json"
	"errors"
	"github.com/go-chi/jwtauth"
	"net/http"
)

type Controller struct {
	tokenAuth *jwtauth.JWTAuth
	serv      IAuthService
	accServ   account.IAccountService
}

func NewAuthController(
	tokenAuth *jwtauth.JWTAuth, hAuth IAuthService, accServ account.IAccountService,
) Controller {
	return Controller{
		tokenAuth: tokenAuth,
		serv:      hAuth,
		accServ:   accServ,
	}
}

//
//func (c *Controller) TestGet(w http.ResponseWriter, r *http.Request) {
//	w.WriteHeader(http.StatusOK)
//	resp.OK(w, r, "test")
//}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req LoginRequest
	err := decoder.Decode(&req)
	if err != nil {
		resp.Bad(w, r, errors.New("EOF: unable to parse token request"+err.Error()))
		return
	}
	if req.UserID == "" || req.Password == "" {
		resp.Bad(w, r, errors.New("name and password must be passed in"))
		return
	}

	//authenticated, err := c.serv.AuthenticateByAccountIDAndPWD(req.AccountID, req.PWD)
	//if err != nil {
	//	resp.Bad(w, r, err)
	//	return
	//}
	//
	//if authenticated == false {
	//	resp.Bad(w, r, errors.New("wrong password or account ID, please try again"))
	//	return
	//}
	//account, err := c.accServ.GetAccountByAccountID(req.AccountID)
	//if err != nil {
	//	resp.Bad(w, r, err)
	//	return
	//}

	// @todo user level
	level := 0
	token := c.serv.getToken(req.UserID, level)
	resp.Data(w, r, token)
}

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

	_, refreshToken, refreshTokenExpiration := c.serv.exchangeRefreshToken(req.Token)
	rToken := RefreshToken{
		Token:      refreshToken,
		Expiration: refreshTokenExpiration,
	}
	resp.Data(w, r, rToken)
}
