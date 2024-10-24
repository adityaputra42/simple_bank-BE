package repository

import (
	"simple_bank_solid/db"
	"simple_bank_solid/model/domain"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(account domain.Account, tx *gorm.DB) (domain.Account, error)
	Update(account domain.Account, tx *gorm.DB) (domain.Account, error)
	Delete(account domain.Account) error
	FindById(accountId int) (domain.Account, error)
	FindAllbyUserId(UserId int) ([]domain.Account, error)
	FindAll() ([]domain.Account, error)
}

type AccountRepositoryImpl struct {
	db *gorm.DB
}

// Create implements AccountReposiotry.
func (a *AccountRepositoryImpl) Create(account domain.Account, tx *gorm.DB) (domain.Account, error) {
	result := tx.Create(&account)

	return account, result.Error
}

// Delete implements AccountReposiotry.
func (a *AccountRepositoryImpl) Delete(account domain.Account) error {
	result := a.db.Delete(&account)
	return result.Error
}

// FindAll implements AccountReposiotry.
func (a *AccountRepositoryImpl) FindAll() ([]domain.Account, error) {
	accounts := []domain.Account{}
	result := a.db.Find(&accounts)
	return accounts, result.Error
}

// FindAllbyUserId implements AccountReposiotry.
func (a *AccountRepositoryImpl) FindAllbyUserId(UserId int) ([]domain.Account, error) {
	accounts := []domain.Account{}
	result := a.db.Find(&accounts, "user_id = ?", UserId)

	return accounts, result.Error
}

// FindById implements AccountReposiotry.
func (a *AccountRepositoryImpl) FindById(accountId int) (domain.Account, error) {
	account := domain.Account{}
	err := a.db.Model(&domain.Account{}).Take(&account, "id =?", accountId).Error

	return account, err
}

// Update implements AccountReposiotry.
func (a *AccountRepositoryImpl) Update(account domain.Account, tx *gorm.DB) (domain.Account, error) {
	result := tx.Save(&account)

	return account, result.Error
}

func NewAccountRepository() AccountRepository {
	db := db.GetConnection()
	return &AccountRepositoryImpl{db: db}
}
