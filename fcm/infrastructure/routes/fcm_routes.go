package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"biblioteca-api/fcm/application"
	"biblioteca-api/fcm/infrastructure/controllers"
	"biblioteca-api/fcm/infrastructure/repository"
)

func FCMRoutes(router *gin.Engine, db *sql.DB) {
	fcmRepo := repository.NewFCMTokenRepositoryPostgres(db)

	registerTokenUseCase := application.NewRegisterTokenUseCase(fcmRepo)
	removeTokenUseCase := application.NewRemoveTokenUseCase(fcmRepo)

	fcmController := controllers.NewFCMController(registerTokenUseCase, removeTokenUseCase)

	fcmGroup := router.Group("/fcm")
	{
		fcmGroup.POST("/register-token", fcmController.RegisterToken)
		fcmGroup.DELETE("/remove-token", fcmController.RemoveToken)
	}
}
