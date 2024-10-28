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
	Delete(user domain.User) error
	FindById(UserId int) (domain.User, error)
	FindByUsername(Username string) (domain.User, error)
	FetchAllUser() ([]domain.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// FetchAllUser implements UserRepository.
func (u *UserRepositoryImpl) FetchAllUser() ([]domain.User, error) {
	users := []domain.User{}
	err := u.db.Model(&domain.User{}).Preload("Accounts").Find(&users, "role =?", "member").Error
	return users, err
}

// Create implements UserRepository.
func (u *UserRepositoryImpl) Create(user domain.User, tx *gorm.DB) (domain.User, error) {
	hash, err := helper.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}
	newUser := domain.User{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
		Password: hash,
		Role:     user.Role,
	}
	result := tx.Create(&newUser)

	return newUser, result.Error
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(user domain.User) error {
	result := u.db.Delete(&user)

	return result.Error
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(UserId int) (domain.User, error) {
	user := domain.User{}
	err := u.db.Model(&domain.User{}).Preload("Accounts").Take(&user, "id =?", UserId).Error

	return user, err
}

// FindByUsername implements UserRepository.
func (u *UserRepositoryImpl) FindByUsername(Username string) (domain.User, error) {
	user := domain.User{}
	err := u.db.Model(&domain.User{}).Preload("Accounts").Take(&user, "username =?", Username).Error

	return user, err
}

// UpdatePassword implements UserRepository.
func (u *UserRepositoryImpl) UpdatePassword(user domain.User) (domain.User, error) {
	hash, err := helper.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}
	newUser := domain.User{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Password: hash,
	}
	result := u.db.Save(&newUser)

	return newUser, result.Error

}

func NewUserRepository() UserRepository {
	con := db.GetConnection()
	return &UserRepositoryImpl{db: con}
}
