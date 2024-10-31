package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"simple_bank_solid/db"
	"simple_bank_solid/helper"
	"simple_bank_solid/model/domain"
	"simple_bank_solid/model/web"
	"simple_bank_solid/token"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenMaker := token.GetTokenMaker()
	authorizationHeader := c.Get(helper.GetHeaderKey())
	if len(authorizationHeader) == 0 {
		err := errors.New("authorization header is not provided")
		return c.Status(http.StatusUnauthorized).JSON(web.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		err := errors.New("invalid authorization header format")
		return c.Status(http.StatusUnauthorized).JSON(web.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != helper.GetTypeBearer() {
		err := fmt.Errorf("unsupported authorization type %s", authorizationType)
		return c.Status(http.StatusUnauthorized).JSON(web.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
	}

	accessToken := fields[1]
	payload, err := tokenMaker.VerifyToken(accessToken)
	if err != nil {

		return c.Status(http.StatusUnauthorized).JSON(web.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
	}
	c.Locals(helper.GetPayloadKey(), payload)

	db := db.GetConnection()
	user := domain.User{}

	err = db.Model(&domain.User{}).Take(&user, "id = ?", payload.UserId).Error

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(web.BaseResponse{
			Status:  http.StatusNotFound,
			Message: "User not found",
		})
	}
	c.Locals("CurrentUser", user)
	return c.Next()

}
