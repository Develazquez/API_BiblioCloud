package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"biblioteca-api/recursos/application"
)

// DeleteRecursoController elimina un recurso
type DeleteRecursoController struct {
	usecase *application.DeleteRecursoUseCase
}

// NewDeleteRecursoController crea una nueva instancia
func NewDeleteRecursoController(usecase *application.DeleteRecursoUseCase) *DeleteRecursoController {
	return &DeleteRecursoController{usecase: usecase}
}

// Handle maneja la petición HTTP
func (c *DeleteRecursoController) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := c.usecase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
