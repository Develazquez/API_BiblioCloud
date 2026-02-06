package application

import (
	"errors"
	"biblioteca-api/prestamos/domain/entities"
	"biblioteca-api/prestamos/domain/repository"
)

type UpdatePrestamoUseCase struct {
	repository repository.PrestamoRepository
}

func NewUpdatePrestamoUseCase(repo repository.PrestamoRepository) *UpdatePrestamoUseCase {
	return &UpdatePrestamoUseCase{repository: repo}
}

func (uc *UpdatePrestamoUseCase) Execute(prestamo *entities.Prestamo) (*entities.Prestamo, error) {
	if prestamo.ID <= 0 {
		return nil, errors.New("ID de préstamo inválido")
	}

	existente, _ := uc.repository.ObtenerPorID(prestamo.ID)
	if existente == nil {
		return nil, errors.New("préstamo no encontrado")
	}

	return uc.repository.Actualizar(prestamo)
}
