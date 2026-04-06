package controllers

import (
	"github.com/gin-gonic/gin"

	"biblioteca-api/recursos/application"
	"biblioteca-api/recursos/domain/entities"
	"biblioteca-api/recursos/infrastructure/dto"
)

type CreateRecursoController struct {
	usecase *application.CreateRecursoUseCase
}

func NewCreateRecursoController(usecase *application.CreateRecursoUseCase) *CreateRecursoController {
	return &CreateRecursoController{usecase: usecase}
}

func (c *CreateRecursoController) Handle(ctx *gin.Context) {
	if err := ctx.Request.ParseMultipartForm(60 << 20); err != nil {
		ctx.JSON(400, gin.H{"error": "error al procesar datos del formulario"})
		return
	}

	recurso := &entities.Recurso{
		Titulo:      ctx.PostForm("titulo"),
		Categoria:   ctx.PostForm("categoria"),
		Descripcion: ctx.PostForm("descripcion"),
		Estado:      entities.EstadoDisponible,
	}


	imagen, _ := ctx.FormFile("imagen") 
	audio, _ := ctx.FormFile("audio")

	if imagen != nil && imagen.Size > 5<<20 { 
		ctx.JSON(400, gin.H{"error": "la imagen no debe superar los 5MB"})
		return
	}
	if audio != nil && audio.Size > 50<<20 { 
		ctx.JSON(400, gin.H{"error": "el audio no debe superar los 50MB"})
		return
	}

	resultado, err := c.usecase.Execute(recurso, imagen, audio)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, dto.NewRecursoResponse(resultado))
}
