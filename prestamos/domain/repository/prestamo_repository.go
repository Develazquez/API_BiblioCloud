package repository

import "biblioteca-api/prestamos/domain/entities"

type PrestamoRepository interface {
	Crear(prestamo *entities.Prestamo) (*entities.Prestamo, error)

	ObtenerPorID(id int) (*entities.Prestamo, error)

	ObtenerTodos() ([]entities.Prestamo, error)

	Actualizar(prestamo *entities.Prestamo) (*entities.Prestamo, error)

	Eliminar(id int) error

	ObtenerPorUsuarioID(usuarioID int) ([]entities.Prestamo, error)

	ObtenerPorRecursoID(recursoID int) ([]entities.Prestamo, error)

	ObtenerPorEstado(estado entities.EstadoPrestamo) ([]entities.Prestamo, error)

	ActualizarEstado(id int, estado entities.EstadoPrestamo) error

	ObtenerActivosPorUsuario(usuarioID int) ([]entities.Prestamo, error)

	ObtenerActivosPorRecurso(recursoID int) ([]entities.Prestamo, error)

	ContarActivosPorUsuario(usuarioID int) (int, error)
}
