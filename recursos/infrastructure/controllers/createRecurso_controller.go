package controllers

import (
	"database/sql"
	"strconv"

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

	creadoPor, _ := strconv.Atoi(ctx.PostForm("creado_por"))

	recurso := &entities.Recurso{
		Titulo:      ctx.PostForm("titulo"),
		Categoria:   ctx.PostForm("categoria"),
		Descripcion: ctx.PostForm("descripcion"),
		Estado:      entities.EstadoDisponible,
		Autor:       sql.NullString{String: ctx.PostForm("autor"), Valid: ctx.PostForm("autor") != ""},
		CreadoPor:   sql.NullInt64{Int64: int64(creadoPor), Valid: creadoPor > 0},
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
