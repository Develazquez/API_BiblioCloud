package repository

import "biblioteca-api/recursos/domain/entities"

// RecursoRepository define el contrato para la persistencia de recursos
type RecursoRepository interface {
	// Crear inserta un nuevo recurso en la base de datos
	Crear(recurso *entities.Recurso) (*entities.Recurso, error)

	// ObtenerPorID obtiene un recurso por su ID
	ObtenerPorID(id int) (*entities.Recurso, error)

	// ObtenerTodos retorna todos los recursos
	ObtenerTodos() ([]entities.Recurso, error)

	// Actualizar actualiza un recurso existente
	Actualizar(recurso *entities.Recurso) (*entities.Recurso, error)

	// Eliminar elimina un recurso por su ID
	Eliminar(id int) error

	// ObtenerPorCategoria obtiene recursos por categoría
	ObtenerPorCategoria(categoria string) ([]entities.Recurso, error)

	// ObtenerPorEstado obtiene recursos por estado
	ObtenerPorEstado(estado entities.EstadoRecurso) ([]entities.Recurso, error)

	// ObtenerPorTitulo busca recursos por título
	ObtenerPorTitulo(titulo string) ([]entities.Recurso, error)

	// ActualizarEstado actualiza el estado de un recurso
	ActualizarEstado(id int, estado entities.EstadoRecurso) error

	// ActualizarImagenURL actualiza la URL de la imagen de un recurso
	ActualizarImagenURL(id int, imagenURL string) error
}
