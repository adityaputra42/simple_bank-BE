package controller

import (
	"simple_bank_solid/api/service"
	"simple_bank_solid/model/web"
	"simple_bank_solid/model/web/request"

	"github.com/gofiber/fiber/v2"
)

type SessionController interface {
	RenewAccessToken(c *fiber.Ctx) error
}

type SessionControllerImpl struct {
	sessionService service.SessionService
}

// RenewAccessToken implements SessionController.
func (s *SessionControllerImpl) RenewAccessToken(c *fiber.Ctx) error {
	req := new(request.RenewAccessTokenRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: "Invalid Message Body",
		})
	}

	accessResponse, err := s.sessionService.RenewAccessToken(*req)
	if err != nil {
		return c.Status(500).JSON(web.BaseResponse{
			Status:  500,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(web.BaseResponse{
		Status:  200,
		Message: "Success",
		Data:    accessResponse,
	})
}

func NewSessionController(sessionService service.SessionService) SessionController {
	return &SessionControllerImpl{sessionService: sessionService}
}
