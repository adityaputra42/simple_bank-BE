package role

import (
	"fmt"
	"net/http"
	"simple_bank_solid/model/domain"
	"simple_bank_solid/model/web"

	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	user := c.Locals("CurrentUser").(domain.User)
	fmt.Printf("user role => %s", user.Role)
	if user.Role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(&web.BaseResponse{Status: http.StatusUnauthorized, Message: "You don't have permission to access this resource"})
	}
	return c.Next()
}
