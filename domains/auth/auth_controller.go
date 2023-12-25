package auth

import (
	"net/http"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// TestGet
// @Summary Exchange password and userID with security token.
// @Tags Auth
// @Description depending on whether two-factor auth is enabled, api will return pre-auth token or auth token
// @Accept  json
// @Produce  json
// @Param account body account.AccountCredentials true "AccountCredentials"
// @Success 200
// @Failure 400
// @Router /auth/token [post]
func (c *AuthController) TestGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("test"))
}
