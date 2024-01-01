package account

import (
	"app/sql/db"
)

type IAccountService interface {
	GetAccount(id int) (db.Account, error)
}

type AccountService struct {
	repo IAccountRepo
}

func NewAccountService(repo IAccountRepo) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) GetAccount(accountID int) (db.Account, error) {
	return s.repo.Get(int64(accountID))
}

//
//func (s *AccountService) GetAccountByAccountIDIncludingPassword(accountID string) (Account, error) {
//	return s.repo.GetAccountByAccountIncludingPassword(accountID)
//}
//
//func (s *AccountService) GetAllAccounts() ([]Account, error) {
//	return s.repo.GetAll()
//}
//
///*
//UpdateAccounts
//Note that this does not change password
//*/
//func (s *AccountService) UpdateAccounts(accounts []Account) ([]Account, error) {
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
//UpdateAccount
//Note that this does not change password
//*/
//func (s *AccountService) UpdateAccount(acc Account, id int) (Account, error) {
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
//func (s *AccountService) UpdateAccountIncludingPassword(acc Account) error {
//	err := s.repo.Save(&acc)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *AccountService) AuthenticateByAccountIDAndPWD(ID string, PWD string) (bool, error) {
//	acc, err := s.repo.GetAccountByAccountIncludingPassword(ID)
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
//func (s *AccountService) CreateAccount(accountID string, pwd string) (Account, error) {
//	encPwd, err := u.GetEncryptedPassword(pwd)
//	if err != nil {
//		return Account{}, err
//	}
//	return s.repo.Create(accountID, encPwd)
//}
//
//func (s *AccountService) DeleteAccountByAccountID(id string) error {
//	return s.repo.DeleteByAccountId(id)
//}
//
//func (s *AccountService) ChangeAccountPwd(accountId, oldPwd, newPwd string) error {
//	acc, err := s.GetAccountByAccountIDIncludingPassword(accountId)
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
//	err = s.UpdateAccountIncludingPassword(acc)
//	if err != nil {
//		return l.Log(err)
//	}
//	return nil
//}
//
//func (s *AccountService) ResetPassword(accountId string) error {
//	acc, err := s.GetAccountByAccountIDIncludingPassword(accountId)
//	if err != nil {
//		return l.Log(err)
//	}
//
//	acc.PWD, err = u.GetEncryptedPassword(vars.FactoryPwd)
//	if err != nil {
//		return l.Log(err)
//	}
//
//	err = s.UpdateAccountIncludingPassword(acc)
//	if err != nil {
//		return l.Log(err)
//	}
//	return nil
//}
