package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"biblioteca-api/usuarios/application"
)

// GetUsuarioPorIDController obtiene un usuario por ID
type GetUsuarioPorIDController struct {
	usecase *application.GetUsuarioPorIDUseCase
}

// NewGetUsuarioPorIDController crea una nueva instancia
func NewGetUsuarioPorIDController(usecase *application.GetUsuarioPorIDUseCase) *GetUsuarioPorIDController {
	return &GetUsuarioPorIDController{usecase: usecase}
}

// Handle maneja la petición HTTP
func (c *GetUsuarioPorIDController) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	resultado, err := c.usecase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultado)
}
