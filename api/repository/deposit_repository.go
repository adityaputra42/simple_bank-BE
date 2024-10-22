package repository

import (
	"simple_bank_solid/db"
	"simple_bank_solid/helper"
	"simple_bank_solid/model/domain"

	"gorm.io/gorm"
)

type DepositRepository interface {
	Create(deposit domain.Deposit, tx *gorm.DB) (domain.Deposit, error)
	Delete(deposit domain.Deposit)
	FindById(depositId int) (domain.Deposit, error)
	FindAll() []domain.Deposit
}

type DepositRepositoryImpl struct {
	db *gorm.DB
}

// Create implements DepositRepository.
func (d *DepositRepositoryImpl) Create(deposit domain.Deposit, tx *gorm.DB) (domain.Deposit, error) {
	result := tx.Create(&deposit)
	helper.PanicIfError(result.Error)
	return deposit, result.Error
}

// Delete implements DepositRepository.
func (d *DepositRepositoryImpl) Delete(deposit domain.Deposit) {
	result := d.db.Delete(&deposit)
	helper.PanicIfError(result.Error)
}

// FindAll implements DepositRepository.
func (d *DepositRepositoryImpl) FindAll() []domain.Deposit {
	deposits := []domain.Deposit{}
	result := d.db.Model(&domain.Deposit{}).Preload("Account").Find(&deposits)
	helper.PanicIfError(result.Error)
	return deposits
}

// FindById implements DepositRepository.
func (d *DepositRepositoryImpl) FindById(depositId int) (domain.Deposit, error) {
	deposit := domain.Deposit{}
	err := d.db.Model(&domain.Deposit{}).Preload("Account").Take(&deposit, "id = ?", depositId).Error
	helper.PanicIfError(err)
	return deposit, err
}

func NewDepositRepository() DepositRepository {
	con := db.GetConnection()
	return &DepositRepositoryImpl{db: con}
}
