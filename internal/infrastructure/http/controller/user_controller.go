package controller

import (
	"github.com/andreluizmicro/desafio-backend/internal/Application/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	createUserService *user.CreateUserService
}

func NewUserController(createUserService *user.CreateUserService) *UserController {
	return &UserController{
		createUserService: createUserService,
	}
}

func (c *UserController) Create(ctx *gin.Context) {
	var input user.CreateUserInputDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	output, err := c.createUserService.Execute(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, &output)
}
