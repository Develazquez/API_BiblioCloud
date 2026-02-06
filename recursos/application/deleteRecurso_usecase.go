package application

import (
	"errors"
	"biblioteca-api/recursos/domain/repository"
)

// DeleteRecursoUseCase elimina un recurso
type DeleteRecursoUseCase struct {
	repository repository.RecursoRepository
}

// NewDeleteRecursoUseCase crea una nueva instancia
func NewDeleteRecursoUseCase(repo repository.RecursoRepository) *DeleteRecursoUseCase {
	return &DeleteRecursoUseCase{repository: repo}
}

// Execute ejecuta el caso de uso de eliminar recurso
func (uc *DeleteRecursoUseCase) Execute(id int) error {
	if id <= 0 {
		return errors.New("ID inválido")
	}

	// Verificar existencia
	existente, _ := uc.repository.ObtenerPorID(id)
	if existente == nil {
		return errors.New("recurso no encontrado")
	}

	return uc.repository.Eliminar(id)
}
