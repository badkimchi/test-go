package user

type UserController struct {
	serv *UserService
}

func NewUserController(
	serv *UserService,
) *UserController {
	return &UserController{
		serv: serv,
	}
}

//
//// GetUser
//// @Security ApiKeyAuth
//// @Summary Gets an user details
//// @Tags UserUsers
//// @Description Gets an user details
//// @Accept  json
//// @Produce  json
//// @Param id path int true "id"
//// @Success 200
//// @Failure 400
//// @Router /users/{id} [get]
//func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
//	accID := chi.URLParam(r, "id")
//	acc, err := c.serv.GetUserByUserID(accID)
//	if err != nil {
//		resp.Bad(w, r, err)
//		return
//	}
//
//	resp.Data(w, r, acc)
//}

//// DeleteUser
//// @Security ApiKeyAuth
//// @Summary Deletes an user
//// @Tags UserUsers
//// @Description Deletes an user
//// @Accept  json
//// @Produce  json
//// @Param id path int true "user id"
//// @Success 200
//// @Failure 400
//// @Router /users/{id} [delete]
//func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
//	err := c.serv.DeleteUserByUserID(chi.URLParam(r, "id"))
//	if err != nil {
//		resp.Bad(w, r, err)
//		return
//	}
//
//	resp.Data(w, r, "user was deleted")
//}

//
//// UpdateUser
//// @Security ApiKeyAuth
//// @Summary Update an user details. Note that this end point does not change password
//// @Tags UserUsers
//// @Description Update an user details. Note that this end point does not change password
//// @Accept  json
//// @Produce  json
//// @Param user body User User "User"
//// @Success 200
//// @Failure 400
//// @Router /users/{id} [put]
//func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
//	id, err := strconv.Atoi(chi.URLParam(r, "id"))
//	if err != nil {
//		resp.Bad(w, r, err)
//		return
//	}
//
//	decoder := json.NewDecoder(r.Body)
//	var req User
//	err = decoder.Decode(&req)
//	if err != nil {
//		resp.Bad(w, r, err)
//		return
//	}
//
//	user, err := c.serv.UpdateUser(req, id)
//	if err != nil {
//		resp.Bad(w, r, err)
//		return
//	}
//
//	resp.Data(w, r, user)
//}
