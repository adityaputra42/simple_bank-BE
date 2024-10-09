package repository

import (
	"simple_bank_solid/db"
	"simple_bank_solid/helper"
	"simple_bank_solid/model/domain"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(account domain.Account) domain.Account
	Update(account domain.Account) domain.Account
	Delete(account domain.Account)
	FindById(accountId int) (domain.Account, error)
	FindAllbyUserId(UserId int) []domain.Account
	FindAll() []domain.Account
}

type AccountRepositoryImpl struct {
	db *gorm.DB
}

// Create implements AccountReposiotry.
func (a *AccountRepositoryImpl) Create(account domain.Account) domain.Account {
	result := a.db.Create(&account)
	helper.PanicIfError(result.Error)
	return account
}

// Delete implements AccountReposiotry.
func (a *AccountRepositoryImpl) Delete(account domain.Account) {
	result := a.db.Delete(&account)
	helper.PanicIfError(result.Error)
}

// FindAll implements AccountReposiotry.
func (a *AccountRepositoryImpl) FindAll() []domain.Account {
	accounts := []domain.Account{}
	result := a.db.Find(&accounts)
	helper.PanicIfError(result.Error)
	return accounts
}

// FindAllbyUserId implements AccountReposiotry.
func (a *AccountRepositoryImpl) FindAllbyUserId(UserId int) []domain.Account {
	accounts := []domain.Account{}
	result := a.db.Find(&accounts, "user_id = ?", UserId)
	helper.PanicIfError(result.Error)
	return accounts
}

// FindById implements AccountReposiotry.
func (a *AccountRepositoryImpl) FindById(accountId int) (domain.Account, error) {
	account := domain.Account{}
	err := a.db.Model(&domain.Account{}).Take(&account, "id =?", accountId).Error
	helper.PanicIfError(err)
	return account, err
}

// Update implements AccountReposiotry.
func (a *AccountRepositoryImpl) Update(account domain.Account) domain.Account {
	result := a.db.Save(&account)
	helper.PanicIfError(result.Error)
	return account
}

func NewAccountRepository() AccountRepository {
	db := db.GetConnection()
	return &AccountRepositoryImpl{db: db}
}
