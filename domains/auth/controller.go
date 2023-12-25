package auth

import (
	"net/http"
)

type Controller struct {
}

func NewAuthController() *Controller {
	return &Controller{}
}

func (c *Controller) TestGet(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("test"))
}
