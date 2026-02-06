package application

import (
	"biblioteca-api/prestamos/domain/repository"
)

type GetTodosPrestamosUseCase struct {
	repository repository.PrestamoRepository
}

func NewGetTodosPrestamosUseCase(repo repository.PrestamoRepository) *GetTodosPrestamosUseCase {
	return &GetTodosPrestamosUseCase{repository: repo}
}

func (uc *GetTodosPrestamosUseCase) Execute() (interface{}, error) {
	return uc.repository.ObtenerTodos()
}
