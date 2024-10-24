package service

import (
	"simple_bank_solid/api/repository"
	"simple_bank_solid/db"
	"simple_bank_solid/helper"
	"simple_bank_solid/model/domain"
	"simple_bank_solid/model/web/request"
	"simple_bank_solid/model/web/response"

	"gorm.io/gorm"
)

type AccountService interface {
	CreateAccount(req request.AccountRequest) (response.AccountResponse, error)
	DeleteAccount(Id int64) error
	FetchAccountById(Id int64) (response.AccountResponse, error)
	FetchAllAccountByUser(UserId int64) ([]response.AccountResponse, error)
	FetchAllAccount() ([]response.AccountResponse, error)
}

type AccountServiceImpl struct {
	accountRepo repository.AccountRepository
	db          *gorm.DB
}

// CreateAccount implements AccountService.
func (a *AccountServiceImpl) CreateAccount(req request.AccountRequest) (response.AccountResponse, error) {
	var accountResp response.AccountResponse
	err := a.db.Transaction(func(tx *gorm.DB) error {
		account := domain.Account{
			UserId:   req.UserId,
			Balance:  0,
			Currency: req.Currency,
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
	err = a.accountRepo.Delete(user)
	if err != nil {
		return err
	}
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
func (a *AccountServiceImpl) FetchAllAccount() ([]response.AccountResponse, error) {
	var listAccount []response.AccountResponse
	accounts, err := a.accountRepo.FindAll()
	if err != nil {
		return listAccount, err
	}
	for _, value := range accounts {
		listAccount = append(listAccount, helper.ToAccountResponse(value))

	}
	return listAccount, nil
}

// FetchAllAccountByUser implements AccountService.
func (a *AccountServiceImpl) FetchAllAccountByUser(UserId int64) ([]response.AccountResponse, error) {
	var listAccount []response.AccountResponse
	accounts, err := a.accountRepo.FindAllbyUserId(int(UserId))
	if err != nil {
		return listAccount, err
	}
	for _, value := range accounts {
		listAccount = append(listAccount, helper.ToAccountResponse(value))

	}
	return listAccount, nil
}

func NewAccountService(AccountRepo repository.AccountRepository) AccountService {
	con := db.GetConnection()
	return &AccountServiceImpl{
		accountRepo: AccountRepo,
		db:          con,
	}
}
