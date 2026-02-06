package application

import (
	"errors"
	"biblioteca-api/prestamos/domain/entities"
	"biblioteca-api/prestamos/domain/repository"
)

type CreatePrestamoUseCase struct {
	prestamoRepository repository.PrestamoRepository
}

func NewCreatePrestamoUseCase(repo repository.PrestamoRepository) *CreatePrestamoUseCase {
	return &CreatePrestamoUseCase{prestamoRepository: repo}
}

func (uc *CreatePrestamoUseCase) Execute(prestamo *entities.Prestamo) (*entities.Prestamo, error) {
	if !prestamo.IsValid() {
		return nil, errors.New("préstamo no válido: datos incompletos")
	}

	return uc.prestamoRepository.Crear(prestamo)
}
