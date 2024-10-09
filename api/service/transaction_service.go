package service

import (
	"simple_bank_solid/api/repository"
	"simple_bank_solid/db"
	"simple_bank_solid/model/web/request"
	"simple_bank_solid/model/web/response"

	"gorm.io/gorm"
)

type TransactionService interface {
	Transfer(req request.TransferRequest) (response.TransferResponse, error)
	FecthTransferById(TransactionId string) (response.TransferResponse, error)
	FecthAllTransferByUserId(UserId string) []response.TransferResponse
	FecthAllTransfer() []response.TransferResponse
}

type TransactionServieImpl struct {
	accountRepo     repository.AccountRepository
	transactionRepo repository.TransactionRepository
	entriesRepo     repository.EntriesRepository
	db              *gorm.DB
}

// FecthAllTransfer implements TransactionService.
func (t *TransactionServieImpl) FecthAllTransfer() []response.TransferResponse {
	panic("unimplemented")
}

// FecthAllTransferByUserId implements TransactionService.
func (t *TransactionServieImpl) FecthAllTransferByUserId(UserId string) []response.TransferResponse {
	panic("unimplemented")
}

// FecthTransferById implements TransactionService.
func (t *TransactionServieImpl) FecthTransferById(TransactionId string) (response.TransferResponse, error) {
	panic("unimplemented")
}

// Transfer implements TransactionService.
func (t *TransactionServieImpl) Transfer(req request.TransferRequest) (response.TransferResponse, error) {
	panic("unimplemented")
}

func NewTranserService(
	AccountRepo repository.AccountRepository,
	TransactionRepo repository.TransactionRepository,
	EntriesRepo repository.EntriesRepository,
) TransactionService {
	con := db.GetConnection()
	return &TransactionServieImpl{
		accountRepo:     AccountRepo,
		transactionRepo: TransactionRepo,
		entriesRepo:     EntriesRepo,
		db:              con,
	}
}
