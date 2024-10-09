package service

import (
	"simple_bank_solid/api/repository"
	"simple_bank_solid/db"
	"simple_bank_solid/model/web/response"

	"gorm.io/gorm"
)

type AccountService interface {
	CreateAccount(UserId int64) (response.AccountResponse, error)
	DeleteAccount(Id int64) error
	FetchAccountById(Id int64) (response.AccountResponse, error)
	FetchAllAccountByUser(UserId int64) []response.AccountResponse
	FetchAllAccount() []response.AccountResponse
}

type AccountServiceImpl struct {
	accountRepo repository.AccountRepository
	db          *gorm.DB
}

// CreateAccount implements AccountService.
func (a *AccountServiceImpl) CreateAccount(UserId int64) (response.AccountResponse, error) {
	panic("unimplemented")
}

// DeleteAccount implements AccountService.
func (a *AccountServiceImpl) DeleteAccount(Id int64) error {
	panic("unimplemented")
}

// FetchAccountById implements AccountService.
func (a *AccountServiceImpl) FetchAccountById(Id int64) (response.AccountResponse, error) {
	panic("unimplemented")
}

// FetchAllAccount implements AccountService.
func (a *AccountServiceImpl) FetchAllAccount() []response.AccountResponse {
	panic("unimplemented")
}

// FetchAllAccountByUser implements AccountService.
func (a *AccountServiceImpl) FetchAllAccountByUser(UserId int64) []response.AccountResponse {
	panic("unimplemented")
}

func NewAccountService(AccountRepo repository.AccountRepository) AccountService {
	con := db.GetConnection()
	return &AccountServiceImpl{
		accountRepo: AccountRepo,
		db:          con,
	}
}
