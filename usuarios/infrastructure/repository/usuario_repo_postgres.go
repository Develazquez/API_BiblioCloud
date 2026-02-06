package repository

import (
	"database/sql"
	"errors"

	"biblioteca-api/usuarios/domain/entities"
	"biblioteca-api/usuarios/domain/repository"
)

type UsuarioRepositoryPostgres struct {
	db *sql.DB
}

func NewUsuarioRepositoryPostgres(db *sql.DB) repository.UsuarioRepository {
	return &UsuarioRepositoryPostgres{db: db}
}

func (r *UsuarioRepositoryPostgres) Crear(usuario *entities.Usuario) (*entities.Usuario, error) {
	query := `INSERT INTO usuarios (nombre, email, password, estado, cantidad_prestamos_actuales)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := r.db.QueryRow(query, usuario.Nombre, usuario.Email, usuario.Password,
		usuario.Estado, usuario.CantidadPrestamosActuales).Scan(&usuario.ID)

	if err != nil {
		return nil, err
	}

	return usuario, nil
}

func (r *UsuarioRepositoryPostgres) ObtenerPorID(id int) (*entities.Usuario, error) {
	query := `SELECT id, nombre, email, password, estado, cantidad_prestamos_actuales
	          FROM usuarios WHERE id = $1`

	usuario := &entities.Usuario{}
	err := r.db.QueryRow(query, id).Scan(
		&usuario.ID,
		&usuario.Nombre,
		&usuario.Email,
		&usuario.Password,
		&usuario.Estado,
		&usuario.CantidadPrestamosActuales,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("usuario no encontrado")
	}

	if err != nil {
		return nil, err
	}

	return usuario, nil
}

func (r *UsuarioRepositoryPostgres) ObtenerPorEmail(email string) (*entities.Usuario, error) {
	query := `SELECT id, nombre, email, password, estado, cantidad_prestamos_actuales
	          FROM usuarios WHERE email = $1`

	usuario := &entities.Usuario{}
	err := r.db.QueryRow(query, email).Scan(
		&usuario.ID,
		&usuario.Nombre,
		&usuario.Email,
		&usuario.Password,
		&usuario.Estado,
		&usuario.CantidadPrestamosActuales,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return usuario, nil
}

func (r *UsuarioRepositoryPostgres) ObtenerTodos() ([]entities.Usuario, error) {
	query := `SELECT id, nombre, email, password, estado, cantidad_prestamos_actuales
	          FROM usuarios ORDER BY id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usuarios := []entities.Usuario{}
	for rows.Next() {
		usuario := entities.Usuario{}
		err := rows.Scan(
			&usuario.ID,
			&usuario.Nombre,
			&usuario.Email,
			&usuario.Password,
			&usuario.Estado,
			&usuario.CantidadPrestamosActuales,
		)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (r *UsuarioRepositoryPostgres) Actualizar(usuario *entities.Usuario) (*entities.Usuario, error) {
	query := `UPDATE usuarios SET nombre = $1, email = $2, password = $3, estado = $4
	          WHERE id = $5`

	_, err := r.db.Exec(query, usuario.Nombre, usuario.Email, usuario.Password,
		usuario.Estado, usuario.ID)

	if err != nil {
		return nil, err
	}

	return usuario, nil
}

func (r *UsuarioRepositoryPostgres) Eliminar(id int) error {
	query := `DELETE FROM usuarios WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UsuarioRepositoryPostgres) ActualizarEstado(id int, estado entities.EstadoUsuario) error {
	query := `UPDATE usuarios SET estado = $1 WHERE id = $2`
	_, err := r.db.Exec(query, estado, id)
	return err
}

func (r *UsuarioRepositoryPostgres) IncrementarPrestamos(id int) error {
	query := `UPDATE usuarios SET cantidad_prestamos_actuales = cantidad_prestamos_actuales + 1 WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UsuarioRepositoryPostgres) DecrementarPrestamos(id int) error {
	query := `UPDATE usuarios SET cantidad_prestamos_actuales = cantidad_prestamos_actuales - 1 WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UsuarioRepositoryPostgres) ObtenerPorEstado(estado entities.EstadoUsuario) ([]entities.Usuario, error) {
	query := `SELECT id, nombre, email, password, estado, cantidad_prestamos_actuales
	          FROM usuarios WHERE estado = $1`

	rows, err := r.db.Query(query, estado)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usuarios := []entities.Usuario{}
	for rows.Next() {
		usuario := entities.Usuario{}
		err := rows.Scan(
			&usuario.ID,
			&usuario.Nombre,
			&usuario.Email,
			&usuario.Password,
			&usuario.Estado,
			&usuario.CantidadPrestamosActuales,
		)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}
