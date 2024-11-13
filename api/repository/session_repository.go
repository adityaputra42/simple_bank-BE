package repository

import (
	"simple_bank_solid/db"
	"simple_bank_solid/model/domain"

	"gorm.io/gorm"
)

type SessionRepository interface {
	Create(userSessions domain.UserSessions) (domain.UserSessions, error)
	FindById(userSessionsId string) (domain.UserSessions, error)
}

type SessionRepositoryImpl struct {
	db *gorm.DB
}

// Create implements SessionRepository.
func (s SessionRepositoryImpl) Create(userSessions domain.UserSessions) (domain.UserSessions, error) {
	result := s.db.Create(&userSessions)

	return userSessions, result.Error
}

// FindById implements SessionRepository.
func (s SessionRepositoryImpl) FindById(userSessionsId string) (domain.UserSessions, error) {
	session := domain.UserSessions{}
	err := s.db.Model(&domain.UserSessions{}).Take(&session, "id =?", userSessionsId).Error
	return session, err
}

func NewSessionRepository() SessionRepository {
	db := db.GetConnection()
	return SessionRepositoryImpl{db: db}
}
