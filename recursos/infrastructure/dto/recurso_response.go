package dto

import (
	"database/sql"

	"biblioteca-api/recursos/domain/entities"
)

// RecursoResponse representa la respuesta de un recurso
type RecursoResponse struct {
	ID          int    `json:"id"`
	Titulo      string `json:"titulo"`
	Categoria   string `json:"categoria"`
	ImagenURL   string `json:"imagen_url"`
	Estado      string `json:"estado"`
	Descripcion string `json:"descripcion"`
}

// RecursoRequest representa la solicitud para crear o actualizar un recurso
type RecursoRequest struct {
	ID          int    `json:"id"`
	Titulo      string `json:"titulo" binding:"required"`
	Categoria   string `json:"categoria" binding:"required"`
	ImagenURL   string `json:"imagen_url"`
	Estado      string `json:"estado"`
	Descripcion string `json:"descripcion"`
}

// NewRecursoResponse convierte una entidad Recurso a RecursoResponse
func NewRecursoResponse(r *entities.Recurso) RecursoResponse {
	return RecursoResponse{
		ID:          r.ID,
		Titulo:      r.Titulo,
		Categoria:   r.Categoria,
		ImagenURL:   r.GetImagenURL(),
		Estado:      r.GetEstadoString(),
		Descripcion: r.Descripcion,
	}
}

// NewRecursoResponseSlice convierte un slice de Recurso a slice de RecursoResponse
func NewRecursoResponseSlice(recursos []entities.Recurso) []RecursoResponse {
	response := make([]RecursoResponse, len(recursos))
	for i, r := range recursos {
		response[i] = NewRecursoResponse(&r)
	}
	return response
}

// ToEntity convierte un RecursoRequest a una entidad Recurso
func (r *RecursoRequest) ToEntity() *entities.Recurso {
	return &entities.Recurso{
		ID:        r.ID,
		Titulo:    r.Titulo,
		Categoria: r.Categoria,
		ImagenURL: sql.NullString{
			String: r.ImagenURL,
			Valid:  r.ImagenURL != "",
		},
		Estado:      entities.EstadoRecurso(r.Estado),
		Descripcion: r.Descripcion,
	}
}
