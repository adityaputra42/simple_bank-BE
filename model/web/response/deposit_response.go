package response

import "time"

type DepositResponse struct {
	ID        int64           `json:"id"`
	Amount    int64           `json:"amount"`
	CreatedAt time.Time       `json:"created_at;autoCreateTime;<-:create"`
	Account   AccountResponse `json:"account"`
}
