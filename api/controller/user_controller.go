package controller

import (
	"simple_bank_solid/api/service"
	"simple_bank_solid/model/web"
	"simple_bank_solid/model/web/request"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Create(c *fiber.Ctx) error
	UpdatePassword(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	FetchUSer(c *fiber.Ctx) error
}

type UserControllerImpl struct {
	userService service.UserService
}

// Create implements UserController.
func (u *UserControllerImpl) Create(c *fiber.Ctx) error {
	req := new(request.CreateUser)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: "Invalid Message Body",
		})
	}

	user, err := u.userService.Create(*req)
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
	panic("unimplemented")
}

// FetchUSer implements UserController.
func (u *UserControllerImpl) FetchUSer(c *fiber.Ctx) error {
	panic("unimplemented")
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

	isLogin, err := u.userService.Login(*req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: "Internal Server Error",
		})
	}

	return c.Status(200).JSON(web.BaseResponse{
		Status:  200,
		Message: "Success",
		Data:    isLogin,
	})

}

// UpdatePassword implements UserController.
func (u *UserControllerImpl) UpdatePassword(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}
