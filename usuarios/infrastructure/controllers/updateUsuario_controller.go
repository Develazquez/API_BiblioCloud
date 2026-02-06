package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"biblioteca-api/usuarios/application"
	"biblioteca-api/usuarios/domain/entities"
)

type UpdateUsuarioController struct {
	usecase *application.UpdateUsuarioUseCase
}

func NewUpdateUsuarioController(usecase *application.UpdateUsuarioUseCase) *UpdateUsuarioController {
	return &UpdateUsuarioController{usecase: usecase}
}

func (c *UpdateUsuarioController) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var usuario entities.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usuario.ID = id

	resultado, err := c.usecase.Execute(&usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultado)
}
