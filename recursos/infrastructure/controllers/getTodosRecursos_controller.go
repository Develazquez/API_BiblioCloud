package controllers

import (
	"encoding/json"
	"net/http"

	"biblioteca-api/recursos/application"
)

// GetTodosRecursosController obtiene todos los recursos
type GetTodosRecursosController struct {
	usecase *application.GetTodosRecursosUseCase
}

// NewGetTodosRecursosController crea una nueva instancia
func NewGetTodosRecursosController(usecase *application.GetTodosRecursosUseCase) *GetTodosRecursosController {
	return &GetTodosRecursosController{usecase: usecase}
}

// Handle maneja la petición HTTP
func (c *GetTodosRecursosController) Handle(w http.ResponseWriter, r *http.Request) {
	resultado, err := c.usecase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultado)
}
