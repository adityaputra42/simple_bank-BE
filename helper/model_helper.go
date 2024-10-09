package helper

import (
	"simple_bank_solid/model/domain"
	"simple_bank_solid/model/web/response"
)

func ToAccountResponse(account domain.Account) response.AccountResponse {
	return response.AccountResponse{
		ID:        account.ID,
		UserId:    account.UserId,
		Balance:   account.Balance,
		Currency:  account.Currency,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}

func ToTranferRespomne(transfer domain.Transaction, from domain.Account, to domain.Account) response.TransferResponse {
	return response.TransferResponse{}

}

func ToUserResponse(user domain.User) response.UserResponse {
	var listAccount []response.AccountResponse

	for _, account := range user.Accounts {
		listAccount = append(listAccount, ToAccountResponse(account))
	}
	return response.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
		Accounts: listAccount,
	}
}
