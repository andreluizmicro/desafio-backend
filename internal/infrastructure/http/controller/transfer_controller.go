package controller

import (
	"errors"
	"github.com/andreluizmicro/desafio-backend/internal/Application/transfer"
	"github.com/andreluizmicro/desafio-backend/internal/domain/exception"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransferController struct {
	createTransferService *transfer.CreateTransferService
}

func NewTransferController(createTransferService *transfer.CreateTransferService) *TransferController {
	return &TransferController{
		createTransferService: createTransferService,
	}
}

func (c *TransferController) Create(ctx *gin.Context) {
	var input transfer.CreateTransferInputDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	output, err := c.createTransferService.Execute(input)
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
