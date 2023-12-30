package account

type AccountController struct {
	serv *AccountService
}

func NewAccountController(
	serv *AccountService,
) *AccountController {
	return &AccountController{
		serv: serv,
	}
}

//
//// GetAccount
//// @Security ApiKeyAuth
//// @Summary Gets an account details
//// @Tags AccountAccounts
//// @Description Gets an account details
//// @Accept  json
//// @Produce  json
//// @Param id path int true "id"
//// @Success 200
//// @Failure 400
//// @Router /accounts/{id} [get]
//func (c *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {
//	accID := chi.URLParam(r, "id")
//	acc, err := c.serv.GetAccountByAccountID(accID)
//	if err != nil {
//		resp.Bad(w, r, err)
//		return
//	}
//
//	resp.Data(w, r, acc)
//}

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
