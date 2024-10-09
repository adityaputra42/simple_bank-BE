package domain

import "time"

type Entries struct {
	ID        int64     `gorm:"primaryKey;column:id;autoIncrement"`
	AccountId int64     `gorm:"column:account_id"`
	Amount    int64     `gorm:"column:amount"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	Account   Account   `gorm:"foreignKey:account_id;references:id"`
}

func (u *Entries) TableName() string {
	return "entries"
}
