package repository

import (
	"simple_bank_solid/db"
	"simple_bank_solid/helper"
	"simple_bank_solid/model/domain"

	"gorm.io/gorm"
)

type EntriesRepository interface {
	Create(entity domain.Entries, tx *gorm.DB) (domain.Entries, error)
	FindById(entityId int) (domain.Entries, error)
	FindAll() []domain.Entries
}

type EntriesRepositoryImpl struct {
	db *gorm.DB
}

// Create implements EntriesRepository.
func (e *EntriesRepositoryImpl) Create(entity domain.Entries, tx *gorm.DB) (domain.Entries, error) {
	result := tx.Create(&entity)
	helper.PanicIfError(result.Error)
	return entity, result.Error
}

// FindAll implements EntriesRepository.
func (e *EntriesRepositoryImpl) FindAll() []domain.Entries {
	entries := []domain.Entries{}
	result := e.db.Find(&entries)
	helper.PanicIfError(result.Error)
	return entries
}

// FindById implements EntriesRepository.
func (e *EntriesRepositoryImpl) FindById(entityId int) (domain.Entries, error) {
	entries := domain.Entries{}
	result := e.db.Model(&domain.Entries{}).Take(&entries, "id =?", entityId)
	helper.PanicIfError(result.Error)
	return entries, result.Error
}

func NewEntriesRepository() EntriesRepository {
	con := db.GetConnection()
	return &EntriesRepositoryImpl{db: con}
}
