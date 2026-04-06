package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	fcmApp "biblioteca-api/fcm/application"
	fcmRepo "biblioteca-api/fcm/infrastructure/repository"
	fcmServices "biblioteca-api/fcm/infrastructure/services"
	"biblioteca-api/prestamos/application"
	"biblioteca-api/prestamos/infrastructure/controllers"
	"biblioteca-api/prestamos/infrastructure/repository"
)

func PrestamoRoutes(router *gin.Engine, db *sql.DB) {
	prestamoRepo := repository.NewPrestamoRepositoryPostgres(db)

	// Instanciar dependencias FCM
	fcmTokenRepo := fcmRepo.NewFCMTokenRepositoryPostgres(db)
	fcmService := fcmServices.NewFirebaseMessagingService()
	sendNotification := fcmApp.NewSendNotificationUseCase(fcmTokenRepo, fcmService)

	createUseCase := application.NewCreatePrestamoUseCase(prestamoRepo, sendNotification)
	getByIDUseCase := application.NewGetPrestamoPorIDUseCase(prestamoRepo)
	getTodosUseCase := application.NewGetTodosPrestamosUseCase(prestamoRepo)
	updateUseCase := application.NewUpdatePrestamoUseCase(prestamoRepo)
	deleteUseCase := application.NewDeletePrestamoUseCase(prestamoRepo)
	devolverUseCase := application.NewDevolverPrestamoUseCase(prestamoRepo, sendNotification)

	createController := controllers.NewCreatePrestamoController(createUseCase)
	getByIDController := controllers.NewGetPrestamoPorIDController(getByIDUseCase)
	getTodosController := controllers.NewGetTodosPrestamosController(getTodosUseCase)
	updateController := controllers.NewUpdatePrestamoController(updateUseCase)
	deleteController := controllers.NewDeletePrestamoController(deleteUseCase)
	devolverController := controllers.NewDevolverPrestamoController(devolverUseCase)

	router.POST("/prestamos", createController.Handle)
	router.GET("/prestamos", getTodosController.Handle)
	router.GET("/prestamos/:id", getByIDController.Handle)
	router.PUT("/prestamos/:id", updateController.Handle)
	router.DELETE("/prestamos/:id", deleteController.Handle)
	router.POST("/prestamos/:id/devolver", devolverController.Handle)
}
