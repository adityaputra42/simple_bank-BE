package request

type AccountRequest struct {
	UserId   int64  `json:"user_id"`
	Currency string `json:"currency"`
}
