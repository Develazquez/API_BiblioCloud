package controllers

import (
	"github.com/gin-gonic/gin"

	"biblioteca-api/prestamos/application"
)

type GetTodosPrestamosController struct {
	usecase *application.GetTodosPrestamosUseCase
}

func NewGetTodosPrestamosController(usecase *application.GetTodosPrestamosUseCase) *GetTodosPrestamosController {
	return &GetTodosPrestamosController{usecase: usecase}
}

func (c *GetTodosPrestamosController) Handle(ctx *gin.Context) {
	resultado, err := c.usecase.Execute()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resultado)
}
