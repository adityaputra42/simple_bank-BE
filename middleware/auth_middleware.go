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
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
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
func AddAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	username string,
	userId int64,
	duration time.Duration,
) {
	token, payload, err := tokenMaker.CreateToken(username, userId, duration)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(helper.GetHeaderKey(), authorizationHeader)
}
