package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"biblioteca-api/usuarios/application"
	"biblioteca-api/usuarios/infrastructure/controllers"
	"biblioteca-api/usuarios/infrastructure/repository"
)

func UsuarioRoutes(router *gin.Engine, db *sql.DB) {
	usuarioRepo := repository.NewUsuarioRepositoryPostgres(db)

	createUseCase := application.NewCreateUsuarioUseCase(usuarioRepo)
	getByIDUseCase := application.NewGetUsuarioPorIDUseCase(usuarioRepo)
	getTodosUseCase := application.NewGetTodosUsuariosUseCase(usuarioRepo)
	updateUseCase := application.NewUpdateUsuarioUseCase(usuarioRepo)
	deleteUseCase := application.NewDeleteUsuarioUseCase(usuarioRepo)
	loginUseCase := application.NewLoginUsuarioUseCase(usuarioRepo)

	createController := controllers.NewCreateUsuarioController(createUseCase)
	getByIDController := controllers.NewGetUsuarioPorIDController(getByIDUseCase)
	getTodosController := controllers.NewGetTodosUsuariosController(getTodosUseCase)
	updateController := controllers.NewUpdateUsuarioController(updateUseCase)
	deleteController := controllers.NewDeleteUsuarioController(deleteUseCase)
	loginController := controllers.NewLoginUsuarioController(loginUseCase)

	router.POST("/usuarios", createController.Handle)
	router.POST("/usuarios/login", loginController.Handle)
	router.GET("/usuarios", getTodosController.Handle)
	router.GET("/usuarios/:id", getByIDController.Handle)
	router.PUT("/usuarios/:id", updateController.Handle)
	router.DELETE("/usuarios/:id", deleteController.Handle)
}
