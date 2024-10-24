package repository

import (
	"simple_bank_solid/db"
	"simple_bank_solid/model/domain"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction domain.Transaction, tx *gorm.DB) (domain.Transaction, error)
	FindById(transactionId string) (domain.Transaction, error)
	FindAllbyUserId(userId int) ([]domain.Transaction, error)
	FindAll() ([]domain.Transaction, error)
}

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

// Create implements TransactionRepository.
func (t *TransactionRepositoryImpl) Create(transaction domain.Transaction, tx *gorm.DB) (domain.Transaction, error) {
	result := tx.Create(&transaction)

	return transaction, result.Error
}

// FindAll implements TransactionRepository.
func (t *TransactionRepositoryImpl) FindAll() ([]domain.Transaction, error) {
	transactions := []domain.Transaction{}
	result := t.db.Model(&domain.Account{}).Preload("FromAccount").Preload("ToAccount").Find(&transactions)

	return transactions, result.Error
}

// FindAllbyUserId implements TransactionRepository.
func (t *TransactionRepositoryImpl) FindAllbyUserId(userId int) ([]domain.Transaction, error) {
	transactions := []domain.Transaction{}
	result := t.db.Model(&domain.Account{}).Preload("FromAccount").Preload("ToAccount").Find(&transactions, "user_id = ?", userId)

	return transactions, result.Error
}

// FindById implements TransactionRepository.
func (t *TransactionRepositoryImpl) FindById(transactionId string) (domain.Transaction, error) {
	transaction := domain.Transaction{}
	err := t.db.Model(&domain.Account{}).Preload("FromAccount").Preload("ToAccount").Take(&transaction, "id =?", transactionId).Error
	return transaction, err
}

func NewTransactionRepository() TransactionRepository {
	con := db.GetConnection()
	return &TransactionRepositoryImpl{db: con}
}
