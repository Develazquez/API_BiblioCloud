package controllers

import (
	"github.com/gin-gonic/gin"

	"biblioteca-api/usuarios/application"
	"biblioteca-api/usuarios/domain/entities"
)

// CreateUsuarioController maneja la creación de usuarios
type CreateUsuarioController struct {
	usecase *application.CreateUsuarioUseCase
}

// NewCreateUsuarioController crea una nueva instancia
func NewCreateUsuarioController(usecase *application.CreateUsuarioUseCase) *CreateUsuarioController {
	return &CreateUsuarioController{usecase: usecase}
}

// Handle maneja la petición HTTP
func (c *CreateUsuarioController) Handle(ctx *gin.Context) {
	var usuario entities.Usuario
	err := ctx.BindJSON(&usuario)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Establecer estado por defecto si no se proporciona
	if usuario.Estado == "" {
		usuario.Estado = entities.EstadoActivo
	}

	resultado, err := c.usecase.Execute(&usuario)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, resultado)
}
