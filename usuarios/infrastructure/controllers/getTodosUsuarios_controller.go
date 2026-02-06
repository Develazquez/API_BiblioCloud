package controllers

import (
	"encoding/json"
	"net/http"

	"biblioteca-api/usuarios/application"
)

// GetTodosUsuariosController obtiene todos los usuarios
type GetTodosUsuariosController struct {
	usecase *application.GetTodosUsuariosUseCase
}

// NewGetTodosUsuariosController crea una nueva instancia
func NewGetTodosUsuariosController(usecase *application.GetTodosUsuariosUseCase) *GetTodosUsuariosController {
	return &GetTodosUsuariosController{usecase: usecase}
}

// Handle maneja la petición HTTP
func (c *GetTodosUsuariosController) Handle(w http.ResponseWriter, r *http.Request) {
	resultado, err := c.usecase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultado)
}
