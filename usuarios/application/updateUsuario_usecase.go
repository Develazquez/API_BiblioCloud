package application

import (
	"errors"
	"biblioteca-api/usuarios/domain/entities"
	"biblioteca-api/usuarios/domain/repository"
)

type UpdateUsuarioUseCase struct {
	repository repository.UsuarioRepository
}

func NewUpdateUsuarioUseCase(repo repository.UsuarioRepository) *UpdateUsuarioUseCase {
	return &UpdateUsuarioUseCase{repository: repo}
}

func (uc *UpdateUsuarioUseCase) Execute(usuario *entities.Usuario) (*entities.Usuario, error) {
	if usuario.ID <= 0 {
		return nil, errors.New("ID de usuario inválido")
	}

	if !usuario.IsValid() {
		return nil, errors.New("usuario no válido: faltan datos requeridos")
	}

	existente, _ := uc.repository.ObtenerPorID(usuario.ID)
	if existente == nil {
		return nil, errors.New("usuario no encontrado")
	}

	return uc.repository.Actualizar(usuario)
}
