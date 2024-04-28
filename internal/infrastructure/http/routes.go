package http

import (
	"fmt"
	"github.com/andreluizmicro/desafio-backend/internal/infrastructure/http/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	userController *controller.UserController,
	accountController *controller.AccountController,
	transferController *controller.TransferController,
	port string,
) {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		users := v1.Group("/users")
		users.POST("/", userController.Create)

		accounts := v1.Group("/accounts")
		accounts.POST("/", accountController.Create)
		accounts.POST("/deposit/:user_id", accountController.Deposit)

		transfers := v1.Group("/transfers")
		transfers.POST("/", transferController.Create)
	}

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		return
	}
}
