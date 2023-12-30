package account

import (
	"app/sql/db"
)

type AccountRepo struct {
	db *db.Queries
}

func NewAccountRepo(db *db.Queries) *AccountRepo {
	return &AccountRepo{
		db: db,
	}
}

//
//func (r *AccountRepo) GetAccountByIDIncludingPassword(id uint) (Account, error) {
//	var acc Account
//	if err := r.db.Where("id = ?", id).First(&acc).Error; err != nil {
//		return Account{}, l.Log(err)
//	}
//	return acc, nil
//}
//
//func (r *AccountRepo) GetAccountByAccountIncludingPassword(accountID string) (Account, error) {
//	var acc Account
//	if err := r.db.Where("account_id = ?", accountID).First(&acc).Error; err != nil {
//		return Account{}, l.Log(errors.New("unable to find the matching account with is " + accountID))
//	}
//	return acc, nil
//}

//func (r *AccountRepo) GetAccountByAccountID(accountID int64) (db.Author, error) {
//	acc, err := r.db.GetAuthor(context.Background(), accountID)
//	if err != nil {
//		return db.Author{}, err
//	}
//	return acc, nil
//}

//
//func (r *AccountRepo) GetAll() ([]Account, error) {
//	var accs []Account
//	if err := r.db.Find(&accs).Error; err != nil {
//		return accs, l.Log(err)
//	}
//
//	var retAccs []Account
//	for _, acc := range accs {
//		acc.PWD = "NOT-SHOWN"
//		retAccs = append(retAccs, acc)
//	}
//	return retAccs, nil
//}

//func (r *AccountRepo) Create(accountID string, pwd string) (Account, error) {
//	//check if same name exists.
//	var acc Account
//	if err := r.db.Where("account_id = ?", accountID).First(&acc).Error; err != nil && err.Error() != "record not found" {
//		return Account{}, l.Log(err)
//	}
//	acc = Account{
//		AccountID:         accountID,
//		PWD:            pwd,
//		PrivilegeTitle: "monitor", //create with the lowest privilege and only allow admins to elevate privilege
//	}
//	if err := r.db.Create(&acc).Error; err != nil {
//		return Account{}, l.Log(err)
//	}
//	acc.PWD = "NOT-SHOWN"
//	return acc, nil
//}
//
//func (r *AccountRepo) Delete(id uint) error {
//	var acc Account
//	if err := r.db.Where("id = ?", id).First(&acc).Unscoped().Delete(&acc).Error; err != nil {
//		return l.Log(err)
//	}
//	return nil
//}
//
//func (r *AccountRepo) DeleteByAccountId(id string) error {
//	var acc Account
//	if err := r.db.Where("account_id = ?", id).First(&acc).Unscoped().Delete(&acc).Error; err != nil {
//		return l.Log(err)
//	}
//	return nil
//}
//
//func (r *AccountRepo) Save(acc *Account) error {
//	if err := r.db.Save(&acc).Error; err != nil {
//		return l.Log(err)
//	}
//	return nil
//}
