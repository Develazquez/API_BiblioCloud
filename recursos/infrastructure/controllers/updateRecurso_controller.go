package controllers

import (
	"encoding/json"
	"net/http"

	"biblioteca-api/recursos/application"
	"biblioteca-api/recursos/domain/entities"
)

// UpdateRecursoController actualiza un recurso
type UpdateRecursoController struct {
	usecase *application.UpdateRecursoUseCase
}

// NewUpdateRecursoController crea una nueva instancia
func NewUpdateRecursoController(usecase *application.UpdateRecursoUseCase) *UpdateRecursoController {
	return &UpdateRecursoController{usecase: usecase}
}

// Handle maneja la petición HTTP
func (c *UpdateRecursoController) Handle(w http.ResponseWriter, r *http.Request) {
	var recurso entities.Recurso
	err := json.NewDecoder(r.Body).Decode(&recurso)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultado, err := c.usecase.Execute(&recurso)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultado)
}
