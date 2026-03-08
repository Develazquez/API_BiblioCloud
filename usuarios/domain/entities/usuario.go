package entities

type EstadoUsuario string

const (
	EstadoActivo EstadoUsuario = "ACTIVO"
	EstadoDeudor EstadoUsuario = "DEUDOR"
)

type Usuario struct {
	ID                        int				`json:"id"`
	Nombre                    string			`json:"nombre"`
	Email                     string			`json:"email"`
	Password                  string			`json:"password"`
	Estado                    EstadoUsuario		`json:"estado"`
	CantidadPrestamosActuales int				`json:"cantidadPrestamosActuales"`
}

func NewUsuario(nombre, email, password string) *Usuario {
	return &Usuario{
		Nombre:                    nombre,
		Email:                     email,
		Password:                  password,
		Estado:                    EstadoActivo,
		CantidadPrestamosActuales: 0,
	}
}

func (u *Usuario) IsValid() bool {
	return u.Nombre != "" && u.Email != "" && u.Password != ""
}

func (u *Usuario) GetEstadoString() string {
	return string(u.Estado)
}

func (u *Usuario) SetEstado(estado EstadoUsuario) {
	u.Estado = estado
}
