package application

import (
	"errors"
	"biblioteca-api/recursos/domain/entities"
	"biblioteca-api/recursos/domain/repository"
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
func (uc *CreateRecursoUseCase) Execute(recurso *entities.Recurso) (*entities.Recurso, error) {
	if !recurso.IsValid() {
		return nil, errors.New("recurso no válido: falta el título")
	}

	return uc.repository.Crear(recurso)
}
