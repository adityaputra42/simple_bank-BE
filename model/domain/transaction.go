package domain

import "time"

type Transaction struct {
	ID            string    `gorm:"primaryKey;column:id"`
	FromAccountId int64     `gorm:"primaryKey;column:from_account_id"`
	ToAccountId   int64     `gorm:"primaryKey;column:to_account_id"`
	Amount        int64     `gorm:"column:amount"`
	Currency      string    `gorm:"column:currency"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	FromAccount   Account   `gorm:"foreignKey:from_account_id;references:id"`
	ToAccount     Account   `gorm:"foreignKey:to_account_id;references:id"`
}

func (u *Transaction) TableName() string {
	return "transactions"
}
