package controller

import (
	"simple_bank_solid/api/service"
	"simple_bank_solid/model/web"
	"simple_bank_solid/model/web/request"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TransactionController interface {
	Transfer(c *fiber.Ctx) error
	FecthTransferById(c *fiber.Ctx) error
	FecthAllTransferByUserId(c *fiber.Ctx) error
	FecthAllTransfer(c *fiber.Ctx) error
}

type TransactionControllerImpl struct {
	transactionService service.TransactionService
}

// FecthAllTransfer implements TransactionController.
func (t TransactionControllerImpl) FecthAllTransfer(c *fiber.Ctx) error {
	results, err := t.transactionService.FecthAllTransfer()
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(web.BaseResponse{
		Status:  20,
		Message: "Success",
		Data:    results,
	})

}

// FecthAllTransferByUserId implements TransactionController.
func (t TransactionControllerImpl) FecthAllTransferByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")
	intValue, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: "Invalid Query Params",
		})
	}
	results, err := t.transactionService.FecthAllTransferByUserId(intValue)

	if err != nil {
		return c.Status(403).JSON(web.BaseResponse{
			Status:  403,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(web.BaseResponse{
		Status:  20,
		Message: "Success",
		Data:    results,
	})

}

// FecthTransferById implements TransactionController.
func (t TransactionControllerImpl) FecthTransferById(c *fiber.Ctx) error {
	txId := c.Params("tx_id")

	results, err := t.transactionService.FecthTransferById(txId)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: "Internal Server Error",
		})
	}
	return c.Status(200).JSON(web.BaseResponse{
		Status:  20,
		Message: "Success",
		Data:    results,
	})
}

// Transfer implements TransactionController.
func (t TransactionControllerImpl) Transfer(c *fiber.Ctx) error {
	req := new(request.TransferRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: "Invalid Message Body",
		})
	}
	result, err := t.transactionService.Transfer(*req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}

	return c.Status(201).JSON(web.BaseResponse{
		Status:  201,
		Message: "Success",
		Data:    result,
	})
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{transactionService: transactionService}
}
