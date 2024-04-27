package http

import (
	"fmt"
	"github.com/andreluizmicro/desafio-backend/internal/infrastructure/http/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	userController *controller.UserController,
	accountController *controller.AccountController,
	port string,
) {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		users := v1.Group("/users")
		users.POST("/", userController.Create)

		accounts := v1.Group("/accounts")
		accounts.POST("/", accountController.Create)
	}

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		return
	}
}
