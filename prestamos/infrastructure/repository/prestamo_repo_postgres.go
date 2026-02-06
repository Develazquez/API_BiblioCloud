package repository

import (
	"database/sql"
	"errors"

	"biblioteca-api/prestamos/domain/entities"
	"biblioteca-api/prestamos/domain/repository"
)

type PrestamoRepositoryPostgres struct {
	db *sql.DB
}

func NewPrestamoRepositoryPostgres(db *sql.DB) repository.PrestamoRepository {
	return &PrestamoRepositoryPostgres{db: db}
}

// Crear inserta un nuevo préstamo
func (r *PrestamoRepositoryPostgres) Crear(prestamo *entities.Prestamo) (*entities.Prestamo, error) {
	query := `INSERT INTO prestamos (usuario_id, recurso_id, fecha_inicio, fecha_limite, estado)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := r.db.QueryRow(query, prestamo.UsuarioID, prestamo.RecursoID, prestamo.FechaInicio,
		prestamo.FechaLimite, prestamo.Estado).Scan(&prestamo.ID)

	if err != nil {
		return nil, err
	}

	return prestamo, nil
}

// ObtenerPorID obtiene un préstamo por ID
func (r *PrestamoRepositoryPostgres) ObtenerPorID(id int) (*entities.Prestamo, error) {
	query := `SELECT id, usuario_id, recurso_id, fecha_inicio, fecha_limite, fecha_devolucion, estado
	          FROM prestamos WHERE id = $1`

	prestamo := &entities.Prestamo{}
	err := r.db.QueryRow(query, id).Scan(
		&prestamo.ID,
		&prestamo.UsuarioID,
		&prestamo.RecursoID,
		&prestamo.FechaInicio,
		&prestamo.FechaLimite,
		&prestamo.FechaDevolucion,
		&prestamo.Estado,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("préstamo no encontrado")
	}

	if err != nil {
		return nil, err
	}

	return prestamo, nil
}

// ObtenerTodos retorna todos los préstamos
func (r *PrestamoRepositoryPostgres) ObtenerTodos() ([]entities.Prestamo, error) {
	query := `SELECT id, usuario_id, recurso_id, fecha_inicio, fecha_limite, fecha_devolucion, estado
	          FROM prestamos ORDER BY id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prestamos := []entities.Prestamo{}
	for rows.Next() {
		prestamo := entities.Prestamo{}
		err := rows.Scan(
			&prestamo.ID,
			&prestamo.UsuarioID,
			&prestamo.RecursoID,
			&prestamo.FechaInicio,
			&prestamo.FechaLimite,
			&prestamo.FechaDevolucion,
			&prestamo.Estado,
		)
		if err != nil {
			return nil, err
		}
		prestamos = append(prestamos, prestamo)
	}

	return prestamos, nil
}

// Actualizar actualiza un préstamo existente
func (r *PrestamoRepositoryPostgres) Actualizar(prestamo *entities.Prestamo) (*entities.Prestamo, error) {
	query := `UPDATE prestamos SET usuario_id = $1, recurso_id = $2, fecha_inicio = $3,
	          fecha_limite = $4, fecha_devolucion = $5, estado = $6 WHERE id = $7`

	_, err := r.db.Exec(query, prestamo.UsuarioID, prestamo.RecursoID, prestamo.FechaInicio,
		prestamo.FechaLimite, prestamo.FechaDevolucion, prestamo.Estado, prestamo.ID)

	if err != nil {
		return nil, err
	}

	return prestamo, nil
}

// Eliminar elimina un préstamo
func (r *PrestamoRepositoryPostgres) Eliminar(id int) error {
	query := `DELETE FROM prestamos WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// ObtenerPorUsuarioID obtiene los préstamos de un usuario
func (r *PrestamoRepositoryPostgres) ObtenerPorUsuarioID(usuarioID int) ([]entities.Prestamo, error) {
	query := `SELECT id, usuario_id, recurso_id, fecha_inicio, fecha_limite, fecha_devolucion, estado
	          FROM prestamos WHERE usuario_id = $1`

	rows, err := r.db.Query(query, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prestamos := []entities.Prestamo{}
	for rows.Next() {
		prestamo := entities.Prestamo{}
		err := rows.Scan(
			&prestamo.ID,
			&prestamo.UsuarioID,
			&prestamo.RecursoID,
			&prestamo.FechaInicio,
			&prestamo.FechaLimite,
			&prestamo.FechaDevolucion,
			&prestamo.Estado,
		)
		if err != nil {
			return nil, err
		}
		prestamos = append(prestamos, prestamo)
	}

	return prestamos, nil
}

// ObtenerPorRecursoID obtiene los préstamos de un recurso
func (r *PrestamoRepositoryPostgres) ObtenerPorRecursoID(recursoID int) ([]entities.Prestamo, error) {
	query := `SELECT id, usuario_id, recurso_id, fecha_inicio, fecha_limite, fecha_devolucion, estado
	          FROM prestamos WHERE recurso_id = $1`

	rows, err := r.db.Query(query, recursoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prestamos := []entities.Prestamo{}
	for rows.Next() {
		prestamo := entities.Prestamo{}
		err := rows.Scan(
			&prestamo.ID,
			&prestamo.UsuarioID,
			&prestamo.RecursoID,
			&prestamo.FechaInicio,
			&prestamo.FechaLimite,
			&prestamo.FechaDevolucion,
			&prestamo.Estado,
		)
		if err != nil {
			return nil, err
		}
		prestamos = append(prestamos, prestamo)
	}

	return prestamos, nil
}

func (r *PrestamoRepositoryPostgres) ObtenerPorEstado(estado entities.EstadoPrestamo) ([]entities.Prestamo, error) {
	query := `SELECT id, usuario_id, recurso_id, fecha_inicio, fecha_limite, fecha_devolucion, estado
	          FROM prestamos WHERE estado = $1`

	rows, err := r.db.Query(query, estado)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prestamos := []entities.Prestamo{}
	for rows.Next() {
		prestamo := entities.Prestamo{}
		err := rows.Scan(
			&prestamo.ID,
			&prestamo.UsuarioID,
			&prestamo.RecursoID,
			&prestamo.FechaInicio,
			&prestamo.FechaLimite,
			&prestamo.FechaDevolucion,
			&prestamo.Estado,
		)
		if err != nil {
			return nil, err
		}
		prestamos = append(prestamos, prestamo)
	}

	return prestamos, nil
}

func (r *PrestamoRepositoryPostgres) ActualizarEstado(id int, estado entities.EstadoPrestamo) error {
	query := `UPDATE prestamos SET estado = $1 WHERE id = $2`
	_, err := r.db.Exec(query, estado, id)
	return err
}

func (r *PrestamoRepositoryPostgres) ObtenerActivosPorUsuario(usuarioID int) ([]entities.Prestamo, error) {
	query := `SELECT id, usuario_id, recurso_id, fecha_inicio, fecha_limite, fecha_devolucion, estado
	          FROM prestamos WHERE usuario_id = $1 AND estado = 'ACTIVO'`

	rows, err := r.db.Query(query, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prestamos := []entities.Prestamo{}
	for rows.Next() {
		prestamo := entities.Prestamo{}
		err := rows.Scan(
			&prestamo.ID,
			&prestamo.UsuarioID,
			&prestamo.RecursoID,
			&prestamo.FechaInicio,
			&prestamo.FechaLimite,
			&prestamo.FechaDevolucion,
			&prestamo.Estado,
		)
		if err != nil {
			return nil, err
		}
		prestamos = append(prestamos, prestamo)
	}

	return prestamos, nil
}

func (r *PrestamoRepositoryPostgres) ObtenerActivosPorRecurso(recursoID int) ([]entities.Prestamo, error) {
	query := `SELECT id, usuario_id, recurso_id, fecha_inicio, fecha_limite, fecha_devolucion, estado
	          FROM prestamos WHERE recurso_id = $1 AND estado = 'ACTIVO'`

	rows, err := r.db.Query(query, recursoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prestamos := []entities.Prestamo{}
	for rows.Next() {
		prestamo := entities.Prestamo{}
		err := rows.Scan(
			&prestamo.ID,
			&prestamo.UsuarioID,
			&prestamo.RecursoID,
			&prestamo.FechaInicio,
			&prestamo.FechaLimite,
			&prestamo.FechaDevolucion,
			&prestamo.Estado,
		)
		if err != nil {
			return nil, err
		}
		prestamos = append(prestamos, prestamo)
	}

	return prestamos, nil
}

func (r *PrestamoRepositoryPostgres) ContarActivosPorUsuario(usuarioID int) (int, error) {
	query := `SELECT COUNT(*) FROM prestamos WHERE usuario_id = $1 AND estado = 'ACTIVO'`

	var count int
	err := r.db.QueryRow(query, usuarioID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
