package application

import (
	"biblioteca-api/recursos/domain/repository"
)

// GetTodosRecursosUseCase obtiene todos los recursos
type GetTodosRecursosUseCase struct {
	repository repository.RecursoRepository
}

// NewGetTodosRecursosUseCase crea una nueva instancia
func NewGetTodosRecursosUseCase(repo repository.RecursoRepository) *GetTodosRecursosUseCase {
	return &GetTodosRecursosUseCase{repository: repo}
}

// Execute ejecuta el caso de uso
func (uc *GetTodosRecursosUseCase) Execute() (interface{}, error) {
	return uc.repository.ObtenerTodos()
}
