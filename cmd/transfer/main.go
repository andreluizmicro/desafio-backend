package main

import (
	"database/sql"
	"github.com/andreluizmicro/desafio-backend/config"
	accountService "github.com/andreluizmicro/desafio-backend/internal/Application/account"
	"github.com/andreluizmicro/desafio-backend/internal/Application/transfer"
	userService "github.com/andreluizmicro/desafio-backend/internal/Application/user"
	"github.com/andreluizmicro/desafio-backend/internal/infrastructure/gateway"
	"github.com/andreluizmicro/desafio-backend/internal/infrastructure/http"
	"github.com/andreluizmicro/desafio-backend/internal/infrastructure/http/controller"
	accountRepository "github.com/andreluizmicro/desafio-backend/internal/infrastructure/repository/account"
	transferRepository "github.com/andreluizmicro/desafio-backend/internal/infrastructure/repository/transfer"
	userRepository "github.com/andreluizmicro/desafio-backend/internal/infrastructure/repository/user"
	"github.com/andreluizmicro/desafio-backend/pkg/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := config.LoadConfig("../")
	if err != nil {
		panic(err)
	}

	db, err := database.NewConnection(cfg)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	createUserRepository := userRepository.NewUserRepository(db)
	createAccountRepository := accountRepository.NewAccountRepository(db)
	createTransferRepository := transferRepository.NewTransferRepository(db)

	authGateway := gateway.NewAuthorizationGateway()
	notifyGateway := gateway.NewNotificationGateway()

	createUserService := userService.NewCreateUserService(createUserRepository)
	createAccountService := accountService.NewCreateAccountService(createAccountRepository, createUserRepository)

	createTransferService := transfer.NewCreateTransferService(
		createAccountRepository,
		createTransferRepository,
		authGateway,
		notifyGateway,
	)

	userController := controller.NewUserController(createUserService)
	accountController := controller.NewAccountController(createAccountService)
	transferController := controller.NewTransferController(createTransferService)

	http.InitRoutes(userController, accountController, transferController, cfg.WebServerPort)
}
