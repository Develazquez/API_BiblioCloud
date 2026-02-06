package controllers

import (
	"encoding/json"
	"net/http"

	"biblioteca-api/usuarios/application"
	"biblioteca-api/usuarios/domain/entities"
)

// CreateUsuarioController maneja la creación de usuarios
type CreateUsuarioController struct {
	usecase *application.CreateUsuarioUseCase
}

// NewCreateUsuarioController crea una nueva instancia
func NewCreateUsuarioController(usecase *application.CreateUsuarioUseCase) *CreateUsuarioController {
	return &CreateUsuarioController{usecase: usecase}
}

// Handle maneja la petición HTTP
func (c *CreateUsuarioController) Handle(w http.ResponseWriter, r *http.Request) {
	var usuario entities.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Establecer estado por defecto si no se proporciona
	if usuario.Estado == "" {
		usuario.Estado = entities.EstadoActivo
	}

	resultado, err := c.usecase.Execute(&usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}
