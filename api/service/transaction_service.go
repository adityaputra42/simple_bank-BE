package service

import (
	"errors"
	"fmt"
	"simple_bank_solid/api/repository"
	"simple_bank_solid/db"
	"simple_bank_solid/helper"
	"simple_bank_solid/model/domain"
	"simple_bank_solid/model/web/request"
	"simple_bank_solid/model/web/response"

	"gorm.io/gorm"
)

type TransactionService interface {
	Transfer(req request.TransferRequest, userId int64) (response.TransferResponse, error)
	FecthTransferById(TransactionId string) (response.TransferResponse, error)
	FecthAllTransferByUserId(UserId int64) ([]response.TransferResponse, error)
	FecthAllTransfer() ([]response.TransferResponse, error)
	DeleteTransfer(TxId string) error
}

type TransactionServieImpl struct {
	accountRepo     repository.AccountRepository
	transactionRepo repository.TransactionRepository
	entriesRepo     repository.EntriesRepository
	db              *gorm.DB
}

// DeleteTransfer implements TransactionService.
func (t *TransactionServieImpl) DeleteTransfer(TxId string) error {

	transfer, err := t.transactionRepo.FindById(TxId)
	if err != nil {
		return fmt.Errorf("Data transfer not found")
	}

	err = t.transactionRepo.Delete(transfer)
	if err != nil {
		return err
	}
	return nil
}

// FecthAllTransfer implements TransactionService.
func (t *TransactionServieImpl) FecthAllTransfer() ([]response.TransferResponse, error) {
	var listTransfer []response.TransferResponse
	transfers, err := t.transactionRepo.FindAll()
	if err != nil {
		return listTransfer, err
	}
	for _, v := range transfers {
		listTransfer = append(listTransfer, helper.ToTranferRespone(v, v.FromAccount, v.ToAccount))
	}
	return listTransfer, nil
}

// FecthAllTransferByUserId implements TransactionService.
func (t *TransactionServieImpl) FecthAllTransferByUserId(UserId int64) ([]response.TransferResponse, error) {
	var listTransfer []response.TransferResponse
	transfers, err := t.transactionRepo.FindAllbyUserId(int(UserId))
	if err != nil {
		return listTransfer, err
	}
	for _, v := range transfers {
		listTransfer = append(listTransfer, helper.ToTranferRespone(v, v.FromAccount, v.ToAccount))
	}
	return listTransfer, nil
}

// FecthTransferById implements TransactionService.
func (t *TransactionServieImpl) FecthTransferById(TransactionId string) (response.TransferResponse, error) {
	transfer, err := t.transactionRepo.FindById(TransactionId)
	if err != nil {
		return response.TransferResponse{}, err
	}
	return helper.ToTranferRespone(transfer, transfer.FromAccount, transfer.ToAccount), nil
}

// Transfer implements TransactionService.
func (t *TransactionServieImpl) Transfer(req request.TransferRequest, userId int64) (response.TransferResponse, error) {
	var response response.TransferResponse
	err := t.db.Transaction(func(tx *gorm.DB) error {
		fromAccount, fromAccountValid := helper.ValidAccount(tx, req.FromAccountID, req.Currency)
		if fromAccountValid != true {
			return errors.New("From account invalid")
		}
		if fromAccount.UserId != userId {
			err := errors.New("from account doesn't belong to the authenticated user")
			return err
		}

		toAccount, toAccountValid := helper.ValidAccount(tx, req.ToAccountID, req.Currency)
		if toAccountValid != true {
			return errors.New("To account invalid")
		}

		tranferReq := domain.Transaction{
			ID:            helper.Generate("TRX-"),
			Amount:        req.Amount,
			FromAccountId: fromAccount.ID,
			ToAccountId:   toAccount.ID,
			Currency:      req.Currency,
		}
		transfer, err := t.transactionRepo.Create(tranferReq, tx)
		if err != nil {
			return err
		}

		entryFrom := domain.Entries{
			AccountId: fromAccount.ID,
			Amount:    req.Amount,
		}
		_, err = t.entriesRepo.Create(entryFrom, tx)

		if err != nil {
			return err
		}

		entryTo := domain.Entries{
			AccountId: toAccount.ID,
			Amount:    req.Amount,
		}
		_, err = t.entriesRepo.Create(entryTo, tx)

		if err != nil {
			return err
		}

		fromAccount.Balance -= req.Amount

		newFromAccount, err := t.accountRepo.Update(fromAccount, tx)
		if err != nil {
			return err
		}

		toAccount.Balance += req.Amount

		newToAccount, err := t.accountRepo.Update(toAccount, tx)
		if err != nil {
			return err
		}

		response = helper.ToTranferRespone(transfer, newFromAccount, newToAccount)
		return nil

	})
	if err != nil {
		return response, err
	}
	return response, nil
}

func NewTransactionService(
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
