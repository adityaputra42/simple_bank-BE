package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"simple_bank_solid/db/mock"
	"simple_bank_solid/helper"
	"simple_bank_solid/middleware"
	"simple_bank_solid/model/web/request"
	"simple_bank_solid/token"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestTransfer(t *testing.T) {
	amount := int64(10)

	user1, _ := randomUser()
	user2, _ := randomUser()
	// user3, _ := randomUser()

	account1 := randomAccount(user1.ID, "IDR")
	account2 := randomAccount(user2.ID, "IDR")
	// account3 := randomAccount(user3.ID, "USD")

	testCases := []struct {
		name          string
		body          fiber.Map
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mock.MockTransactionService)
		checkResponse func(t *testing.T, recorder *http.Response)
	}{
		{
			name: "Ok",
			body: fiber.Map{
				"from_account_id": account1.ID,
				"to_account_id":   account2.ID,
				"amount":          amount,
				"currency":        "IDR",
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {

				middleware.AddAuthorization(t, request, tokenMaker, helper.GetTypeBearer(), user1.Username, user1.ID, time.Minute)
			},
			buildStubs: func(store *mock.MockTransactionService) {

				arg := request.TransferRequest{
					FromAccountID: account1.ID,
					ToAccountID:   account2.ID,
					Amount:        amount,
					Currency:      "IDR",
				}
				store.EXPECT().Transfer(gomock.Eq(arg), gomock.Eq(user1.ID)).Times(1)
			},
			checkResponse: func(t *testing.T, recorder *http.Response) {

				require.Equal(t, http.StatusOK, recorder.StatusCode)

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

			store := mock.NewMockTransactionService(ctrl)
			tc.buildStubs(store)

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			// Setup Fiber app dan route
			app := fiber.New()
			url := "/transfers"

			req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			req.Header.Set("Content-Type", "application/json")
			tc.setupAuth(t, req, token.GetTokenMaker())

			controller := NewTransactionController(store)
			app.Post("/transfers", controller.Transfer) // Gunakan :id sebagai path param

			// Execute request
			resp, err := app.Test(req)
			require.NoError(t, err)

			tc.checkResponse(t, resp)
		})
	}
}
