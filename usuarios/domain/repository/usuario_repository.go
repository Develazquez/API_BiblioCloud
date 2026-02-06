package repository

import "biblioteca-api/usuarios/domain/entities"

type UsuarioRepository interface {
	Crear(usuario *entities.Usuario) (*entities.Usuario, error)

	ObtenerPorID(id int) (*entities.Usuario, error)

	ObtenerPorEmail(email string) (*entities.Usuario, error)

	ObtenerTodos() ([]entities.Usuario, error)

	Actualizar(usuario *entities.Usuario) (*entities.Usuario, error)

	Eliminar(id int) error

	ActualizarEstado(id int, estado entities.EstadoUsuario) error

	IncrementarPrestamos(id int) error

	DecrementarPrestamos(id int) error

	ObtenerPorEstado(estado entities.EstadoUsuario) ([]entities.Usuario, error)
}
