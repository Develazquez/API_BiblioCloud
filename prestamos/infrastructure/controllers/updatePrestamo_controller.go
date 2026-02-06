package controllers

import (
	"github.com/gin-gonic/gin"

	"biblioteca-api/prestamos/application"
	"biblioteca-api/prestamos/domain/entities"
)

type UpdatePrestamoController struct {
	usecase *application.UpdatePrestamoUseCase
}

func NewUpdatePrestamoController(usecase *application.UpdatePrestamoUseCase) *UpdatePrestamoController {
	return &UpdatePrestamoController{usecase: usecase}
}

func (c *UpdatePrestamoController) Handle(ctx *gin.Context) {
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

	ctx.JSON(200, resultado)
}
