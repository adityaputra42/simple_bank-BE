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
	Delete(transaction domain.Transaction) error
	FindAll() ([]domain.Transaction, error)
}

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

// Delete implements TransactionRepository.
func (t *TransactionRepositoryImpl) Delete(transaction domain.Transaction) error {
	err := t.db.Delete(transaction).Error
	return err
}

// Create implements TransactionRepository.
func (t *TransactionRepositoryImpl) Create(transaction domain.Transaction, tx *gorm.DB) (domain.Transaction, error) {
	result := tx.Create(&transaction)
	return transaction, result.Error
}

// FindAll implements TransactionRepository.
func (t *TransactionRepositoryImpl) FindAll() ([]domain.Transaction, error) {
	transactions := []domain.Transaction{}
	result := t.db.Model(&domain.Transaction{}).Preload("FromAccount").Preload("ToAccount").Find(&transactions)

	return transactions, result.Error
}

// FindAllbyUserId implements TransactionRepository.
func (t *TransactionRepositoryImpl) FindAllbyUserId(userId int) ([]domain.Transaction, error) {
	transactions := []domain.Transaction{}
	result := t.db.Model(&domain.Transaction{}).Joins("JOIN accounts AS from_accounts ON from_accounts.id = transactions.from_account_id").
		Joins("JOIN accounts AS to_accounts ON to_accounts.id = transactions.to_account_id").
		Where("from_accounts.user_id = ? OR to_accounts.user_id = ?", userId, userId).Preload("FromAccount").Preload("ToAccount").Find(&transactions)
	return transactions, result.Error
}

// FindById implements TransactionRepository.
func (t *TransactionRepositoryImpl) FindById(transactionId string) (domain.Transaction, error) {
	transaction := domain.Transaction{}
	err := t.db.Model(&domain.Transaction{}).Preload("FromAccount").Preload("ToAccount").Take(&transaction, "id =?", transactionId).Error
	return transaction, err
}

func NewTransactionRepository() TransactionRepository {
	con := db.GetConnection()
	return &TransactionRepositoryImpl{db: con}
}
