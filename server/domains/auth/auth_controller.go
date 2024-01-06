package auth

import (
	"app/domains/account"
	"app/util/resp"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/jwtauth"
	"golang.org/x/oauth2"
	"net/http"
)

type IController interface {
}

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
	//googleUrl := "https://www.googleapis.com/oauth2/v3/userinfo"
	//googleUrl = ""
	var Endpoint = oauth2.Endpoint{
		AuthURL:       "https://accounts.google.com/o/oauth2/auth",
		TokenURL:      "https://oauth2.googleapis.com/token",
		DeviceAuthURL: "https://oauth2.googleapis.com/device/code",
		AuthStyle:     oauth2.AuthStyleInParams,
	}
	conf := oauth2.Config{
		ClientID:     "773325553700-oluqkagk36js85vlqh55dselui6dvpar.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:     Endpoint,
		RedirectURL:  "postmessage",
		Scopes:       []string{"https://www.googleapis.com/auth/drive.metadata.readonly"},
	}
	token, err := conf.Exchange(context.Background(), tokenReq.Token)
	if err != nil {
		resp.Data(w, r, token)
		return
	}

	//cID := "773325553700-oluqkagk36js85vlqh55dselui6dvpar.apps.googleusercontent.com"
	//cS := "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"
	//redirectUrl := "postmessage"
	//grantType := "authorization_code"
	//body := fmt.Sprintf("{'code':%s, 'client_id':%s, 'client_secret': %s, 'redirect_uri': %s, 'grant_type': %s}",
	//	tokenReq.Token, cID, cS, redirectUrl, grantType)
	//req, err := http.NewRequest("GET", "https://oauth2.googleapis.com/token", bytes.NewBuffer([]byte(body)))
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenReq.Token))
	//if err != nil {
	//	resp.Bad(w, r, err)
	//	return
	//}
	//
	////req.Header.Set("Content-type", "application/json; charset=utf-8")
	////req.Header.Set("Accept", "text/plain")
	//
	//client := &http.Client{Timeout: time.Millisecond * 1000}
	//
	//apiResp, err := client.Do(req)
	//if err != nil {
	//	resp.Bad(w, r, err)
	//	return
	//}
	//data, err := io.ReadAll(apiResp.Body)
	//resp.OK(w, r, string(data))

	//googleUrl := "https://www.googleapis.com/oauth2/v3/userinfo"
	//req, err := http.NewRequest("GET", googleUrl, bytes.NewBuffer([]byte("{}")))
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenReq.Token))
	//if err != nil {
	//	resp.Bad(w, r, err)
	//	return
	//}
	//
	////req.Header.Set("Content-type", "application/json; charset=utf-8")
	////req.Header.Set("Accept", "text/plain")
	//
	//client := &http.Client{Timeout: time.Millisecond * 1000}
	//
	//apiResp, err := client.Do(req)
	//if err != nil {
	//	resp.Bad(w, r, err)
	//	return
	//}
	//
	//if apiResp.StatusCode == http.StatusOK {
	//	data, err := io.ReadAll(apiResp.Body)
	//	if err == nil {
	//		resp.OK(w, r, string(data))
	//		return
	//	}
	//}
	// @todo user level
	level := "0"
	t := c.serv.getToken("sang", level)
	resp.Data(w, r, t)
}

func (c *Controller) OAuthLogin(w http.ResponseWriter, r *http.Request) {
	//decoder := json.NewDecoder(r.Body)
	var req OAuthRequest

	req.Token = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjkxNDEzY2Y0ZmEwY2I5MmEzYzNmNWEwNTQ1MDkxMzJjNDc2NjA5MzciLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI3NzMzMjU1NTM3MDAtb2x1cWthZ2szNmpzODV2bHFoNTVkc2VsdWk2ZHZwYXIuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI3NzMzMjU1NTM3MDAtb2x1cWthZ2szNmpzODV2bHFoNTVkc2VsdWk2ZHZwYXIuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMDkzMzc4NTYxOTcwNTM0NDI0NzUiLCJlbWFpbCI6InNhbmdwYXJrLmVuZ0BnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmJmIjoxNzA0NTc1MjQ0LCJuYW1lIjoiU2FuZyBQYXJrIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FDZzhvY0xWNG9rQ0VxQzJfQlpEZXBtODVjcHduWmVQajNORW1zdmNMc0hzcXpvST1zOTYtYyIsImdpdmVuX25hbWUiOiJTYW5nIiwiZmFtaWx5X25hbWUiOiJQYXJrIiwibG9jYWxlIjoiZW4iLCJpYXQiOjE3MDQ1NzU1NDQsImV4cCI6MTcwNDU3OTE0NCwianRpIjoiMTQ0MzQ3MTc2MmYxZWZkYzI2ZGM4Zjc4ZDI4OTY5Y2I4ODRhMDljMyJ9.LZnSM5IZH6jCEeWjC-be_ElwN4UjBQM-ZPcdYI4WWOUF1jTjpYnxXOeJkU-OSskXUg_6lQzvh2Vs1gONDkJQFcqJwqSjx7E7ghjWz2WsJEkpON6QaP_CaQOgRcXxJ0Zzl6B4oWVjQgJodXOz14W24dtJPBfjJ4J5yjadsDgvPCrjAZ76qpXYMdbAZdtHLJjBAMCn7j4HDduj-FOSoVBP5d6E7M46VhAIzUtoSUwW5WCgit7eO_dMvaVcd19gA-yREQUDN6Cf7LydrvMTqNz8ygO2mvy2IrRvoGDo2w_ZXIq1SRapMv2UFkDPYpXCCefcy7F-7qTzPGkqZKhx50nx0Q"

	//err := decoder.Decode(&req)
	//if err != nil {
	//	resp.Bad(w, r, errors.New("EOF: unable to parse token request"+err.Error()))
	//	return
	//}
	conf := oauth2.Config{
		ClientID:     "773325553700-oluqkagk36js85vlqh55dselui6dvpar.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:     oauth2.Endpoint{},
		RedirectURL:  "http://localhost",
		Scopes:       []string{"https://www.googleapis.com/auth/drive.metadata.readonly"},
	}
	token, err := conf.Exchange(context.Background(), req.Token)
	if err != nil {
		resp.Data(w, r, token)
		return
	}
	resp.Bad(w, r, err)
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
