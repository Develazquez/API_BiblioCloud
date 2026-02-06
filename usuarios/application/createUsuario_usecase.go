package application

import (
	"errors"
	"biblioteca-api/usuarios/domain/entities"
	"biblioteca-api/usuarios/domain/repository"
)

type CreateUsuarioUseCase struct {
	repository repository.UsuarioRepository
}

func NewCreateUsuarioUseCase(repo repository.UsuarioRepository) *CreateUsuarioUseCase {
	return &CreateUsuarioUseCase{repository: repo}
}

func (uc *CreateUsuarioUseCase) Execute(usuario *entities.Usuario) (*entities.Usuario, error) {
	if !usuario.IsValid() {
		return nil, errors.New("usuario no válido: faltan datos requeridos")
	}

	existente, _ := uc.repository.ObtenerPorEmail(usuario.Email)
	if existente != nil {
		return nil, errors.New("el email ya está registrado")
	}

	return uc.repository.Crear(usuario)
}
