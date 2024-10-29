package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `gorm:"primaryKey;column:id;autoIncrement"`
	Username  string         `gorm:"column:username"`
	Password  string         `gorm:"column:password"`
	FullName  string         `gorm:"column:full_name"`
	Email     string         `gorm:"column:email"`
	Role      string         `gorm:"column:role"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	Accounts  []Account      `gorm:"foreignKey:user_id;references:id"`
}

func (u *User) TableName() string {
	return "users"
}
