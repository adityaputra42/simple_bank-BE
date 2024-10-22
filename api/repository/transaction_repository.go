package repository

import (
	"simple_bank_solid/db"
	"simple_bank_solid/helper"
	"simple_bank_solid/model/domain"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction domain.Transaction, tx *gorm.DB) (domain.Transaction, error)
	FindById(transactionId string) (domain.Transaction, error)
	FindAllbyUserId(userId int) []domain.Transaction
	FindAll() []domain.Transaction
}

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

// Create implements TransactionRepository.
func (t *TransactionRepositoryImpl) Create(transaction domain.Transaction, tx *gorm.DB) (domain.Transaction, error) {
	result := tx.Create(&transaction)
	helper.PanicIfError(result.Error)
	return transaction, result.Error
}

// FindAll implements TransactionRepository.
func (t *TransactionRepositoryImpl) FindAll() []domain.Transaction {
	transactions := []domain.Transaction{}
	result := t.db.Model(&domain.Account{}).Preload("FromAccount").Preload("ToAccount").Find(&transactions)
	helper.PanicIfError(result.Error)
	return transactions
}

// FindAllbyUserId implements TransactionRepository.
func (t *TransactionRepositoryImpl) FindAllbyUserId(userId int) []domain.Transaction {
	transactions := []domain.Transaction{}
	result := t.db.Model(&domain.Account{}).Preload("FromAccount").Preload("ToAccount").Find(&transactions, "user_id = ?", userId)
	helper.PanicIfError(result.Error)
	return transactions
}

// FindById implements TransactionRepository.
func (t *TransactionRepositoryImpl) FindById(transactionId string) (domain.Transaction, error) {
	transaction := domain.Transaction{}
	err := t.db.Model(&domain.Account{}).Preload("FromAccount").Preload("ToAccount").Take(&transaction, "id =?", transactionId).Error
	helper.PanicIfError(err)
	return transaction, err
}

func NewTransactionRepository() TransactionRepository {
	con := db.GetConnection()
	return &TransactionRepositoryImpl{db: con}
}
