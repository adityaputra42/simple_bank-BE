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

type DepositServie interface {
	CreateDeposit(req request.DepositRequest, userId int64) (response.DepositResponse, error)
	FetchDepositById(DepositId int64) (response.DepositResponse, error)
	FetchAllDeposit() ([]response.DepositResponse, error)
	FetchAllDepositByUserId(userId int64) ([]response.DepositResponse, error)
	Delete(depositId int64) error
}

type DepositServieImpl struct {
	accountRepo repository.AccountRepository
	depositRepo repository.DepositRepository
	entriesRepo repository.EntriesRepository
	db          *gorm.DB
}

// Delete implements DepositServie.
func (d *DepositServieImpl) Delete(depositId int64) error {
	deposit, err := d.depositRepo.FindById(int(depositId))
	if err != nil {
		return fmt.Errorf("Deposit not found")
	}
	err = d.depositRepo.Delete(deposit)
	return err
}

// FetchAllDepositByUserId implements DepositServie.
func (d *DepositServieImpl) FetchAllDepositByUserId(userId int64) ([]response.DepositResponse, error) {
	listDeposit := []response.DepositResponse{}
	deposit, err := d.depositRepo.FindAllbyUser(userId)
	if err != nil {
		return listDeposit, err
	}

	for _, v := range deposit {
		listDeposit = append(listDeposit, helper.ToDepositRespone(v, v.Account))
	}
	return listDeposit, nil

}

// CreateDeposit implements DepositServie.
func (d *DepositServieImpl) CreateDeposit(req request.DepositRequest, userId int64) (response.DepositResponse, error) {
	var response response.DepositResponse
	err := d.db.Transaction(func(tx *gorm.DB) error {
		account, valid := helper.ValidAccount(tx, req.AccountId, req.Currency)

		if valid != true {
			return errors.New("Account not valid")
		}

		if account.UserId != userId {
			return errors.New("Account Invalid Permition")
		}
		depositReq := domain.Deposit{
			Amount: req.Amount, Currency: req.Currency, AccountId: req.AccountId,
		}
		deposit, err := d.depositRepo.Create(depositReq, tx)
		if err != nil {
			return err
		}
		entryReq := domain.Entries{
			AccountId: req.AccountId,
			Amount:    req.Amount,
		}
		_, err = d.entriesRepo.Create(entryReq, tx)

		if err != nil {
			return err
		}

		account.Balance += req.Amount
		newAccount, err := d.accountRepo.Update(account, tx)
		if err != nil {
			return err
		}
		response = helper.ToDepositRespone(deposit, newAccount)

		return nil

	})
	if err != nil {
		return response, err
	}
	return response, nil
}

// FetchAllDeposit implements DepositServie.
func (d *DepositServieImpl) FetchAllDeposit() ([]response.DepositResponse, error) {
	var listDeposit []response.DepositResponse
	deposits, err := d.depositRepo.FindAll()
	if err != nil {
		return listDeposit, err
	}
	for _, v := range deposits {
		listDeposit = append(listDeposit, helper.ToDepositRespone(v, v.Account))
	}
	return listDeposit, nil
}

// FetchDepositById implements DepositServie.
func (d *DepositServieImpl) FetchDepositById(DepositId int64) (response.DepositResponse, error) {
	deposit, err := d.depositRepo.FindById(int(DepositId))
	if err != nil {
		return helper.ToDepositRespone(deposit, deposit.Account), err
	}
	return helper.ToDepositRespone(deposit, deposit.Account), nil

}

func NewDepositService(
	AccountRepo repository.AccountRepository,
	DepositRepo repository.DepositRepository,
	EntriesRepo repository.EntriesRepository,
) DepositServie {
	con := db.GetConnection()
	return &DepositServieImpl{
		accountRepo: AccountRepo,
		depositRepo: DepositRepo,
		entriesRepo: EntriesRepo,
		db:          con,
	}
}
