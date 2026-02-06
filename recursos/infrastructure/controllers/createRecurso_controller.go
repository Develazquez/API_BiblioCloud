package controllers

import (
	"github.com/gin-gonic/gin"

	"biblioteca-api/recursos/application"
	"biblioteca-api/recursos/domain/entities"
)

type CreateRecursoController struct {
	usecase *application.CreateRecursoUseCase
}

func NewCreateRecursoController(usecase *application.CreateRecursoUseCase) *CreateRecursoController {
	return &CreateRecursoController{usecase: usecase}
}

func (c *CreateRecursoController) Handle(ctx *gin.Context) {
	var recurso entities.Recurso
	err := ctx.BindJSON(&recurso)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resultado, err := c.usecase.Execute(&recurso)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, resultado)
}
