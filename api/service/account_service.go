package service

import (
	"simple_bank_solid/api/repository"
	"simple_bank_solid/db"
	"simple_bank_solid/helper"
	"simple_bank_solid/model/domain"
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
	var accountResp response.AccountResponse
	err := a.db.Transaction(func(tx *gorm.DB) error {
		account := domain.Account{
			UserId: UserId, Balance: 0,
			Currency: "IDR",
		}
		account, err := a.accountRepo.Create(account, tx)
		if err != nil {
			return err
		}
		accountResp = response.AccountResponse{
			ID:        account.ID,
			UserId:    account.UserId,
			Balance:   account.Balance,
			Currency:  account.Currency,
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
		}
		return err
	})

	if err != nil {
		return accountResp, err
	}

	return accountResp, nil
}

// DeleteAccount implements AccountService.
func (a *AccountServiceImpl) DeleteAccount(Id int64) error {
	user, err := a.accountRepo.FindById(int(Id))
	if err != nil {
		return err
	}
	a.accountRepo.Delete(user)
	return nil
}

// FetchAccountById implements AccountService.
func (a *AccountServiceImpl) FetchAccountById(Id int64) (response.AccountResponse, error) {
	account, err := a.accountRepo.FindById(int(Id))
	if err != nil {
		return helper.ToAccountResponse(account), err
	}
	return helper.ToAccountResponse(account), nil
}

// FetchAllAccount implements AccountService.
func (a *AccountServiceImpl) FetchAllAccount() []response.AccountResponse {
	var listAccount []response.AccountResponse
	accounts := a.accountRepo.FindAll()
	for _, value := range accounts {
		listAccount = append(listAccount, helper.ToAccountResponse(value))

	}
	return listAccount
}

// FetchAllAccountByUser implements AccountService.
func (a *AccountServiceImpl) FetchAllAccountByUser(UserId int64) []response.AccountResponse {
	var listAccount []response.AccountResponse
	accounts := a.accountRepo.FindAllbyUserId(int(UserId))
	for _, value := range accounts {
		listAccount = append(listAccount, helper.ToAccountResponse(value))

	}
	return listAccount
}

func NewAccountService(AccountRepo repository.AccountRepository) AccountService {
	con := db.GetConnection()
	return &AccountServiceImpl{
		accountRepo: AccountRepo,
		db:          con,
	}
}
