package application

import (
	"errors"
	"biblioteca-api/recursos/domain/repository"
)

// GetRecursoPorIDUseCase obtiene un recurso por ID
type GetRecursoPorIDUseCase struct {
	repository repository.RecursoRepository
}

// NewGetRecursoPorIDUseCase crea una nueva instancia
func NewGetRecursoPorIDUseCase(repo repository.RecursoRepository) *GetRecursoPorIDUseCase {
	return &GetRecursoPorIDUseCase{repository: repo}
}

// Execute ejecuta el caso de uso
func (uc *GetRecursoPorIDUseCase) Execute(id int) (interface{}, error) {
	if id <= 0 {
		return nil, errors.New("ID inválido")
	}

	return uc.repository.ObtenerPorID(id)
}
