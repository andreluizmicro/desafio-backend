package controller

import (
	"errors"
	"github.com/andreluizmicro/desafio-backend/internal/Application/account"
	"github.com/andreluizmicro/desafio-backend/internal/domain/exception"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	createAccountService *account.CreateAccountService
}

func NewAccountController(createAccountService *account.CreateAccountService) *AccountController {
	return &AccountController{
		createAccountService: createAccountService,
	}
}

func (c *AccountController) Create(ctx *gin.Context) {
	var input account.CreateAccountInputDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	output, err := c.createAccountService.Execute(input)
	if err != nil {
		if errors.Is(err, exception.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, &output)
}
