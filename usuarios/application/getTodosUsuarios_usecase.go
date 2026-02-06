package application

import (
	"biblioteca-api/usuarios/domain/repository"
)

type GetTodosUsuariosUseCase struct {
	repository repository.UsuarioRepository
}

func NewGetTodosUsuariosUseCase(repo repository.UsuarioRepository) *GetTodosUsuariosUseCase {
	return &GetTodosUsuariosUseCase{repository: repo}
}

func (uc *GetTodosUsuariosUseCase) Execute() (interface{}, error) {
	return uc.repository.ObtenerTodos()
}
