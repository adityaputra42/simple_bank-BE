package service

import (
	"simple_bank_solid/api/repository"
	"simple_bank_solid/db"
	"simple_bank_solid/model/web/request"
	"simple_bank_solid/model/web/response"

	"gorm.io/gorm"
)

type DepositServie interface {
	CreateDeposit(req request.DepositRequest) (response.DepositResponse, error)
	FetchDepositById(DepositId int64) (response.DepositResponse, error)
	FetchAllDeposit() []response.DepositResponse
}

type DepositServieImpl struct {
	accountRepo repository.AccountRepository
	depositRepo repository.DepositRepository
	entriesRepo repository.EntriesRepository
	db          *gorm.DB
}

// CreateDeposit implements DepositServie.
func (d *DepositServieImpl) CreateDeposit(req request.DepositRequest) (response.DepositResponse, error) {
	panic("unimplemented")
}

// FetchAllDeposit implements DepositServie.
func (d *DepositServieImpl) FetchAllDeposit() []response.DepositResponse {
	panic("unimplemented")
}

// FetchDepositById implements DepositServie.
func (d *DepositServieImpl) FetchDepositById(DepositId int64) (response.DepositResponse, error) {
	panic("unimplemented")
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
