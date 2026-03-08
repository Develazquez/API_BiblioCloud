package controllers

import (
	"github.com/gin-gonic/gin"

	"biblioteca-api/usuarios/application"
	"biblioteca-api/usuarios/domain/entities"
)

type usuarioResponse struct {
	ID                        int                    `json:"id"`
	Nombre                    string                 `json:"nombre"`
	Email                     string                 `json:"email"`
	Estado                    entities.EstadoUsuario `json:"estado"`
	CantidadPrestamosActuales int                    `json:"cantidadPrestamosActuales"`
}

type CreateUsuarioController struct {
	usecase *application.CreateUsuarioUseCase
}

func NewCreateUsuarioController(usecase *application.CreateUsuarioUseCase) *CreateUsuarioController {
	return &CreateUsuarioController{usecase: usecase}
}

func (c *CreateUsuarioController) Handle(ctx *gin.Context) {
	var usuario entities.Usuario
	err := ctx.BindJSON(&usuario)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if usuario.Estado == "" {
		usuario.Estado = entities.EstadoActivo
	}

	resultado, err := c.usecase.Execute(&usuario)
	if err != nil {
		if err.Error() == "el email ya está registrado" {
			ctx.JSON(409, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, usuarioResponse{
		ID:                        resultado.ID,
		Nombre:                    resultado.Nombre,
		Email:                     resultado.Email,
		Estado:                    resultado.Estado,
		CantidadPrestamosActuales: resultado.CantidadPrestamosActuales,
	})
}