package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"biblioteca-api/usuarios/application"
)

// DeleteUsuarioController elimina un usuario
type DeleteUsuarioController struct {
	usecase *application.DeleteUsuarioUseCase
}

// NewDeleteUsuarioController crea una nueva instancia
func NewDeleteUsuarioController(usecase *application.DeleteUsuarioUseCase) *DeleteUsuarioController {
	return &DeleteUsuarioController{usecase: usecase}
}

// Handle maneja la petición HTTP
func (c *DeleteUsuarioController) Handle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.usecase.Execute(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(204)
}
