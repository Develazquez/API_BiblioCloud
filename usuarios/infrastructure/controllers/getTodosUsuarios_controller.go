package controllers

import (
	"github.com/gin-gonic/gin"

	"biblioteca-api/usuarios/application"
)

// GetTodosUsuariosController obtiene todos los usuarios
type GetTodosUsuariosController struct {
	usecase *application.GetTodosUsuariosUseCase
}

// NewGetTodosUsuariosController crea una nueva instancia
func NewGetTodosUsuariosController(usecase *application.GetTodosUsuariosUseCase) *GetTodosUsuariosController {
	return &GetTodosUsuariosController{usecase: usecase}
}

// Handle maneja la petición HTTP
func (c *GetTodosUsuariosController) Handle(ctx *gin.Context) {
	resultado, err := c.usecase.Execute()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resultado)
}
