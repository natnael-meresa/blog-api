package persistence

import (
	"fmt"
	"twof/blog-api/internal/constant/model"

	"gorm.io/gorm"
)

type UserPersistence interface {
	CreateUser(user *model.User) (err error)
	GetAllUsers(user *[]model.User) (err error)
	GetUser(user *model.User) (err error)
	GetUserById(uint, *model.User) (err error)
	GetUserByEmail(Email string) (model.User, error)
	UpdateUserPass(userEmail string, password string) (err error)
}

type userPersistence struct {
	db *gorm.DB
}

func UserInit(db *gorm.DB) UserPersistence {
	return &userPersistence{
		db,
	}
}

func (u userPersistence) CreateUser(user *model.User) (err error) {

	if err = u.db.Create(user).Error; err != nil {
		fmt.Println("her is the error")
		return err
	}

	return nil
}

func (u *userPersistence) GetUser(user *model.User) (err error) {
	if err = u.db.Where("email = ?", user.Email).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userPersistence) GetUserByEmail(Email string) (model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", Email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *userPersistence) GetUserById(userId uint, user *model.User) (err error) {
	if err = u.db.Where("ID = ?", userId).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userPersistence) GetAllUsers(user *[]model.User) (err error) {
	if err = u.db.Find(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userPersistence) UpdateUserPass(userEmail string, password string) (err error) {
	if err = u.db.Model(&model.User{}).Where("email = ?", userEmail).Update("password", password).Error; err != nil {
		return err
	}

	return nil
}
