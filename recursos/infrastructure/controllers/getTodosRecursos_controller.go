package controllers

import (
	"github.com/gin-gonic/gin"

	"biblioteca-api/recursos/application"
	"biblioteca-api/recursos/domain/entities"
	"biblioteca-api/recursos/infrastructure/dto"
)

type GetTodosRecursosController struct {
	usecase *application.GetTodosRecursosUseCase
}

func NewGetTodosRecursosController(usecase *application.GetTodosRecursosUseCase) *GetTodosRecursosController {
	return &GetTodosRecursosController{usecase: usecase}
}

func (c *GetTodosRecursosController) Handle(ctx *gin.Context) {
	resultado, err := c.usecase.Execute()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	recursos := resultado.([]entities.Recurso)
	ctx.JSON(200, dto.NewRecursoResponseSlice(recursos))
}
