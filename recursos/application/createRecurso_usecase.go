package application

import (
	"database/sql"
	"errors"
	"mime/multipart"

	"biblioteca-api/recursos/domain/entities"
	"biblioteca-api/recursos/domain/repository"
	"biblioteca-api/shared/cloudinary"
)

// CreateRecursoUseCase maneja la creación de recursos
type CreateRecursoUseCase struct {
	repository repository.RecursoRepository
}

// NewCreateRecursoUseCase crea una nueva instancia
func NewCreateRecursoUseCase(repo repository.RecursoRepository) *CreateRecursoUseCase {
	return &CreateRecursoUseCase{repository: repo}
}

// Execute ejecuta el caso de uso de crear recurso
func (uc *CreateRecursoUseCase) Execute(recurso *entities.Recurso, imagen *multipart.FileHeader, audio *multipart.FileHeader) (*entities.Recurso, error) {
	if !recurso.IsValid() {
		return nil, errors.New("recurso no válido: falta el título")
	}

	// 1. Subir Imagen (si existe)
	if imagen != nil {
		imgUrl, err := cloudinary.UploadImage(imagen)
		if err != nil {
			return nil, err
		}
		recurso.ImagenURL = sql.NullString{String: imgUrl, Valid: true}
	}

	// 2. Subir Audio (si existe)
	if audio != nil {
		audioUrl, err := cloudinary.UploadAudio(audio)
		if err != nil {
			return nil, err
		}
		recurso.AudioURL = sql.NullString{String: audioUrl, Valid: true}
	}

	return uc.repository.Crear(recurso)
}
