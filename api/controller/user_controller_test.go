package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"simple_bank_solid/db/mock"
	"simple_bank_solid/helper"
	"simple_bank_solid/model/web/request"
	"simple_bank_solid/model/web/response"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	user, password := randomUser()

	testCases := []struct {
		name          string
		body          fiber.Map
		buildStubs    func(service *mock.MockUserService)
		checkResponse func(t *testing.T, recorder *http.Response)
	}{
		{
			name: "Ok",
			body: fiber.Map{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(service *mock.MockUserService) {
				arg := request.CreateUser{
					Username: user.Username,
					FullName: user.FullName,
					Email:    user.Email,
					Password: password,
				}
				service.EXPECT().CreateUser(arg).Times(1).Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *http.Response) {
				require.Equal(t, http.StatusCreated, recorder.StatusCode)
				RequireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "InternalError",
			body: fiber.Map{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(service *mock.MockUserService) {

				service.EXPECT().CreateUser(gomock.Any()).Times(1).Return(response.UserResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *http.Response) {
				require.Equal(t, http.StatusInternalServerError, recorder.StatusCode)

			},
		},
		{
			name: "DuplicateUsername",
			body: fiber.Map{
				"username":  user.Username,
				"password":  password,
				"full_name": user.Username,
				"email":     user.Email,
			},
			buildStubs: func(service *mock.MockUserService) {

				service.EXPECT().CreateUser(gomock.Any()).Times(1).Return(response.UserResponse{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *http.Response) {
				require.Equal(t, http.StatusInternalServerError, recorder.StatusCode)

			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock.NewMockUserService(ctrl)
			tc.buildStubs(store)

			app := fiber.New()
			reqJson, err := json.Marshal(tc.body)
			require.NoError(t, err)

			controller := NewUserController(store)
			app.Post("/users", controller.CreateUser)

			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(reqJson))
			req.Header.Set("Content-Type", "application/json")

			// Execute request
			resp, err := app.Test(req)
			require.NoError(t, err)

			tc.checkResponse(t, resp)
		})
	}
}

func randomUser() (user response.UserResponse, password string) {
	password = helper.RandomString(6)

	user = response.UserResponse{
		ID:       helper.RandomInt(1, 1000),
		Username: helper.RandomOwner(),
		FullName: helper.RandomOwner(),
		Email:    helper.RandomEmail(),
	}
	return user, password
}

func RequireBodyMatchUser(t *testing.T, body io.Reader, user response.UserResponse) {
	bodyBytes, err := io.ReadAll(body)
	require.NoError(t, err)

	var responseBody map[string]interface{}
	err = json.Unmarshal(bodyBytes, &responseBody)
	require.NoError(t, err)

	fmt.Println("response body user => ", responseBody)
	fmt.Println("got user => ", responseBody["data"])

	require.Equal(t, user.Email, responseBody["data"].(map[string]interface{})["email"])
}
