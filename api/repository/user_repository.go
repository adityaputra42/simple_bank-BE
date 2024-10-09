package repository

import (
	"simple_bank_solid/db"
	"simple_bank_solid/helper"
	"simple_bank_solid/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user domain.User, tx *gorm.DB) (domain.User, error)
	UpdatePassword(user domain.User) (domain.User, error)
	Delete(user domain.User)
	FindById(UserId int) (domain.User, error)
	FindByUsername(Username string) (domain.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// Create implements UserRepository.
func (u *UserRepositoryImpl) Create(user domain.User, tx *gorm.DB) (domain.User, error) {
	hash, err := helper.HashPassword(user.Password)
	helper.PanicIfError(err)
	newUser := domain.User{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
		Password: hash,
	}
	result := tx.Create(&newUser)
	helper.PanicIfError(result.Error)

	return newUser, result.Error
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(user domain.User) {
	result := u.db.Delete(&user)
	helper.PanicIfError(result.Error)
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(UserId int) (domain.User, error) {
	user := domain.User{}
	err := u.db.Model(&domain.User{}).Preload("Accounts").Take(&user, "id =?", UserId).Error
	helper.PanicIfError(err)
	return user, err
}

// FindByUsername implements UserRepository.
func (u *UserRepositoryImpl) FindByUsername(Username string) (domain.User, error) {
	user := domain.User{}
	err := u.db.Model(&domain.User{}).Preload("Accounts").Take(&user, "username =?", Username).Error
	helper.PanicIfError(err)
	return user, err
}

// UpdatePassword implements UserRepository.
func (u *UserRepositoryImpl) UpdatePassword(user domain.User) (domain.User, error) {
	hash, err := helper.HashPassword(user.Password)
	helper.PanicIfError(err)
	newUser := domain.User{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Password: hash,
	}
	result := u.db.Save(&newUser)
	helper.PanicIfError(result.Error)
	return newUser, result.Error

}

func NewUserRepository() UserRepository {
	con := db.GetConnection()
	return &UserRepositoryImpl{db: con}
}
