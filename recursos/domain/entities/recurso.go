package entities

import "database/sql"

// EstadoRecurso representa los estados posibles de un recurso
type EstadoRecurso string

const (
	EstadoDisponible   EstadoRecurso = "DISPONIBLE"
	EstadoPrestado     EstadoRecurso = "PRESTADO"
	EstadoNoDisponible EstadoRecurso = "NO_DISPONIBLE"
)

// Recurso representa la entidad de recurso en el dominio
type Recurso struct {
	ID          int
	Titulo      string
	Categoria   string
	ImagenURL   sql.NullString
	AudioURL    sql.NullString // NUEVO campo
	Autor       sql.NullString // NUEVO campo
	CreadoPor   sql.NullInt64  // NUEVO campo
	Estado      EstadoRecurso
	Descripcion string
}

// NewRecurso crea una nueva instancia de Recurso
func NewRecurso(titulo, categoria, imagenURL, audioURL, descripcion, autor string, creadoPor int) *Recurso {
	var creadoPorVal sql.NullInt64
	if creadoPor > 0 {
		creadoPorVal = sql.NullInt64{Int64: int64(creadoPor), Valid: true}
	} else {
		creadoPorVal = sql.NullInt64{Valid: false}
	}

	return &Recurso{
		Titulo:    titulo,
		Categoria: categoria,
		ImagenURL: sql.NullString{
			String: imagenURL,
			Valid:  imagenURL != "",
		},
		AudioURL: sql.NullString{
			String: audioURL,
			Valid:  audioURL != "",
		},
		Autor: sql.NullString{
			String: autor,
			Valid:  autor != "",
		},
		CreadoPor:   creadoPorVal,
		Estado:      EstadoDisponible,
		Descripcion: descripcion,
	}
}

// IsValid valida si el recurso tiene datos válidos
func (r *Recurso) IsValid() bool {
	return r.Titulo != ""
}

// GetEstadoString obtiene el estado como string
func (r *Recurso) GetEstadoString() string {
	return string(r.Estado)
}

// SetEstado cambia el estado del recurso
func (r *Recurso) SetEstado(estado EstadoRecurso) {
	r.Estado = estado
}

// IsDisponible verifica si el recurso está disponible
func (r *Recurso) IsDisponible() bool {
	return r.Estado == EstadoDisponible
}

// GetImagenURL obtiene la URL de la imagen de forma segura
func (r *Recurso) GetImagenURL() string {
	if r.ImagenURL.Valid {
		return r.ImagenURL.String
	}
	return ""
}

// GetAudioURL obtiene la URL del audio de forma segura
func (r *Recurso) GetAudioURL() string {
	if r.AudioURL.Valid {
		return r.AudioURL.String
	}
	return ""
}

// GetAutor obtiene el autor de forma segura
func (r *Recurso) GetAutor() string {
	if r.Autor.Valid {
		return r.Autor.String
	}
	return ""
}

// GetCreadoPor obtiene el CreadoPor de forma segura
func (r *Recurso) GetCreadoPor() int {
	if r.CreadoPor.Valid {
		return int(r.CreadoPor.Int64)
	}
	return 0
}
