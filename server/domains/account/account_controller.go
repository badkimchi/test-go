package account

import (
	"app/util/resp"
	"errors"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type AccountController struct {
	serv IAccountService
}

func NewAccountController(
	serv IAccountService,
) *AccountController {
	return &AccountController{
		serv: serv,
	}
}

func (c *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if len(idStr) == 0 {
		resp.Bad(w, r, errors.New("id must be set"))
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Bad(w, r, err)
		return
	}
	acc, err := c.serv.GetAccount(id)
	if err != nil {
		resp.Bad(w, r, err)
		return
	}

	resp.Data(w, r, acc)
}

//// DeleteAccount
//// @Security ApiKeyAuth
//// @Summary Deletes an account
//// @Tags AccountAccounts
//// @Description Deletes an account
//// @Accept  json
//// @Produce  json
//// @Param id path int true "account id"
//// @Success 200
//// @Failure 400
//// @Router /accounts/{id} [delete]
//func (c *AccountController) DeleteAccount(w http.ResponseWriter, r *http.Request) {
//	err := c.serv.DeleteAccountByAccountID(chi.URLParam(r, "id"))
//	if err != nil {
//		resp.Bad(w, r, err)
//		return
//	}
//
//	resp.Data(w, r, "account was deleted")
//}

//
//// UpdateAccount
//// @Security ApiKeyAuth
//// @Summary Update an account details. Note that this end point does not change password
//// @Tags AccountAccounts
//// @Description Update an account details. Note that this end point does not change password
//// @Accept  json
//// @Produce  json
//// @Param account body Account Account "Account"
//// @Success 200
//// @Failure 400
//// @Router /accounts/{id} [put]
//func (c *AccountController) UpdateAccount(w http.ResponseWriter, r *http.Request) {
//	id, err := strconv.Atoi(chi.URLParam(r, "id"))
//	if err != nil {
//		resp.Bad(w, r, err)
//		return
//	}
//
//	decoder := json.NewDecoder(r.Body)
//	var req Account
//	err = decoder.Decode(&req)
//	if err != nil {
//		resp.Bad(w, r, err)
//		return
//	}
//
//	account, err := c.serv.UpdateAccount(req, id)
//	if err != nil {
//		resp.Bad(w, r, err)
//		return
//	}
//
//	resp.Data(w, r, account)
//}
