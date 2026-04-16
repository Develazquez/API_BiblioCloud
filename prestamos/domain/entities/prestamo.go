package entities

import "time"

type EstadoPrestamo string

const (
	EstadoActivo   EstadoPrestamo = "ACTIVO"
	EstadoDevuelto EstadoPrestamo = "DEVUELTO"
	EstadoVencido  EstadoPrestamo = "VENCIDO"
)

type Prestamo struct {
	ID              int            `json:"id"`
	UsuarioID       int            `json:"usuario_id"`
	RecursoID       int            `json:"recurso_id"`
	FechaInicio     time.Time      `json:"fecha_prestamo"`
	FechaLimite     time.Time      `json:"fecha_devolucion_estimada"`
	FechaDevolucion *time.Time     `json:"fecha_devolucion,omitempty"`
	Estado          EstadoPrestamo `json:"estado"`
}

func NewPrestamo(usuarioID, recursoID int, fechaLimite time.Time) *Prestamo {
	return &Prestamo{
		UsuarioID:   usuarioID,
		RecursoID:   recursoID,
		FechaInicio: time.Now(),
		FechaLimite: fechaLimite,
		Estado:      EstadoActivo,
	}
}

func (p *Prestamo) IsValid() bool {
	return p.UsuarioID > 0 && p.RecursoID > 0 && p.FechaLimite.After(time.Now())
}

func (p *Prestamo) GetEstadoString() string {
	return string(p.Estado)
}

func (p *Prestamo) SetEstado(estado EstadoPrestamo) {
	p.Estado = estado
}

func (p *Prestamo) Devolver() {
	ahora := time.Now()
	p.FechaDevolucion = &ahora
	p.Estado = EstadoDevuelto
}

func (p *Prestamo) IsActivo() bool {
	return p.Estado == EstadoActivo
}

func (p *Prestamo) IsVencido() bool {
	return (p.IsActivo() || p.Estado == EstadoVencido) && time.Now().After(p.FechaLimite)
}
