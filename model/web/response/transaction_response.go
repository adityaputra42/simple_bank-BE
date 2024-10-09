package response

import "time"

type TransferResponse struct {
	TransactionID string          `json:"transaction_id"`
	From          AccountResponse `json:"from"`
	To            AccountResponse `json:"to"`
	Amount        int64           `json:"amount"`
	Currency      string          `json:"currency"`
	CreatedAt     time.Time       `json:"created_at"`
}
