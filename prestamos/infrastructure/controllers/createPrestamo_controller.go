package controllers

import (
	"github.com/gin-gonic/gin"

	"biblioteca-api/prestamos/application"
	"biblioteca-api/prestamos/domain/entities"
)

type CreatePrestamoController struct {
	usecase *application.CreatePrestamoUseCase
}

func NewCreatePrestamoController(usecase *application.CreatePrestamoUseCase) *CreatePrestamoController {
	return &CreatePrestamoController{usecase: usecase}
}

func (c *CreatePrestamoController) Handle(ctx *gin.Context) {
	var prestamo entities.Prestamo
	err := ctx.BindJSON(&prestamo)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resultado, err := c.usecase.Execute(&prestamo)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, resultado)
}
