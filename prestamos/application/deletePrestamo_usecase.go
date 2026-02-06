package application

import (
	"errors"
	"biblioteca-api/prestamos/domain/repository"
)

type DeletePrestamoUseCase struct {
	repository repository.PrestamoRepository
}

func NewDeletePrestamoUseCase(repo repository.PrestamoRepository) *DeletePrestamoUseCase {
	return &DeletePrestamoUseCase{repository: repo}
}

func (uc *DeletePrestamoUseCase) Execute(id int) error {
	if id <= 0 {
		return errors.New("ID inválido")
	}

	existente, _ := uc.repository.ObtenerPorID(id)
	if existente == nil {
		return errors.New("préstamo no encontrado")
	}

	return uc.repository.Eliminar(id)
}
