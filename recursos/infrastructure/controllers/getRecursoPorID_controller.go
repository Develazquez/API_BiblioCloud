package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"biblioteca-api/recursos/application"
	"biblioteca-api/recursos/domain/entities"
	"biblioteca-api/recursos/infrastructure/dto"
)

type GetRecursoPorIDController struct {
	usecase *application.GetRecursoPorIDUseCase
}

func NewGetRecursoPorIDController(usecase *application.GetRecursoPorIDUseCase) *GetRecursoPorIDController {
	return &GetRecursoPorIDController{usecase: usecase}
}

func (c *GetRecursoPorIDController) Handle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	resultado, err := c.usecase.Execute(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	recurso := resultado.(*entities.Recurso)
	ctx.JSON(200, dto.NewRecursoResponse(recurso))
}
