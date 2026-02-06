package routes

import (
	"database/sql"

	"github.com/gorilla/mux"

	"biblioteca-api/recursos/application"
	"biblioteca-api/recursos/infrastructure/controllers"
	"biblioteca-api/recursos/infrastructure/repository"
)

// RecursoRoutes configura las rutas para recursos
func RecursoRoutes(router *mux.Router, db *sql.DB) {
	// Crear instancias del repositorio
	recursoRepo := repository.NewRecursoRepositoryPostgres(db)

	// Crear instancias de los use cases
	createUseCase := application.NewCreateRecursoUseCase(recursoRepo)
	getByIDUseCase := application.NewGetRecursoPorIDUseCase(recursoRepo)
	getTodosUseCase := application.NewGetTodosRecursosUseCase(recursoRepo)
	updateUseCase := application.NewUpdateRecursoUseCase(recursoRepo)
	deleteUseCase := application.NewDeleteRecursoUseCase(recursoRepo)

	// Crear instancias de los controladores
	createController := controllers.NewCreateRecursoController(createUseCase)
	getByIDController := controllers.NewGetRecursoPorIDController(getByIDUseCase)
	getTodosController := controllers.NewGetTodosRecursosController(getTodosUseCase)
	updateController := controllers.NewUpdateRecursoController(updateUseCase)
	deleteController := controllers.NewDeleteRecursoController(deleteUseCase)

	// Definir las rutas
	router.HandleFunc("/recursos", createController.Handle).Methods("POST")
	router.HandleFunc("/recursos", getTodosController.Handle).Methods("GET")
	router.HandleFunc("/recursos/{id}", getByIDController.Handle).Methods("GET")
	router.HandleFunc("/recursos/{id}", updateController.Handle).Methods("PUT")
	router.HandleFunc("/recursos/{id}", deleteController.Handle).Methods("DELETE")
}
