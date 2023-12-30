package account

type AccountService struct {
	repo *AccountRepo
}

func NewAccountService(repo *AccountRepo) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

//func (s *AccountService) GetAccountByUserID(userID string) (Account, error) {
//	return s.repo.GetAccountByUserID(userID)
//}
//
//func (s *AccountService) GetAccountByUserIDIncludingPassword(userID string) (Account, error) {
//	return s.repo.GetAccountByUserIncludingPassword(userID)
//}
//
//func (s *AccountService) GetAllAccounts() ([]Account, error) {
//	return s.repo.GetAll()
//}
//
///*
//UpdateUsers
//Note that this does not change password
//*/
//func (s *AccountService) UpdateUsers(accounts []Account) ([]Account, error) {
//	for _, acc := range accounts {
//		accOld, err := s.repo.GetAccountByIDIncludingPassword(acc.ID)
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
//	return accounts, nil
//}
//
///*
//UpdateUser
//Note that this does not change password
//*/
//func (s *AccountService) UpdateUser(acc Account, id int) (Account, error) {
//	accOld, err := s.repo.GetAccountByIDIncludingPassword(uint(id))
//	if err != nil {
//		return Account{}, err
//	}
//	acc.PWD = accOld.PWD
//	err = s.repo.Save(&acc)
//	if err != nil {
//		return Account{}, err
//	}
//
//	return acc, nil
//}
//
//func (s *AccountService) UpdateUserIncludingPassword(acc Account) error {
//	err := s.repo.Save(&acc)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *AccountService) AuthenticateByUserIDAndPWD(ID string, PWD string) (bool, error) {
//	acc, err := s.repo.GetAccountByUserIncludingPassword(ID)
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
//func (s *AccountService) CreateAccount(userID string, pwd string) (Account, error) {
//	encPwd, err := u.GetEncryptedPassword(pwd)
//	if err != nil {
//		return Account{}, err
//	}
//	return s.repo.Create(userID, encPwd)
//}
//
//func (s *AccountService) DeleteAccountByUserID(id string) error {
//	return s.repo.DeleteByUserId(id)
//}
//
//func (s *AccountService) ChangeUserPwd(userId, oldPwd, newPwd string) error {
//	acc, err := s.GetAccountByUserIDIncludingPassword(userId)
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
//func (s *AccountService) ResetPassword(userId string) error {
//	acc, err := s.GetAccountByUserIDIncludingPassword(userId)
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
