package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"biblioteca-api/usuarios/application"
	"biblioteca-api/usuarios/domain/entities"
)

type UpdateUsuarioController struct {
	usecase *application.UpdateUsuarioUseCase
}

func NewUpdateUsuarioController(usecase *application.UpdateUsuarioUseCase) *UpdateUsuarioController {
	return &UpdateUsuarioController{usecase: usecase}
}

func (c *UpdateUsuarioController) Handle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var usuario entities.Usuario
	err := ctx.BindJSON(&usuario)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	usuario.ID = id

	resultado, err := c.usecase.Execute(&usuario)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resultado)
}
