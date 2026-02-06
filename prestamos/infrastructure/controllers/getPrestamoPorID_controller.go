package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"biblioteca-api/prestamos/application"
)

type GetPrestamoPorIDController struct {
	usecase *application.GetPrestamoPorIDUseCase
}

func NewGetPrestamoPorIDController(usecase *application.GetPrestamoPorIDUseCase) *GetPrestamoPorIDController {
	return &GetPrestamoPorIDController{usecase: usecase}
}

func (c *GetPrestamoPorIDController) Handle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	resultado, err := c.usecase.Execute(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resultado)
}
