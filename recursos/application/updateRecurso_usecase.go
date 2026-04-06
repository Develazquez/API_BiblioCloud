package application

import (
	"database/sql"
	"errors"
	"mime/multipart"

	"biblioteca-api/recursos/domain/entities"
	"biblioteca-api/recursos/domain/repository"
	"biblioteca-api/shared/cloudinary"
)

// UpdateRecursoUseCase actualiza un recurso
type UpdateRecursoUseCase struct {
	repository repository.RecursoRepository
}

// NewUpdateRecursoUseCase crea una nueva instancia
func NewUpdateRecursoUseCase(repo repository.RecursoRepository) *UpdateRecursoUseCase {
	return &UpdateRecursoUseCase{repository: repo}
}

// Execute ejecuta el caso de uso de actualizar recurso
func (uc *UpdateRecursoUseCase) Execute(recurso *entities.Recurso, imagen *multipart.FileHeader, audio *multipart.FileHeader) (*entities.Recurso, error) {
	if recurso.ID <= 0 {
		return nil, errors.New("ID de recurso inválido")
	}

	if !recurso.IsValid() {
		return nil, errors.New("recurso no válido: falta el título")
	}

	// Verificar existencia
	existente, _ := uc.repository.ObtenerPorID(recurso.ID)
	if existente == nil {
		return nil, errors.New("recurso no encontrado")
	}

	// 1. Manejo Imagen
	if imagen != nil {
		imgUrl, err := cloudinary.UploadImage(imagen)
		if err == nil {
			recurso.ImagenURL = sql.NullString{String: imgUrl, Valid: true}
		}
	} else {
		// Retener existente
		recurso.ImagenURL = existente.ImagenURL
	}

	// 2. Manejo Audio
	if audio != nil {
		audioUrl, err := cloudinary.UploadAudio(audio)
		if err == nil {
			recurso.AudioURL = sql.NullString{String: audioUrl, Valid: true}
		}
	} else {
		// Retener existente
		recurso.AudioURL = existente.AudioURL
	}

	return uc.repository.Actualizar(recurso)
}
