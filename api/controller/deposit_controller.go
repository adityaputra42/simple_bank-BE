package controller

import (
	"simple_bank_solid/api/service"
	"simple_bank_solid/model/web"
	"simple_bank_solid/model/web/request"
	"strconv"

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
	req := new(request.DepositRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: "Invalid Message Body",
		})
	}
	response, err := d.depositService.CreateDeposit(*req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}

	return c.Status(201).JSON(web.BaseResponse{
		Status:  201,
		Message: "Success",
		Data:    response,
	})

}

// FetchAllDeposit implements DepositController.
func (d *DepositControllerImpl) FetchAllDeposit(c *fiber.Ctx) error {
	response, err := d.depositService.FetchAllDeposit()
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}

	return c.Status(201).JSON(web.BaseResponse{
		Status:  201,
		Message: "Success",
		Data:    response,
	})

}

// FetchDepositById implements DepositController.
func (d *DepositControllerImpl) FetchDepositById(c *fiber.Ctx) error {
	depositId := c.Params("id")

	id, err := strconv.Atoi(depositId)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	response, err := d.depositService.FetchDepositById(int64(id))
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}

	return c.Status(201).JSON(web.BaseResponse{
		Status:  201,
		Message: "Success",
		Data:    response,
	})

}

func NewDepositController(depositService service.DepositServie) DepositController {
	return &DepositControllerImpl{depositService: depositService}
}
