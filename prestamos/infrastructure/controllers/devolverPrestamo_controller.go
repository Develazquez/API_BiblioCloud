package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"biblioteca-api/prestamos/application"
)

type DevolverPrestamoController struct {
	usecase *application.DevolverPrestamoUseCase
}

func NewDevolverPrestamoController(usecase *application.DevolverPrestamoUseCase) *DevolverPrestamoController {
	return &DevolverPrestamoController{usecase: usecase}
}

func (c *DevolverPrestamoController) Handle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.usecase.Execute(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(200)
}
