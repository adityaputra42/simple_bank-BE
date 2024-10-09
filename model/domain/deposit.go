package domain

import "time"

type Deposit struct {
	ID        int64     `gorm:"primaryKey;column:id;autoIncrement"`
	Amount    int64     `gorm:"column:amount"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	AccountId int64     `gorm:"column:account_id"`
	Account   Account   `gorm:"foreignKey:account_id;references:id"`
}

func (u *Deposit) TableName() string {
	return "deposits"
}
