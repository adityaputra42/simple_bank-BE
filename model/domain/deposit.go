package domain

import (
	"time"

	"gorm.io/gorm"
)

type Deposit struct {
	ID        int64          `gorm:"primaryKey;column:id;autoIncrement"`
	Amount    int64          `gorm:"column:amount"`
	Currency  string         `gorm:"column:currency"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime;<-:create"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	AccountId int64          `gorm:"column:account_id"`
	Account   Account        `gorm:"foreignKey:account_id;references:id"`
}

func (u *Deposit) TableName() string {
	return "deposits"
}
