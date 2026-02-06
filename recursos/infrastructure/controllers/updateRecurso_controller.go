package controllers

import (
	"github.com/gin-gonic/gin"

	"biblioteca-api/recursos/application"
	"biblioteca-api/recursos/infrastructure/dto"
)

type UpdateRecursoController struct {
	usecase *application.UpdateRecursoUseCase
}

func NewUpdateRecursoController(usecase *application.UpdateRecursoUseCase) *UpdateRecursoController {
	return &UpdateRecursoController{usecase: usecase}
}

func (c *UpdateRecursoController) Handle(ctx *gin.Context) {
	var request dto.RecursoRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resultado, err := c.usecase.Execute(request.ToEntity())
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, dto.NewRecursoResponse(resultado))
}
