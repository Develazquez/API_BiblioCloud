package repository

type FCMTokenRepository interface {
	Save(usuarioID int, token string) error
	Delete(usuarioID int, token string) error
	DeleteByUsuarioID(usuarioID int) error
	GetTokensByUsuarioID(usuarioID int) ([]string, error)
}
