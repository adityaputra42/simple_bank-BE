package service

import (
	"simple_bank_solid/api/repository"
	"simple_bank_solid/db"
	"simple_bank_solid/model/domain"
	"simple_bank_solid/model/web/request"
	"simple_bank_solid/model/web/response"

	"gorm.io/gorm"
)

type UserService interface {
	Create(req request.CreateUser) (response.UserResponse, error)
	Login(username, password string) (bool, error)
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
func (u UserServiceImpl) Create(req request.CreateUser) (response.UserResponse, error) {
	var user response.UserResponse

	err := u.db.Transaction(func(tx *gorm.DB) error {

		userData := domain.User{
			FullName: req.FullName,
			Email:    req.Email,
			Username: req.Username,
			Password: req.Password,
		}
		userData = u.userRepo.Create(userData, tx)

		account := domain.Account{
			UserId:   userData.ID,
			Balance:  0,
			Currency: "IDR",
		}
		account = u.accountRepo.Create(account)

		return nil
	})

	if err != nil {
		return user, err
	}

	return user, nil

}

// Delete implements UserService.
func (u UserServiceImpl) Delete(userId int) error {
	panic("unimplemented")
}

// FecthUser implements UserService.
func (u UserServiceImpl) FecthUser(username string) (response.UserResponse, error) {
	panic("unimplemented")
}

// Login implements UserService.
func (u UserServiceImpl) Login(username string, password string) (bool, error) {
	panic("unimplemented")
}

// UpdatePassword implements UserService.
func (u UserServiceImpl) UpdatePassword(req request.UpdateUser, username string) (response.UserResponse, error) {
	panic("unimplemented")
}

func NewUserService(AccountRepo repository.AccountRepository, UserRepo repository.UserRepository) UserService {
	con := db.GetConnection()
	return UserServiceImpl{
		accountRepo: AccountRepo,
		userRepo:    UserRepo,
		db:          con,
	}
}
