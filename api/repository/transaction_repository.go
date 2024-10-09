package repository

import (
	"simple_bank_solid/db"
	"simple_bank_solid/helper"
	"simple_bank_solid/model/domain"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction domain.Transaction) domain.Transaction
	FindById(transactionId int) (domain.Transaction, error)
	FindAllbyUserId(userId int) []domain.Transaction
	FindAll() []domain.Transaction
}

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

// Create implements TransactionRepository.
func (t *TransactionRepositoryImpl) Create(transaction domain.Transaction) domain.Transaction {
	result := t.db.Create(&transaction)
	helper.PanicIfError(result.Error)
	return transaction
}

// FindAll implements TransactionRepository.
func (t *TransactionRepositoryImpl) FindAll() []domain.Transaction {
	transactions := []domain.Transaction{}
	result := t.db.Find(&transactions)
	helper.PanicIfError(result.Error)
	return transactions
}

// FindAllbyUserId implements TransactionRepository.
func (t *TransactionRepositoryImpl) FindAllbyUserId(userId int) []domain.Transaction {
	transactions := []domain.Transaction{}
	result := t.db.Find(&transactions, "user_id = ?", userId)
	helper.PanicIfError(result.Error)
	return transactions
}

// FindById implements TransactionRepository.
func (t *TransactionRepositoryImpl) FindById(transactionId int) (domain.Transaction, error) {
	transaction := domain.Transaction{}
	err := t.db.Model(&domain.Account{}).Take(&transaction, "id =?", transactionId).Error
	helper.PanicIfError(err)
	return transaction, err
}

func NewTransactionRepository() TransactionRepository {
	con := db.GetConnection()
	return &TransactionRepositoryImpl{db: con}
}
