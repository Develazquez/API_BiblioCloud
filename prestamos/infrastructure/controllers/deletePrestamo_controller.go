package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"biblioteca-api/prestamos/application"
)

type DeletePrestamoController struct {
	usecase *application.DeletePrestamoUseCase
}

func NewDeletePrestamoController(usecase *application.DeletePrestamoUseCase) *DeletePrestamoController {
	return &DeletePrestamoController{usecase: usecase}
}

func (c *DeletePrestamoController) Handle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.usecase.Execute(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(204)
}
