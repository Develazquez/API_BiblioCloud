package controllers

import (
	"github.com/gin-gonic/gin"

	"biblioteca-api/usuarios/application"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUsuarioController struct {
	usecase *application.LoginUsuarioUseCase
}

func NewLoginUsuarioController(usecase *application.LoginUsuarioUseCase) *LoginUsuarioController {
	return &LoginUsuarioController{usecase: usecase}
}

func (c *LoginUsuarioController) Handle(ctx *gin.Context) {
	var request LoginRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "email y contraseña son requeridos"})
		return
	}

	usuario, err := c.usecase.Execute(request.Email, request.Password)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"mensaje": "login exitoso",
		"usuario": usuario,
	})
}
