package domain

import (
	"time"
)

type UserSessions struct {
	ID           string    `gorm:"primaryKey;column:id"`
	UserId       int64     `gorm:"primaryKey;column:user_id"`
	RefreshToken string    `gorm:"column:refresh_token"`
	UserAgent    string    `gorm:"column:user_agent"`
	ClientIp     string    `gorm:"column:client_ip"`
	IsBlocked    bool      `gorm:"column:is_blocked"`
	ExpiredAt    time.Time `gorm:"column:expired_at"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	User         User      `gorm:"foreignKey:user_id;references:id"`
}

func (u *UserSessions) TableName() string {
	return "user_sessions"
}
