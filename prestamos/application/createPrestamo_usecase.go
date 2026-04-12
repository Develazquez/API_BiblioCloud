package application

import (
	fcmApp "biblioteca-api/fcm/application"
	"biblioteca-api/prestamos/domain/entities"
	"biblioteca-api/prestamos/domain/repository"
	"errors"
	"strconv"
	"time"
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
	if prestamo.UsuarioID <= 0 {
		return nil, errors.New("el ID de usuario debe ser mayor a 0")
	}
	if prestamo.RecursoID <= 0 {
		return nil, errors.New("el ID de recurso debe ser mayor a 0")
	}
	if prestamo.FechaLimite.IsZero() {
		return nil, errors.New("la fecha límite de devolución es requerida")
	}
	// Si la fecha límite es anterior a ahora (con un pequeño margen de 1 min por desfase de red/servidor)
	if prestamo.FechaLimite.Before(time.Now().Add(-1 * time.Minute)) {
		return nil, errors.New("la fecha límite no puede ser en el pasado")
	}

	// Asignar valores por defecto si no vienen en el JSON
	if prestamo.FechaInicio.IsZero() {
		prestamo.FechaInicio = time.Now()
	}
	if prestamo.Estado == "" {
		prestamo.Estado = entities.EstadoActivo
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
