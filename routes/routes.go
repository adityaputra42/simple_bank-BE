package routes

import (
	"simple_bank_solid/config"
	"simple_bank_solid/middleware"
	"simple_bank_solid/middleware/role"
	"simple_bank_solid/token"

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
	api := app.Group("/api/v1")
	{
		api.Post("/register", userController.CreateUser)
		api.Post("/login", userController.Login)
		api.Post("/admin/register", userController.CreateAdmin)
	}

	api_user := api.Group("/users").Use(middleware.AuthMiddleware, role.MemberAuth)
	{
		api_user.Get("/me", userController.FetchUSer)
		api_user.Get("/change_password", userController.UpdatePassword)

		api_user.Post("/accounts/create", accountController.CreateAccount)
		api_user.Get("/accounts", accountController.FetchAllAccountByUser)
		api_user.Get("/accounts/:account_id", accountController.FetchAccountById)
		api_user.Get("/accounts/delete/:account_id", accountController.DeleteAccount)

		api_user.Post("/deposit", depositController.CreateDeposit)
		api_user.Get("/deposit", depositController.FetchAllDeposit)
		api_user.Get("/deposit/:id", depositController.FetchDepositById)
		api_user.Get("/deposit/delete/:id", depositController.Delete)

		api_user.Post("/transfer", transactionController.Transfer)
		api_user.Get("/transfer", transactionController.FecthAllTransferByUserId)
		api_user.Get("/transfer/:tx_id", transactionController.FecthTransferById)
		api_user.Get("/transfer/delete/:tx_id", transactionController.DeleteTransfer)

	}

	api_admin := api.Group("/admin").Use(middleware.AuthMiddleware, role.AdminAuth)
	{
		api_admin.Get("/me", userController.FetchUSer)
		api_admin.Get("/accounts", accountController.FetchAllAccount)
		api_admin.Get("/deposit", depositController.FetchAllDeposit)
		api_admin.Delete("/deposit", depositController.Delete)
		api_admin.Get("/transfer", transactionController.FecthAllTransfer)
		api_admin.Get("/users", userController.FetchAllUSer)
		api_admin.Get("/users/delete/:id", userController.Delete)
	}

}
