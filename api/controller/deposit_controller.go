package controller

import (
	"simple_bank_solid/api/service"

	"github.com/gofiber/fiber/v2"
)

type DepositController interface {
	CreateDeposit(c *fiber.Ctx) error
	FetchDepositById(c *fiber.Ctx) error
	FetchAllDeposit(c *fiber.Ctx) error
}

type DepositControllerImpl struct {
	depositService service.DepositServie
}

// CreateDeposit implements DepositController.
func (d *DepositControllerImpl) CreateDeposit(c *fiber.Ctx) error {
	panic("unimplemented")
}

// FetchAllDeposit implements DepositController.
func (d *DepositControllerImpl) FetchAllDeposit(c *fiber.Ctx) error {
	panic("unimplemented")
}

// FetchDepositById implements DepositController.
func (d *DepositControllerImpl) FetchDepositById(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewDepositController(depositService service.DepositServie) DepositController {
	return &DepositControllerImpl{depositService: depositService}
}
