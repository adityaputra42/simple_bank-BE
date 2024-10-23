package request

type DepositRequest struct {
	AccountId int64  `json:"account_id"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
}
