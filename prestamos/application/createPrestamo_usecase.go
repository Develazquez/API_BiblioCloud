package application

import (
	fcmApp "biblioteca-api/fcm/application"
	"biblioteca-api/prestamos/domain/entities"
	"biblioteca-api/prestamos/domain/repository"
	"errors"
	"strconv"
)

type CreatePrestamoUseCase struct {
	prestamoRepository repository.PrestamoRepository
	sendNotification   *fcmApp.SendNotificationUseCase
}

func NewCreatePrestamoUseCase(repo repository.PrestamoRepository, sendNotification *fcmApp.SendNotificationUseCase) *CreatePrestamoUseCase {
	return &CreatePrestamoUseCase{
		prestamoRepository: repo,
		sendNotification:   sendNotification,
	}
}

func (uc *CreatePrestamoUseCase) Execute(prestamo *entities.Prestamo) (*entities.Prestamo, error) {
	if !prestamo.IsValid() {
		return nil, errors.New("préstamo no válido: datos incompletos")
	}

	creado, err := uc.prestamoRepository.Crear(prestamo)
	if err != nil {
		return nil, err
	}

	if uc.sendNotification != nil {
		go func() {
			data := map[string]string{
				"type":    "LOAN_CREATED",
				"loanId":  strconv.Itoa(creado.ID),
				"bookId":  strconv.Itoa(creado.RecursoID),
				"title":   "Préstamo confirmado",
				"message": "Has tomado prestado un recurso exitosamente.",
			}
			uc.sendNotification.Execute(creado.UsuarioID, data)
		}()
	}

	return creado, nil
}
