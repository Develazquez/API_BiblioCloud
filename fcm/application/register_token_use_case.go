package application

import "biblioteca-api/fcm/domain/repository"

type RegisterTokenUseCase struct {
	repo repository.FCMTokenRepository
}

func NewRegisterTokenUseCase(repo repository.FCMTokenRepository) *RegisterTokenUseCase {
	return &RegisterTokenUseCase{repo: repo}
}

func (uc *RegisterTokenUseCase) Execute(usuarioID int, token string) error {
	return uc.repo.Save(usuarioID, token)
}
