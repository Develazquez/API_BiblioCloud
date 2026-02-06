package routes

import (
	"database/sql"

	"github.com/gorilla/mux"

	"biblioteca-api/prestamos/application"
	"biblioteca-api/prestamos/infrastructure/controllers"
	"biblioteca-api/prestamos/infrastructure/repository"
)

func PrestamoRoutes(router *mux.Router, db *sql.DB) {
	prestamoRepo := repository.NewPrestamoRepositoryPostgres(db)

	createUseCase := application.NewCreatePrestamoUseCase(prestamoRepo)
	getByIDUseCase := application.NewGetPrestamoPorIDUseCase(prestamoRepo)
	getTodosUseCase := application.NewGetTodosPrestamosUseCase(prestamoRepo)
	updateUseCase := application.NewUpdatePrestamoUseCase(prestamoRepo)
	deleteUseCase := application.NewDeletePrestamoUseCase(prestamoRepo)
	devolverUseCase := application.NewDevolverPrestamoUseCase(prestamoRepo)

	createController := controllers.NewCreatePrestamoController(createUseCase)
	getByIDController := controllers.NewGetPrestamoPorIDController(getByIDUseCase)
	getTodosController := controllers.NewGetTodosPrestamosController(getTodosUseCase)
	updateController := controllers.NewUpdatePrestamoController(updateUseCase)
	deleteController := controllers.NewDeletePrestamoController(deleteUseCase)
	devolverController := controllers.NewDevolverPrestamoController(devolverUseCase)

	router.HandleFunc("/prestamos", createController.Handle).Methods("POST")
	router.HandleFunc("/prestamos", getTodosController.Handle).Methods("GET")
	router.HandleFunc("/prestamos/{id}", getByIDController.Handle).Methods("GET")
	router.HandleFunc("/prestamos/{id}", updateController.Handle).Methods("PUT")
	router.HandleFunc("/prestamos/{id}", deleteController.Handle).Methods("DELETE")
	router.HandleFunc("/prestamos/{id}/devolver", devolverController.Handle).Methods("POST")
}
