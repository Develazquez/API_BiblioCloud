package controllers

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"

	"biblioteca-api/recursos/application"
	"biblioteca-api/recursos/domain/entities"
	"biblioteca-api/recursos/infrastructure/dto"
)

type UpdateRecursoController struct {
	usecase *application.UpdateRecursoUseCase
}

func NewUpdateRecursoController(usecase *application.UpdateRecursoUseCase) *UpdateRecursoController {
	return &UpdateRecursoController{usecase: usecase}
}

func (c *UpdateRecursoController) Handle(ctx *gin.Context) {
	// Parsear el payload asegurando limite de peso de peticiones
	if err := ctx.Request.ParseMultipartForm(60 << 20); err != nil {
		ctx.JSON(400, gin.H{"error": "error al procesar datos del formulario"})
		return
	}

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "id de recurso inválido"})
		return
	}

	creadoPor, _ := strconv.Atoi(ctx.PostForm("creado_por"))

	// Mapear campos literales
	recurso := &entities.Recurso{
		ID:          id,
		Titulo:      ctx.PostForm("titulo"),
		Categoria:   ctx.PostForm("categoria"),
		Descripcion: ctx.PostForm("descripcion"),
		Autor:       sql.NullString{String: ctx.PostForm("autor"), Valid: ctx.PostForm("autor") != ""},
		CreadoPor:   sql.NullInt64{Int64: int64(creadoPor), Valid: creadoPor > 0},
	}

	// Extraer Archivos
	imagen, _ := ctx.FormFile("imagen") // Ignora el error, nil = no envió archivo
	audio, _ := ctx.FormFile("audio")

	// Validaciones de tamaño
	if imagen != nil && imagen.Size > 5<<20 { // 5MB
		ctx.JSON(400, gin.H{"error": "la imagen no debe superar los 5MB"})
		return
	}
	if audio != nil && audio.Size > 50<<20 { // 50MB
		ctx.JSON(400, gin.H{"error": "el audio no debe superar los 50MB"})
		return
	}

	// Ejecutar logica de negocio
	resultado, err := c.usecase.Execute(recurso, imagen, audio)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, dto.NewRecursoResponse(resultado))
}
