package application

import (
	"errors"
	"biblioteca-api/prestamos/domain/repository"
)

type GetPrestamoPorIDUseCase struct {
	repository repository.PrestamoRepository
}

func NewGetPrestamoPorIDUseCase(repo repository.PrestamoRepository) *GetPrestamoPorIDUseCase {
	return &GetPrestamoPorIDUseCase{repository: repo}
}

func (uc *GetPrestamoPorIDUseCase) Execute(id int) (interface{}, error) {
	if id <= 0 {
		return nil, errors.New("ID inválido")
	}

	return uc.repository.ObtenerPorID(id)
}
