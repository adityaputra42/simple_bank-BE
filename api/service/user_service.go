package service

import (
	"database/sql"
	"fmt"
	"simple_bank_solid/api/repository"
	"simple_bank_solid/db"
	"simple_bank_solid/helper"
	"simple_bank_solid/model/domain"
	"simple_bank_solid/model/web/request"
	"simple_bank_solid/model/web/response"

	"gorm.io/gorm"
)

type UserService interface {
	Create(req request.CreateUser) (response.UserResponse, error)
	Login(req request.LoginUser) (bool, response.UserResponse, error)
	UpdatePassword(req request.UpdateUser, username string) (response.UserResponse, error)
	Delete(userId int) error
	FecthUser(username string) (response.UserResponse, error)
}

type UserServiceImpl struct {
	accountRepo repository.AccountRepository
	userRepo    repository.UserRepository
	db          *gorm.DB
}

// Create implements UserService.
func (u *UserServiceImpl) Create(req request.CreateUser) (response.UserResponse, error) {
	var user response.UserResponse

	err := u.db.Transaction(func(tx *gorm.DB) error {

		userData := domain.User{
			FullName: req.FullName,
			Email:    req.Email,
			Username: req.Username,
			Password: req.Password,
		}
		userData, err := u.userRepo.Create(userData, tx)
		if err != nil {
			return err
		}
		account := domain.Account{
			UserId:   userData.ID,
			Balance:  0,
			Currency: "IDR",
		}
		account, err = u.accountRepo.Create(account, tx)
		if err != nil {
			return err
		}
		user = response.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			FullName: user.FullName,
			Email:    user.Email,
			Accounts: []response.AccountResponse{
				helper.ToAccountResponse(account),
			},
		}
		return err
	})

	if err != nil {
		return user, err
	}

	return user, nil

}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(userId int) error {
	user, err := u.userRepo.FindById(userId)
	if err != nil {
		return err
	}
	u.userRepo.Delete(user)
	return nil
}

// FecthUser implements UserService.
func (u *UserServiceImpl) FecthUser(username string) (response.UserResponse, error) {
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		helper.PanicIfError(err)
	}
	return helper.ToUserResponse(user), nil
}

// Login implements UserService.
func (u *UserServiceImpl) Login(req request.LoginUser) (bool, response.UserResponse, error) {

	user, err := u.userRepo.FindByUsername(req.Username)
	if err == sql.ErrNoRows {
		fmt.Println("User not found")
		return false, helper.ToUserResponse(user), fmt.Errorf("User not found")
	}
	if err != nil {
		return false, helper.ToUserResponse(user), fmt.Errorf("Query error")
	}
	match, err := helper.CheckPasswordHash(req.Password, user.Password)
	if !match {

		return false, helper.ToUserResponse(user), fmt.Errorf("hash and password doesn't match")
	}
	return true, helper.ToUserResponse(user), nil
}

// UpdatePassword implements UserService.
func (u *UserServiceImpl) UpdatePassword(req request.UpdateUser, username string) (response.UserResponse, error) {
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		helper.PanicIfError(err)
	}
	match, err := helper.CheckPasswordHash(req.OldPassword, user.Password)
	if !match {
		fmt.Println("hash and password doesn't match")
		helper.PanicIfError(err)
	}
	user.Password = req.Password
	user, err = u.userRepo.UpdatePassword(user)
	if err != nil {
		helper.PanicIfError(err)
	}
	return helper.ToUserResponse(user), nil

}

func NewUserService(AccountRepo repository.AccountRepository, UserRepo repository.UserRepository) UserService {
	con := db.GetConnection()
	return &UserServiceImpl{
		accountRepo: AccountRepo,
		userRepo:    UserRepo,
		db:          con,
	}
}
