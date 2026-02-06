package repository

import (
	"database/sql"
	"errors"

	"biblioteca-api/recursos/domain/entities"
	"biblioteca-api/recursos/domain/repository"
)

// RecursoRepositoryPostgres implementa la interfaz de repositorio
type RecursoRepositoryPostgres struct {
	db *sql.DB
}

// NewRecursoRepositoryPostgres crea una nueva instancia
func NewRecursoRepositoryPostgres(db *sql.DB) repository.RecursoRepository {
	return &RecursoRepositoryPostgres{db: db}
}

// Crear inserta un nuevo recurso
func (r *RecursoRepositoryPostgres) Crear(recurso *entities.Recurso) (*entities.Recurso, error) {
	query := `INSERT INTO recursos (titulo, categoria, imagen_url, estado, descripcion)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := r.db.QueryRow(query, recurso.Titulo, recurso.Categoria, recurso.ImagenURL,
		recurso.Estado, recurso.Descripcion).Scan(&recurso.ID)

	if err != nil {
		return nil, err
	}

	return recurso, nil
}

// ObtenerPorID obtiene un recurso por ID
func (r *RecursoRepositoryPostgres) ObtenerPorID(id int) (*entities.Recurso, error) {
	query := `SELECT id, titulo, categoria, imagen_url, estado, descripcion
	          FROM recursos WHERE id = $1`

	recurso := &entities.Recurso{}
	err := r.db.QueryRow(query, id).Scan(
		&recurso.ID,
		&recurso.Titulo,
		&recurso.Categoria,
		&recurso.ImagenURL,
		&recurso.Estado,
		&recurso.Descripcion,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("recurso no encontrado")
	}

	if err != nil {
		return nil, err
	}

	return recurso, nil
}

// ObtenerTodos retorna todos los recursos
func (r *RecursoRepositoryPostgres) ObtenerTodos() ([]entities.Recurso, error) {
	query := `SELECT id, titulo, categoria, imagen_url, estado, descripcion
	          FROM recursos ORDER BY id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recursos := []entities.Recurso{}
	for rows.Next() {
		recurso := entities.Recurso{}
		err := rows.Scan(
			&recurso.ID,
			&recurso.Titulo,
			&recurso.Categoria,
			&recurso.ImagenURL,
			&recurso.Estado,
			&recurso.Descripcion,
		)
		if err != nil {
			return nil, err
		}
		recursos = append(recursos, recurso)
	}

	return recursos, nil
}

// Actualizar actualiza un recurso existente
func (r *RecursoRepositoryPostgres) Actualizar(recurso *entities.Recurso) (*entities.Recurso, error) {
	query := `UPDATE recursos SET titulo = $1, categoria = $2, imagen_url = $3, 
	          estado = $4, descripcion = $5 WHERE id = $6`

	_, err := r.db.Exec(query, recurso.Titulo, recurso.Categoria, recurso.ImagenURL,
		recurso.Estado, recurso.Descripcion, recurso.ID)

	if err != nil {
		return nil, err
	}

	return recurso, nil
}

// Eliminar elimina un recurso
func (r *RecursoRepositoryPostgres) Eliminar(id int) error {
	query := `DELETE FROM recursos WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// ObtenerPorCategoria obtiene recursos por categoría
func (r *RecursoRepositoryPostgres) ObtenerPorCategoria(categoria string) ([]entities.Recurso, error) {
	query := `SELECT id, titulo, categoria, imagen_url, estado, descripcion
	          FROM recursos WHERE categoria = $1`

	rows, err := r.db.Query(query, categoria)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recursos := []entities.Recurso{}
	for rows.Next() {
		recurso := entities.Recurso{}
		err := rows.Scan(
			&recurso.ID,
			&recurso.Titulo,
			&recurso.Categoria,
			&recurso.ImagenURL,
			&recurso.Estado,
			&recurso.Descripcion,
		)
		if err != nil {
			return nil, err
		}
		recursos = append(recursos, recurso)
	}

	return recursos, nil
}

// ObtenerPorEstado obtiene recursos por estado
func (r *RecursoRepositoryPostgres) ObtenerPorEstado(estado entities.EstadoRecurso) ([]entities.Recurso, error) {
	query := `SELECT id, titulo, categoria, imagen_url, estado, descripcion
	          FROM recursos WHERE estado = $1`

	rows, err := r.db.Query(query, estado)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recursos := []entities.Recurso{}
	for rows.Next() {
		recurso := entities.Recurso{}
		err := rows.Scan(
			&recurso.ID,
			&recurso.Titulo,
			&recurso.Categoria,
			&recurso.ImagenURL,
			&recurso.Estado,
			&recurso.Descripcion,
		)
		if err != nil {
			return nil, err
		}
		recursos = append(recursos, recurso)
	}

	return recursos, nil
}

// ObtenerPorTitulo busca recursos por título
func (r *RecursoRepositoryPostgres) ObtenerPorTitulo(titulo string) ([]entities.Recurso, error) {
	query := `SELECT id, titulo, categoria, imagen_url, estado, descripcion
	          FROM recursos WHERE titulo ILIKE $1`

	rows, err := r.db.Query(query, "%"+titulo+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recursos := []entities.Recurso{}
	for rows.Next() {
		recurso := entities.Recurso{}
		err := rows.Scan(
			&recurso.ID,
			&recurso.Titulo,
			&recurso.Categoria,
			&recurso.ImagenURL,
			&recurso.Estado,
			&recurso.Descripcion,
		)
		if err != nil {
			return nil, err
		}
		recursos = append(recursos, recurso)
	}

	return recursos, nil
}

// ActualizarEstado actualiza el estado de un recurso
func (r *RecursoRepositoryPostgres) ActualizarEstado(id int, estado entities.EstadoRecurso) error {
	query := `UPDATE recursos SET estado = $1 WHERE id = $2`
	_, err := r.db.Exec(query, estado, id)
	return err
}

// ActualizarImagenURL actualiza la URL de la imagen
func (r *RecursoRepositoryPostgres) ActualizarImagenURL(id int, imagenURL string) error {
	query := `UPDATE recursos SET imagen_url = $1 WHERE id = $2`
	_, err := r.db.Exec(query, imagenURL, id)
	return err
}
