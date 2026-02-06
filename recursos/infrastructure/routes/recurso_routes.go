package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"biblioteca-api/recursos/application"
	"biblioteca-api/recursos/infrastructure/controllers"
	"biblioteca-api/recursos/infrastructure/repository"
)

func RecursoRoutes(router *gin.Engine, db *sql.DB) {
	recursoRepo := repository.NewRecursoRepositoryPostgres(db)

	createUseCase := application.NewCreateRecursoUseCase(recursoRepo)
	getByIDUseCase := application.NewGetRecursoPorIDUseCase(recursoRepo)
	getTodosUseCase := application.NewGetTodosRecursosUseCase(recursoRepo)
	updateUseCase := application.NewUpdateRecursoUseCase(recursoRepo)
	deleteUseCase := application.NewDeleteRecursoUseCase(recursoRepo)

	createController := controllers.NewCreateRecursoController(createUseCase)
	getByIDController := controllers.NewGetRecursoPorIDController(getByIDUseCase)
	getTodosController := controllers.NewGetTodosRecursosController(getTodosUseCase)
	updateController := controllers.NewUpdateRecursoController(updateUseCase)
	deleteController := controllers.NewDeleteRecursoController(deleteUseCase)

	router.POST("/recursos", createController.Handle)
	router.GET("/recursos", getTodosController.Handle)
	router.GET("/recursos/:id", getByIDController.Handle)
	router.PUT("/recursos/:id", updateController.Handle)
	router.DELETE("/recursos/:id", deleteController.Handle)
}
