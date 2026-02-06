package routes

import (
	"database/sql"

	"github.com/gorilla/mux"

	"biblioteca-api/usuarios/application"
	"biblioteca-api/usuarios/infrastructure/controllers"
	"biblioteca-api/usuarios/infrastructure/repository"
)

func UsuarioRoutes(router *mux.Router, db *sql.DB) {
	usuarioRepo := repository.NewUsuarioRepositoryPostgres(db)

	createUseCase := application.NewCreateUsuarioUseCase(usuarioRepo)
	getByIDUseCase := application.NewGetUsuarioPorIDUseCase(usuarioRepo)
	getTodosUseCase := application.NewGetTodosUsuariosUseCase(usuarioRepo)
	updateUseCase := application.NewUpdateUsuarioUseCase(usuarioRepo)
	deleteUseCase := application.NewDeleteUsuarioUseCase(usuarioRepo)

	createController := controllers.NewCreateUsuarioController(createUseCase)
	getByIDController := controllers.NewGetUsuarioPorIDController(getByIDUseCase)
	getTodosController := controllers.NewGetTodosUsuariosController(getTodosUseCase)
	updateController := controllers.NewUpdateUsuarioController(updateUseCase)
	deleteController := controllers.NewDeleteUsuarioController(deleteUseCase)

	router.HandleFunc("/usuarios", createController.Handle).Methods("POST")
	router.HandleFunc("/usuarios", getTodosController.Handle).Methods("GET")
	router.HandleFunc("/usuarios/{id}", getByIDController.Handle).Methods("GET")
	router.HandleFunc("/usuarios/{id}", updateController.Handle).Methods("PUT")
	router.HandleFunc("/usuarios/{id}", deleteController.Handle).Methods("DELETE")
}
