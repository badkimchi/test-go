package account

import (
	"app/sql/db"
	"context"
)

type IAccountRepo interface {
	Get(id int64) (db.Account, error)
}

type AccountRepo struct {
	db *db.Queries
}

func NewAccountRepo(db *db.Queries) *AccountRepo {
	return &AccountRepo{
		db: db,
	}
}

//
//func (r *AccountRepo) GetAccountByIDIncludingPassword(id uint) (LoginInfo, error) {
//	var acc LoginInfo
//	if err := r.db.Where("id = ?", id).First(&acc).Error; err != nil {
//		return LoginInfo{}, l.Log(err)
//	}
//	return acc, nil
//}
//
//func (r *AccountRepo) GetAccountByAccountIncludingPassword(accountID string) (LoginInfo, error) {
//	var acc LoginInfo
//	if err := r.db.Where("account_id = ?", accountID).First(&acc).Error; err != nil {
//		return LoginInfo{}, l.Log(errors.New("unable to find the matching account with is " + accountID))
//	}
//	return acc, nil
//}

func (r *AccountRepo) Get(accountID int64) (db.Account, error) {
	acc, err := r.db.GetAccount(context.Background(), accountID)
	if err != nil {
		return db.Account{}, err
	}
	return acc, nil
}

//
//func (r *AccountRepo) GetAll() ([]LoginInfo, error) {
//	var accs []LoginInfo
//	if err := r.db.Find(&accs).Error; err != nil {
//		return accs, l.Log(err)
//	}
//
//	var retAccs []LoginInfo
//	for _, acc := range accs {
//		acc.PWD = "NOT-SHOWN"
//		retAccs = append(retAccs, acc)
//	}
//	return retAccs, nil
//}

//func (r *AccountRepo) Create(accountID string, pwd string) (LoginInfo, error) {
//	//check if same name exists.
//	var acc LoginInfo
//	if err := r.db.Where("account_id = ?", accountID).First(&acc).Error; err != nil && err.Error() != "record not found" {
//		return LoginInfo{}, l.Log(err)
//	}
//	acc = LoginInfo{
//		AccountID:         accountID,
//		PWD:            pwd,
//		PrivilegeTitle: "monitor", //create with the lowest privilege and only allow admins to elevate privilege
//	}
//	if err := r.db.Create(&acc).Error; err != nil {
//		return LoginInfo{}, l.Log(err)
//	}
//	acc.PWD = "NOT-SHOWN"
//	return acc, nil
//}
//
//func (r *AccountRepo) Delete(id uint) error {
//	var acc LoginInfo
//	if err := r.db.Where("id = ?", id).First(&acc).Unscoped().Delete(&acc).Error; err != nil {
//		return l.Log(err)
//	}
//	return nil
//}
//
//func (r *AccountRepo) DeleteByAccountId(id string) error {
//	var acc LoginInfo
//	if err := r.db.Where("account_id = ?", id).First(&acc).Unscoped().Delete(&acc).Error; err != nil {
//		return l.Log(err)
//	}
//	return nil
//}
//
//func (r *AccountRepo) Save(acc *LoginInfo) error {
//	if err := r.db.Save(&acc).Error; err != nil {
//		return l.Log(err)
//	}
//	return nil
//}
