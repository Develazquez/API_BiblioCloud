package application

import (
	"biblioteca-api/fcm/domain/repository"
	"biblioteca-api/fcm/infrastructure/services"
)

type SendNotificationUseCase struct {
	repo repository.FCMTokenRepository
	fcm  services.FirebaseMessagingService
}

func NewSendNotificationUseCase(repo repository.FCMTokenRepository, fcm services.FirebaseMessagingService) *SendNotificationUseCase {
	return &SendNotificationUseCase{repo: repo, fcm: fcm}
}

func (uc *SendNotificationUseCase) Execute(usuarioID int, data map[string]string) {
	tokens, err := uc.repo.GetTokensByUsuarioID(usuarioID)
	if err != nil || len(tokens) == 0 {
		return
	}

	for _, token := range tokens {
		_ = uc.fcm.SendMessage(token, data)
	}
}
