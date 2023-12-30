package user

import (
	"app/sql/db"
	"context"
)

type UserRepo struct {
	db *db.Queries
}

func NewUserRepo(db *db.Queries) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

//
//func (r *UserRepo) GetUserByIDIncludingPassword(id uint) (User, error) {
//	var acc User
//	if err := r.db.Where("id = ?", id).First(&acc).Error; err != nil {
//		return User{}, l.Log(err)
//	}
//	return acc, nil
//}
//
//func (r *UserRepo) GetUserByUserIncludingPassword(userID string) (User, error) {
//	var acc User
//	if err := r.db.Where("user_id = ?", userID).First(&acc).Error; err != nil {
//		return User{}, l.Log(errors.New("unable to find the matching user with is " + userID))
//	}
//	return acc, nil
//}

func (r *UserRepo) GetUserByUserID(userID int64) (db.Author, error) {
	acc, err := r.db.GetAuthor(context.Background(), userID)
	if err != nil {
		return db.Author{}, err
	}
	return acc, nil
}

//
//func (r *UserRepo) GetAll() ([]User, error) {
//	var accs []User
//	if err := r.db.Find(&accs).Error; err != nil {
//		return accs, l.Log(err)
//	}
//
//	var retAccs []User
//	for _, acc := range accs {
//		acc.PWD = "NOT-SHOWN"
//		retAccs = append(retAccs, acc)
//	}
//	return retAccs, nil
//}

//func (r *UserRepo) Create(userID string, pwd string) (User, error) {
//	//check if same name exists.
//	var acc User
//	if err := r.db.Where("user_id = ?", userID).First(&acc).Error; err != nil && err.Error() != "record not found" {
//		return User{}, l.Log(err)
//	}
//	acc = User{
//		UserID:         userID,
//		PWD:            pwd,
//		PrivilegeTitle: "monitor", //create with the lowest privilege and only allow admins to elevate privilege
//	}
//	if err := r.db.Create(&acc).Error; err != nil {
//		return User{}, l.Log(err)
//	}
//	acc.PWD = "NOT-SHOWN"
//	return acc, nil
//}
//
//func (r *UserRepo) Delete(id uint) error {
//	var acc User
//	if err := r.db.Where("id = ?", id).First(&acc).Unscoped().Delete(&acc).Error; err != nil {
//		return l.Log(err)
//	}
//	return nil
//}
//
//func (r *UserRepo) DeleteByUserId(id string) error {
//	var acc User
//	if err := r.db.Where("user_id = ?", id).First(&acc).Unscoped().Delete(&acc).Error; err != nil {
//		return l.Log(err)
//	}
//	return nil
//}
//
//func (r *UserRepo) Save(acc *User) error {
//	if err := r.db.Save(&acc).Error; err != nil {
//		return l.Log(err)
//	}
//	return nil
//}
