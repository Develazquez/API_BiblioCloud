package application

import (
	"biblioteca-api/usuarios/domain/entities"
	"biblioteca-api/usuarios/domain/repository"
	"errors"
)

type LoginUsuarioUseCase struct {
	repository repository.UsuarioRepository
}

func NewLoginUsuarioUseCase(repo repository.UsuarioRepository) *LoginUsuarioUseCase {
	return &LoginUsuarioUseCase{repository: repo}
}

func (uc *LoginUsuarioUseCase) Execute(email, password string) (*entities.Usuario, error) {
	if email == "" || password == "" {
		return nil, errors.New("email y contraseña son requeridos")
	}

	usuario, err := uc.repository.ObtenerPorEmail(email)
	if err != nil || usuario == nil {
		return nil, errors.New("usuario no encontrado")
	}

	if usuario.Password != password {
		return nil, errors.New("contraseña incorrecta")
	}

	return usuario, nil
}
