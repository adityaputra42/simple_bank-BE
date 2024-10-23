package controller

import (
	"simple_bank_solid/api/service"

	"github.com/gofiber/fiber/v2"
)

type AccountController interface {
	CreateAccount(c *fiber.Ctx) error
	DeleteAccount(c *fiber.Ctx) error
	FetchAccountById(c *fiber.Ctx) error
	FetchAllAccountByUser(c *fiber.Ctx) error
	FetchAllAccount(c *fiber.Ctx) error
}

type AccountControllerImpl struct {
	accountService service.AccountService
}

// CreateAccount implements AccountController.
func (a *AccountControllerImpl) CreateAccount(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteAccount implements AccountController.
func (a *AccountControllerImpl) DeleteAccount(c *fiber.Ctx) error {
	panic("unimplemented")
}

// FetchAccountById implements AccountController.
func (a *AccountControllerImpl) FetchAccountById(c *fiber.Ctx) error {
	panic("unimplemented")
}

// FetchAllAccount implements AccountController.
func (a *AccountControllerImpl) FetchAllAccount(c *fiber.Ctx) error {
	panic("unimplemented")
}

// FetchAllAccountByUser implements AccountController.
func (a *AccountControllerImpl) FetchAllAccountByUser(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewAccountController(accountService service.AccountService) AccountController {
	return &AccountControllerImpl{accountService: accountService}
}
