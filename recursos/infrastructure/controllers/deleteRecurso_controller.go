package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"biblioteca-api/recursos/application"
)

type DeleteRecursoController struct {
	usecase *application.DeleteRecursoUseCase
}

func NewDeleteRecursoController(usecase *application.DeleteRecursoUseCase) *DeleteRecursoController {
	return &DeleteRecursoController{usecase: usecase}
}

func (c *DeleteRecursoController) Handle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.usecase.Execute(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(204)
}
