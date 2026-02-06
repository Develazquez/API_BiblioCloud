package application

import (
	"errors"
	"biblioteca-api/usuarios/domain/repository"
)

type GetUsuarioPorIDUseCase struct {
	repository repository.UsuarioRepository
}

func NewGetUsuarioPorIDUseCase(repo repository.UsuarioRepository) *GetUsuarioPorIDUseCase {
	return &GetUsuarioPorIDUseCase{repository: repo}
}

func (uc *GetUsuarioPorIDUseCase) Execute(id int) (interface{}, error) {
	if id <= 0 {
		return nil, errors.New("ID inválido")
	}

	return uc.repository.ObtenerPorID(id)
}
