package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"biblioteca-api/usuarios/application"
)

// GetUsuarioPorIDController obtiene un usuario por ID
type GetUsuarioPorIDController struct {
	usecase *application.GetUsuarioPorIDUseCase
}

// NewGetUsuarioPorIDController crea una nueva instancia
func NewGetUsuarioPorIDController(usecase *application.GetUsuarioPorIDUseCase) *GetUsuarioPorIDController {
	return &GetUsuarioPorIDController{usecase: usecase}
}

// Handle maneja la petición HTTP
func (c *GetUsuarioPorIDController) Handle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	resultado, err := c.usecase.Execute(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resultado)
}
