package repository

import (
	"simple_bank_solid/db"
	"simple_bank_solid/model/domain"

	"gorm.io/gorm"
)

type DepositRepository interface {
	Create(deposit domain.Deposit, tx *gorm.DB) (domain.Deposit, error)
	Delete(deposit domain.Deposit) error
	FindById(depositId int) (domain.Deposit, error)
	FindAllbyUser(userId int64) ([]domain.Deposit, error)
	FindAll() ([]domain.Deposit, error)
}

type DepositRepositoryImpl struct {
	db *gorm.DB
}

// FindAllbyUser implements DepositRepository.
func (d *DepositRepositoryImpl) FindAllbyUser(userId int64) ([]domain.Deposit, error) {
	deposits := []domain.Deposit{}
	result := d.db.Model(&domain.Deposit{}).Joins("JOIN accounts ON accounts.id = deposits.account_id").Where("accounts.user_id = ?", userId).Preload("Account").Find(&deposits)

	return deposits, result.Error
}

// Create implements DepositRepository.
func (d *DepositRepositoryImpl) Create(deposit domain.Deposit, tx *gorm.DB) (domain.Deposit, error) {
	result := tx.Create(&deposit)

	return deposit, result.Error
}

// Delete implements DepositRepository.
func (d *DepositRepositoryImpl) Delete(deposit domain.Deposit) error {
	result := d.db.Delete(&deposit)
	return result.Error
}

// FindAll implements DepositRepository.
func (d *DepositRepositoryImpl) FindAll() ([]domain.Deposit, error) {
	deposits := []domain.Deposit{}
	result := d.db.Model(&domain.Deposit{}).Preload("Account").Find(&deposits)

	return deposits, result.Error
}

// FindById implements DepositRepository.
func (d *DepositRepositoryImpl) FindById(depositId int) (domain.Deposit, error) {
	deposit := domain.Deposit{}
	err := d.db.Model(&domain.Deposit{}).Preload("Account").Take(&deposit, "id = ?", depositId).Error

	return deposit, err
}

func NewDepositRepository() DepositRepository {
	con := db.GetConnection()
	return &DepositRepositoryImpl{db: con}
}
