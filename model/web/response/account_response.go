package response

import "time"

type AccountResponse struct {
	ID        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
