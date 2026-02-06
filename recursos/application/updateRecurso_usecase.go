package application

import (
	"errors"
	"biblioteca-api/recursos/domain/entities"
	"biblioteca-api/recursos/domain/repository"
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
func (uc *UpdateRecursoUseCase) Execute(recurso *entities.Recurso) (*entities.Recurso, error) {
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

	return uc.repository.Actualizar(recurso)
}
