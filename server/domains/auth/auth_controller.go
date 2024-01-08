package auth

import (
	"app/conf"
	"app/domains/account"
	"app/util/resp"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/jwtauth"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

type IController interface {
}

type Controller struct {
	config    *conf.Config
	tokenAuth *jwtauth.JWTAuth
	serv      IAuthService
	accServ   account.IAccountService
}

func NewAuthController(
	config *conf.Config,
	tokenAuth *jwtauth.JWTAuth,
	hAuth IAuthService,
	accServ account.IAccountService,
) Controller {
	return Controller{
		config:    config,
		tokenAuth: tokenAuth,
		serv:      hAuth,
		accServ:   accServ,
	}
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var tokenReq OAuthRequest
	err := decoder.Decode(&tokenReq)
	if err != nil {
		resp.Bad(w, r, errors.New("EOF: unable to parse token request"+err.Error()))
		return
	}
	if tokenReq.Token == "" {
		resp.Bad(w, r, errors.New("token must be passed in"))
		return
	}
	var Endpoint = oauth2.Endpoint{
		AuthURL:       "https://accounts.google.com/o/oauth2/auth",
		TokenURL:      "https://oauth2.googleapis.com/token",
		DeviceAuthURL: "https://oauth2.googleapis.com/device/code",
		AuthStyle:     oauth2.AuthStyleInParams,
	}
	oauth := oauth2.Config{
		ClientID:     c.config.GoogleClientID,
		ClientSecret: c.config.GoogleClientSecret,
		Endpoint:     Endpoint,
		RedirectURL:  "postmessage",
		Scopes:       []string{"https://www.googleapis.com/auth/drive.metadata.readonly"},
	}
	token, err := oauth.Exchange(context.Background(), tokenReq.Token)
	if err != nil {
		resp.Data(w, r, token)
		return
	}

	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", bytes.NewBuffer([]byte("")))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	if err != nil {
		resp.Bad(w, r, err)
		return
	}
	client := &http.Client{Timeout: time.Millisecond * 1000}
	apiResp, err := client.Do(req)
	if err != nil {
		resp.Bad(w, r, err)
		return
	}

	decoder = json.NewDecoder(apiResp.Body)
	var info UserInfoDto
	err = decoder.Decode(&info)
	if err != nil {
		resp.Bad(w, r, err)
		return
	}
	level := "0"
	info.Jwt = c.serv.getJwt(info.Email, level)
	resp.Data(w, r, info)
}

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
