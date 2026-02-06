package entities

// EstadoRecurso representa los estados posibles de un recurso
type EstadoRecurso string

const (
	EstadoDisponible EstadoRecurso = "DISPONIBLE"
	EstadoPrestado   EstadoRecurso = "PRESTADO"
)

// Recurso representa la entidad de recurso en el dominio
type Recurso struct {
	ID          int
	Titulo      string
	Categoria   string
	ImagenURL   string
	Estado      EstadoRecurso
	Descripcion string
}

// NewRecurso crea una nueva instancia de Recurso
func NewRecurso(titulo, categoria, imagenURL, descripcion string) *Recurso {
	return &Recurso{
		Titulo:      titulo,
		Categoria:   categoria,
		ImagenURL:   imagenURL,
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
