//go:build wireinject
// +build wireinject

package routes

import (
	"simple_bank_solid/api/controller"
	"simple_bank_solid/api/repository"
	"simple_bank_solid/api/service"

	"github.com/google/wire"
)

var UserSet = wire.NewSet(
	repository.NewUserRepository,
	repository.NewAccountRepository,
	service.NewUserService,
	controller.NewUserController,
)

var AccountSet = wire.NewSet(
	repository.NewAccountRepository,
	service.NewAccountService,
	controller.NewAccountController,
)

var DepositSet = wire.NewSet(
	repository.NewAccountRepository,
	repository.NewDepositRepository,
	repository.NewEntriesRepository,
	service.NewDepositService,
	controller.NewDepositController,
)

var TransactionSet = wire.NewSet(
	repository.NewAccountRepository,
	repository.NewTransactionRepository,
	repository.NewEntriesRepository,
	service.NewTransactionService,
	controller.NewTransactionController,
)

func InitializeUserController() controller.UserController {
	wire.Build(
		UserSet,
	)
	return nil
}

func InitializeAccountController() controller.AccountController {
	wire.Build(
		AccountSet,
	)
	return nil
}

func InitializeDepositController() controller.DepositController {
	wire.Build(

		DepositSet,
	)
	return nil
}

func InitializeTransactionController() controller.TransactionController {
	wire.Build(
		TransactionSet,
	)
	return nil
}
