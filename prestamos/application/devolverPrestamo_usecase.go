package application

import (
	"errors"
	"biblioteca-api/prestamos/domain/repository"
)

type DevolverPrestamoUseCase struct {
	repository repository.PrestamoRepository
}

func NewDevolverPrestamoUseCase(repo repository.PrestamoRepository) *DevolverPrestamoUseCase {
	return &DevolverPrestamoUseCase{repository: repo}
}

func (uc *DevolverPrestamoUseCase) Execute(id int) error {
	if id <= 0 {
		return errors.New("ID inválido")
	}

	prestamo, err := uc.repository.ObtenerPorID(id)
	if err != nil || prestamo == nil {
		return errors.New("préstamo no encontrado")
	}

	if !prestamo.IsActivo() {
		return errors.New("el préstamo ya ha sido devuelto")
	}

	prestamo.Devolver()
	_, err = uc.repository.Actualizar(prestamo)
	return err
}
