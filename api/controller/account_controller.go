package controller

import (
	"simple_bank_solid/api/service"
	"simple_bank_solid/model/web"
	"simple_bank_solid/model/web/request"
	"strconv"

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
	req := new(request.AccountRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: "Invalid Message Body",
		})
	}
	account, err := a.accountService.CreateAccount(*req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}

	return c.Status(201).JSON(web.BaseResponse{
		Status:  201,
		Message: "Success",
		Data:    account,
	})
}

// DeleteAccount implements AccountController.
func (a *AccountControllerImpl) DeleteAccount(c *fiber.Ctx) error {
	accountId := c.Params("accountId")

	id, err := strconv.Atoi(accountId)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	err = a.accountService.DeleteAccount(int64(id))
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(web.BaseResponse{
		Status:  200,
		Message: "Ok",
	})
}

// FetchAccountById implements AccountController.
func (a *AccountControllerImpl) FetchAccountById(c *fiber.Ctx) error {
	accountId := c.Params("accountId")
	id, err := strconv.Atoi(accountId)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}

	result, err := a.accountService.FetchAccountById(int64(id))
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(web.BaseResponse{
		Status:  200,
		Message: "Ok",
		Data:    result,
	})
}

// FetchAllAccount implements AccountController.
func (a *AccountControllerImpl) FetchAllAccount(c *fiber.Ctx) error {
	result, err := a.accountService.FetchAllAccount()
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(web.BaseResponse{
		Status:  200,
		Message: "Ok",
		Data:    result,
	})
}

// FetchAllAccountByUser implements AccountController.
func (a *AccountControllerImpl) FetchAllAccountByUser(c *fiber.Ctx) error {
	userId := c.Params("userId")

	id, err := strconv.Atoi(userId)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	response, err := a.accountService.FetchAllAccountByUser(int64(id))
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

func NewAccountController(accountService service.AccountService) AccountController {
	return &AccountControllerImpl{accountService: accountService}
}
