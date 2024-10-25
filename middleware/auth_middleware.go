package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"simple_bank_solid/helper"
	"simple_bank_solid/token"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenMaker := token.GetTokenMaker()
	authorizationHeader := c.Get(helper.GetHeaderKey())
	if len(authorizationHeader) == 0 {
		err := errors.New("authorization header is not provided")
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": err,
		})
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		err := errors.New("invalid authorization header format")
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": err,
		})
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != helper.GetTypeBeare() {
		err := fmt.Errorf("unsupported authorization type %s", authorizationType)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": err,
		})
	}

	accessToken := fields[1]
	payload, err := tokenMaker.VerifyToken(accessToken)
	if err != nil {

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": err,
		})
	}

	c.Locals(helper.GetPayloadKey(), payload)

	return c.Next()

}
