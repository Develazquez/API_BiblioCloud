package application

import (
	fcmApp "biblioteca-api/fcm/application"
	"biblioteca-api/prestamos/domain/repository"
	"errors"
	"strconv"
)

type DevolverPrestamoUseCase struct {
	repository       repository.PrestamoRepository
	sendNotification *fcmApp.SendNotificationUseCase
}

func NewDevolverPrestamoUseCase(repo repository.PrestamoRepository, sendNotification *fcmApp.SendNotificationUseCase) *DevolverPrestamoUseCase {
	return &DevolverPrestamoUseCase{
		repository:       repo,
		sendNotification: sendNotification,
	}
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
	if err != nil {
		return err
	}

	if uc.sendNotification != nil {
		go func() {
			data := map[string]string{
				"type":    "LOAN_RETURNED",
				"loanId":  strconv.Itoa(prestamo.ID),
				"bookId":  strconv.Itoa(prestamo.RecursoID),
				"title":   "Préstamo devuelto",
				"message": "Has devuelto el recurso exitosamente.",
			}
			uc.sendNotification.Execute(prestamo.UsuarioID, data)
		}()
	}

	return nil
}
