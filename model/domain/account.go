package domain

import "time"

type Account struct {
	ID                  int64         `gorm:"primaryKey;column:id;autoIncrement"`
	UserId              int64         `gorm:"primaryKey;column:user_id"`
	Balance             int64         `gorm:"column:balance"`
	Currency            string        `gorm:"column:currency"`
	CreatedAt           time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt           time.Time     `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User                User          `gorm:"foreignKey:user_id;references:id"`
	SendTransactions    []Transaction `gorm:"foreignKey:from_account_id;references:id"`
	ReceiveTransactions []Transaction `gorm:"foreignKey:to_account_id;references:id"`
	Entries             []Entries     `gorm:"foreignKey:account_id;references:id"`
}

func (u *Account) TableName() string {
	return "accounts"
}
