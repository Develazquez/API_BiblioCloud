package controllers

import (
	"encoding/json"
	"net/http"

	"biblioteca-api/recursos/application"
	"biblioteca-api/recursos/domain/entities"
)

// CreateRecursoController maneja la creación de recursos
type CreateRecursoController struct {
	usecase *application.CreateRecursoUseCase
}

// NewCreateRecursoController crea una nueva instancia
func NewCreateRecursoController(usecase *application.CreateRecursoUseCase) *CreateRecursoController {
	return &CreateRecursoController{usecase: usecase}
}

// Handle maneja la petición HTTP
func (c *CreateRecursoController) Handle(w http.ResponseWriter, r *http.Request) {
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}
