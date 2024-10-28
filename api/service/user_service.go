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
	"simple_bank_solid/token"
	"time"

	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(req request.CreateUser) (response.UserResponse, error)
	CreateAdmin(req request.CreateUser) (response.UserResponse, error)
	Login(req request.LoginUser) (bool, response.LoginResponse, error)
	UpdatePassword(req request.UpdateUser, username string) (response.UserResponse, error)
	Delete(username string) error
	FecthUser(username string) (response.UserResponse, error)
	FecthAllUser() ([]response.UserResponse, error)
}

type UserServiceImpl struct {
	accountRepo repository.AccountRepository
	userRepo    repository.UserRepository
	db          *gorm.DB
	tokenMaker  token.Maker
}

// FecthAllUser implements UserService.
func (u *UserServiceImpl) FecthAllUser() ([]response.UserResponse, error) {
	users := []response.UserResponse{}
	result, err := u.userRepo.FetchAllUser()
	if err != nil {
		return users, err
	}

	for _, value := range result {
		users = append(users, helper.ToUserResponse(value))

	}
	return users, nil
}

// CreateAdmin implements UserService.
func (u *UserServiceImpl) CreateAdmin(req request.CreateUser) (response.UserResponse, error) {
	var admin response.UserResponse
	err := u.db.Transaction(func(tx *gorm.DB) error {
		userData := domain.User{
			FullName: req.FullName,
			Email:    req.Email,
			Username: req.Username,
			Password: req.Password,
			Role:     "admin",
		}
		userData, err := u.userRepo.Create(userData, tx)
		if err != nil {
			return err
		}
		admin = response.UserResponse{
			ID:        userData.ID,
			Username:  userData.Username,
			FullName:  userData.FullName,
			Email:     userData.Email,
			Role:      userData.Role,
			CreatedAt: userData.CreatedAt,
			UpdatedAt: userData.UpdatedAt,
		}
		return err
	})
	if err != nil {
		return admin, err
	}

	return admin, nil

}

// Create implements UserService.
func (u *UserServiceImpl) CreateUser(req request.CreateUser) (response.UserResponse, error) {
	var user response.UserResponse

	err := u.db.Transaction(func(tx *gorm.DB) error {

		userData := domain.User{
			FullName: req.FullName,
			Email:    req.Email,
			Username: req.Username,
			Password: req.Password,
			Role:     "member",
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
			ID:        userData.ID,
			Username:  userData.Username,
			FullName:  userData.FullName,
			Email:     userData.Email,
			Role:      userData.Role,
			CreatedAt: userData.CreatedAt,
			UpdatedAt: userData.UpdatedAt,
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
func (u *UserServiceImpl) Delete(username string) error {

	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return err
	}

	err = u.userRepo.Delete(user)
	if err != nil {
		return err
	}
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
func (u *UserServiceImpl) Login(req request.LoginUser) (bool, response.LoginResponse, error) {

	user, err := u.userRepo.FindByUsername(req.Username)
	if err == sql.ErrNoRows {
		fmt.Println("User not found")
		return false, response.LoginResponse{}, fmt.Errorf("User not found")
	}
	if err != nil {
		return false, response.LoginResponse{}, fmt.Errorf("Query error")
	}
	match, err := helper.CheckPasswordHash(req.Password, user.Password)
	if !match {

		return false, response.LoginResponse{}, fmt.Errorf("hash and password doesn't match")
	}

	accessToken, err := u.tokenMaker.CreateToken(user.Username, user.ID, time.Minute*15)

	if err != nil {
		return false, response.LoginResponse{}, err
	}
	return true, response.LoginResponse{AccessToken: accessToken, User: helper.ToUserResponse(user)}, nil
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
	maker := token.GetTokenMaker()
	return &UserServiceImpl{
		accountRepo: AccountRepo,
		userRepo:    UserRepo,
		db:          con,
		tokenMaker:  maker,
	}
}
