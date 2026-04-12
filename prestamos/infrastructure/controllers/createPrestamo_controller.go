package controllers

import (
	"log"

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
		log.Printf("[ERROR] Error al parsear JSON de préstamo: %v", err)
		ctx.JSON(400, gin.H{"error": "JSON inválido o campos faltantes: " + err.Error()})
		return
	}

	resultado, err := c.usecase.Execute(&prestamo)
	if err != nil {
		log.Printf("[ERROR] Error en el caso de uso de préstamo: %v", err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, resultado)
}
