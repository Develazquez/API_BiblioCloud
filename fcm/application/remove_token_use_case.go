package application

import "biblioteca-api/fcm/domain/repository"

type RemoveTokenUseCase struct {
	repo repository.FCMTokenRepository
}

func NewRemoveTokenUseCase(repo repository.FCMTokenRepository) *RemoveTokenUseCase {
	return &RemoveTokenUseCase{repo: repo}
}

func (uc *RemoveTokenUseCase) Execute(usuarioID int, token string) error {
	return uc.repo.Delete(usuarioID, token)
}
