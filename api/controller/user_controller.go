package controller

import (
	"simple_bank_solid/api/service"
	"simple_bank_solid/helper"
	"simple_bank_solid/model/web"
	"simple_bank_solid/model/web/request"
	"simple_bank_solid/token"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	CreateUser(c *fiber.Ctx) error
	CreateAdmin(c *fiber.Ctx) error
	UpdatePassword(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	FetchUSer(c *fiber.Ctx) error
	FetchAllUSer(c *fiber.Ctx) error
}

type UserControllerImpl struct {
	userService service.UserService
}

// FetchAllUSer implements UserController.
func (u *UserControllerImpl) FetchAllUSer(c *fiber.Ctx) error {
	users, err := u.userService.FecthAllUser()
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	return c.Status(201).JSON(web.BaseResponse{
		Status:  201,
		Message: "Success",
		Data:    users,
	})
}

// CreateAdmin implements UserController.
func (u *UserControllerImpl) CreateAdmin(c *fiber.Ctx) error {
	req := new(request.CreateUser)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: "Invalid Message Body",
		})
	}

	user, err := u.userService.CreateAdmin(*req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: "Internal Server Error",
		})
	}
	return c.Status(201).JSON(web.BaseResponse{
		Status:  201,
		Message: "Success",
		Data:    user,
	})
}

// Create implements UserController.
func (u *UserControllerImpl) CreateUser(c *fiber.Ctx) error {
	req := new(request.CreateUser)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: "Invalid Message Body",
		})
	}

	user, err := u.userService.CreateUser(*req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: "Internal Server Error",
		})
	}
	return c.Status(201).JSON(web.BaseResponse{
		Status:  201,
		Message: "Success",
		Data:    user,
	})
}

// Delete implements UserController.
func (u *UserControllerImpl) Delete(c *fiber.Ctx) error {
	authPayload := c.Locals(helper.GetPayloadKey()).(*token.Payload)
	err := u.userService.Delete(authPayload.Username)
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

// FetchUSer implements UserController.
func (u *UserControllerImpl) FetchUSer(c *fiber.Ctx) error {
	authPayload := c.Locals(helper.GetPayloadKey()).(*token.Payload)

	response, err := u.userService.FecthUser(authPayload.Username)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(web.BaseResponse{
		Status:  200,
		Message: "Success",
		Data:    response,
	})
}

// Login implements UserController.
func (u *UserControllerImpl) Login(c *fiber.Ctx) error {
	req := new(request.LoginUser)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: "Invalid Message Body",
		})
	}

	isLogin, user, err := u.userService.Login(*req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	if !isLogin {
		return c.Status(401).JSON(web.BaseResponse{
			Status:  401,
			Message: "Unauthorized",
		})
	}

	return c.Status(200).JSON(web.BaseResponse{
		Status:  200,
		Message: "Success",
		Data:    user,
	})

}

// UpdatePassword implements UserController.
func (u *UserControllerImpl) UpdatePassword(c *fiber.Ctx) error {
	req := new(request.UpdateUser)
	authPayload := c.Locals(helper.GetPayloadKey()).(*token.Payload)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: "Invalid Message Body",
		})
	}
	user, err := u.userService.UpdatePassword(*req, authPayload.Username)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(web.BaseResponse{
		Status:  200,
		Message: "Success",
		Data:    user,
	})

}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}
