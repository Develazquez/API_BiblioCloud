package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"biblioteca-api/usuarios/application"
)

// DeleteUsuarioController elimina un usuario
type DeleteUsuarioController struct {
	usecase *application.DeleteUsuarioUseCase
}

// NewDeleteUsuarioController crea una nueva instancia
func NewDeleteUsuarioController(usecase *application.DeleteUsuarioUseCase) *DeleteUsuarioController {
	return &DeleteUsuarioController{usecase: usecase}
}

// Handle maneja la petición HTTP
func (c *DeleteUsuarioController) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := c.usecase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
