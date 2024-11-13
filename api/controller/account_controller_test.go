package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"simple_bank_solid/db/mock"
	"simple_bank_solid/helper"
	"simple_bank_solid/middleware"
	"simple_bank_solid/model/domain"
	"simple_bank_solid/token"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccount(t *testing.T) {
	user, _ := randomUser()
	account := randomAccount(user.ID, helper.RandomCurrency())

	testCases := []struct {
		name          string
		accountID     int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(service *mock.MockAccountService)
		checkResponse func(t *testing.T, recorder *http.Response)
	}{
		{
			name:      "Ok",
			accountID: account.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				middleware.AddAuthorization(t, request, tokenMaker, helper.GetTypeBearer(), user.Username, user.ID, time.Minute)
			},
			buildStubs: func(store *mock.MockAccountService) {
				store.EXPECT().FetchAccountById(gomock.Eq(account.ID)).Return(helper.ToAccountResponse(account), sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *http.Response) {
				require.Equal(t, http.StatusOK, recorder.StatusCode) // Ubah ke StatusOK jika tes sukses
				RequireBodyMatchAccount(t, recorder.Body, account)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Inisialisasi token maker
			token.InitTokenMaker(helper.RandomString(32))

			store := mock.NewMockAccountService(ctrl)
			tc.buildStubs(store)

			// Setup Fiber app dan route
			app := fiber.New()
			url := fmt.Sprintf("/accounts/%d", tc.accountID)

			req := httptest.NewRequest(http.MethodGet, url, nil)
			req.Header.Set("Content-Type", "application/json")
			tc.setupAuth(t, req, token.GetTokenMaker())

			controller := NewAccountController(store)
			app.Get("/accounts/:id", controller.FetchAccountById) // Gunakan :id sebagai path param

			// Execute request
			resp, err := app.Test(req)
			require.NoError(t, err)

			tc.checkResponse(t, resp)
		})
	}
}

func randomAccount(userId int64, Currency string) domain.Account {
	return domain.Account{
		ID:       helper.RandomInt(1, 1000),
		UserId:   userId,
		Balance:  helper.RandomBalance(),
		Currency: Currency,
	}
}

func RequireBodyMatchAccount(t *testing.T, body io.Reader, account domain.Account) {
	bodyBytes, err := io.ReadAll(body)
	require.NoError(t, err)

	var responseBody map[string]interface{}
	err = json.Unmarshal(bodyBytes, &responseBody)
	require.NoError(t, err)

	require.Equal(t, account.ID, int64(responseBody["data"].(map[string]interface{})["id"].(float64)))
}
