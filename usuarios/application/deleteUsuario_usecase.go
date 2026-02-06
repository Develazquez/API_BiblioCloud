package application

import (
	"errors"
	"biblioteca-api/usuarios/domain/repository"
)

type DeleteUsuarioUseCase struct {
	repository repository.UsuarioRepository
}

func NewDeleteUsuarioUseCase(repo repository.UsuarioRepository) *DeleteUsuarioUseCase {
	return &DeleteUsuarioUseCase{repository: repo}
}

func (uc *DeleteUsuarioUseCase) Execute(id int) error {
	if id <= 0 {
		return errors.New("ID inválido")
	}

	existente, _ := uc.repository.ObtenerPorID(id)
	if existente == nil {
		return errors.New("usuario no encontrado")
	}

	return uc.repository.Eliminar(id)
}
