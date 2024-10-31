package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"simple_bank_solid/helper"
	"simple_bank_solid/middleware"
	"simple_bank_solid/token"

	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func AddAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	username string,
	userId int64,
	duration time.Duration,
) {
	token, err := tokenMaker.CreateToken(username, userId, duration)
	require.NoError(t, err)
	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(helper.GetHeaderKey(), authorizationHeader)
}

func TestAuthMiddleware(t *testing.T) {
	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		checkResponse func(t *testing.T, resp *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				AddAuthorization(t, request, tokenMaker, helper.GetTypeBearer(), "user", 1, time.Minute)
			},
			checkResponse: func(t *testing.T, resp *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, resp.Code)
			},
		},
		{
			name: "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				// Do nothing, no authorization header
			},
			checkResponse: func(t *testing.T, resp *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, resp.Code)
			},
		},
		{
			name: "UnsupportedAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				AddAuthorization(t, request, tokenMaker, "unsupported", "user", 1, time.Minute)
			},
			checkResponse: func(t *testing.T, resp *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, resp.Code)
			},
		},
		{
			name: "InvalidAuthorizationFormat",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				AddAuthorization(t, request, tokenMaker, "", "user", 1, time.Minute)
			},
			checkResponse: func(t *testing.T, resp *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, resp.Code)
			},
		},
		{
			name: "ExpiredToken",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				AddAuthorization(t, request, tokenMaker, helper.GetTypeBearer(), "user", 1, -time.Minute)
			},
			checkResponse: func(t *testing.T, resp *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, resp.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New()

			// Setup the middleware and route
			app.Get("/auth", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
				return c.Status(http.StatusOK).JSON(fiber.Map{})
			})

			// Create request and response recorder
			req := httptest.NewRequest(http.MethodGet, "/auth", nil)
			resp := httptest.NewRecorder()

			// Apply the setupAuth to configure the request headers
			tc.setupAuth(t, req, token.GetTokenMaker())

			// Perform the test
			app.Test(req, -1)

			// Check the response
			tc.checkResponse(t, resp)
		})
	}
}
