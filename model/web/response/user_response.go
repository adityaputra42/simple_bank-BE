package response

import "time"

type UserResponse struct {
	ID        int64             `json:"id"`
	Username  string            `json:"username"`
	FullName  string            `json:"full_name"`
	Email     string            `json:"email"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Accounts  []AccountResponse `json:"accounts"`
}
