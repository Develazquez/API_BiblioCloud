package controllers

import (
	"log"
	"net/http"
	"strconv"

	"biblioteca-api/fcm/application"

	"github.com/gin-gonic/gin"
)

type FCMController struct {
	registerTokenUseCase *application.RegisterTokenUseCase
	removeTokenUseCase   *application.RemoveTokenUseCase
}

func NewFCMController(registerToken *application.RegisterTokenUseCase, removeToken *application.RemoveTokenUseCase) *FCMController {
	return &FCMController{
		registerTokenUseCase: registerToken,
		removeTokenUseCase:   removeToken,
	}
}

type TokenRequest struct {
	Token string `json:"token" binding:"required"`
}

func (c *FCMController) RegisterToken(ctx *gin.Context) {
	// Simulamos obtener el ID del JWT (requiere implementarse el middleware de autenticación)
	userIDStr := ctx.GetHeader("X-User-Id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autorizado o falta header X-User-Id"})
		return
	}

	var req TokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido, se requiere 'token'"})
		return
	}

	if err := c.registerTokenUseCase.Execute(userID, req.Token); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar el token"})
		return
	}

	log.Printf("Token registrado exitosamente para el usuario %d", userID)
	ctx.JSON(http.StatusOK, gin.H{"message": "Token registrado exitosamente"})
}

func (c *FCMController) RemoveToken(ctx *gin.Context) {
	// Simulamos obtener el ID del JWT (requiere implementarse el middleware de autenticación)
	userIDStr := ctx.GetHeader("X-User-Id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autorizado o falta header X-User-Id"})
		return
	}

	var req TokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido, se requiere 'token'"})
		return
	}

	if err := c.removeTokenUseCase.Execute(userID, req.Token); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al remover el token"})
		return
	}

	log.Printf("Token removido exitosamente para el usuario %d", userID)
	ctx.JSON(http.StatusOK, gin.H{"message": "Token removido exitosamente"})
}
