package routes

import (
	"simple_bank_solid/config"
	"simple_bank_solid/token"
	"simple_bank_solid/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitServer(config config.Configuration) error {

	err := token.InitTokenMaker(config.SecretKey)
	if err != nil {
		return err
	}

	app := fiber.New()
	RouteInit(app)
	return app.Listen(config.ServerAddress)
}

func RouteInit(app *fiber.App) {

	accountController := InitializeAccountController()
	userController := InitializeUserController()
	depositController := InitializeDepositController()
	transactionController := InitializeTransactionController()
	api_user := app.Group("/api/v1/user")
	{

		api_user.Post("/create", userController.Create)
	}
	api_account := app.Group("/api/v1/account").Use(middleware.AuthMiddleware)
	{

		api_account.Post("/create", accountController.CreateAccount)
		api_account.Post("/deposit", depositController.CreateDeposit)
		api_account.Post("/transfer", transactionController.Transfer)
	}

}
