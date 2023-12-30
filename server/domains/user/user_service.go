package user

type UserService struct {
	repo *UserRepo
}

func NewUserService(repo *UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

//func (s *UserService) GetUserByUserID(userID string) (User, error) {
//	return s.repo.GetUserByUserID(userID)
//}
//
//func (s *UserService) GetUserByUserIDIncludingPassword(userID string) (User, error) {
//	return s.repo.GetUserByUserIncludingPassword(userID)
//}
//
//func (s *UserService) GetAllUsers() ([]User, error) {
//	return s.repo.GetAll()
//}
//
///*
//UpdateUsers
//Note that this does not change password
//*/
//func (s *UserService) UpdateUsers(users []User) ([]User, error) {
//	for _, acc := range users {
//		accOld, err := s.repo.GetUserByIDIncludingPassword(acc.ID)
//		if err != nil {
//			return nil, err
//		}
//		acc.PWD = accOld.PWD
//		err = s.repo.Save(&acc)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	return users, nil
//}
//
///*
//UpdateUser
//Note that this does not change password
//*/
//func (s *UserService) UpdateUser(acc User, id int) (User, error) {
//	accOld, err := s.repo.GetUserByIDIncludingPassword(uint(id))
//	if err != nil {
//		return User{}, err
//	}
//	acc.PWD = accOld.PWD
//	err = s.repo.Save(&acc)
//	if err != nil {
//		return User{}, err
//	}
//
//	return acc, nil
//}
//
//func (s *UserService) UpdateUserIncludingPassword(acc User) error {
//	err := s.repo.Save(&acc)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *UserService) AuthenticateByUserIDAndPWD(ID string, PWD string) (bool, error) {
//	acc, err := s.repo.GetUserByUserIncludingPassword(ID)
//	if err != nil {
//		return false, err
//	}
//
//	if u.ComparePwdHashWithPwd(acc.PWD, PWD) {
//		return true, nil
//	}
//
//	return false, errors.New("incorrect password")
//}
//
//func (s *UserService) CreateUser(userID string, pwd string) (User, error) {
//	encPwd, err := u.GetEncryptedPassword(pwd)
//	if err != nil {
//		return User{}, err
//	}
//	return s.repo.Create(userID, encPwd)
//}
//
//func (s *UserService) DeleteUserByUserID(id string) error {
//	return s.repo.DeleteByUserId(id)
//}
//
//func (s *UserService) ChangeUserPwd(userId, oldPwd, newPwd string) error {
//	acc, err := s.GetUserByUserIDIncludingPassword(userId)
//	if err != nil {
//		return l.Log(err)
//	}
//
//	if !u.ComparePwdHashWithPwd(acc.PWD, oldPwd) {
//		return l.Log(errors.New("old password is incorrect"))
//	}
//
//	acc.PWD, err = u.GetEncryptedPassword(newPwd)
//	if err != nil {
//		return l.Log(err)
//	}
//
//	err = s.UpdateUserIncludingPassword(acc)
//	if err != nil {
//		return l.Log(err)
//	}
//	return nil
//}
//
//func (s *UserService) ResetPassword(userId string) error {
//	acc, err := s.GetUserByUserIDIncludingPassword(userId)
//	if err != nil {
//		return l.Log(err)
//	}
//
//	acc.PWD, err = u.GetEncryptedPassword(vars.FactoryPwd)
//	if err != nil {
//		return l.Log(err)
//	}
//
//	err = s.UpdateUserIncludingPassword(acc)
//	if err != nil {
//		return l.Log(err)
//	}
//	return nil
//}
