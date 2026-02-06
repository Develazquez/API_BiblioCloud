package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"biblioteca-api/recursos/application"
)

// GetRecursoPorIDController obtiene un recurso por ID
type GetRecursoPorIDController struct {
	usecase *application.GetRecursoPorIDUseCase
}

// NewGetRecursoPorIDController crea una nueva instancia
func NewGetRecursoPorIDController(usecase *application.GetRecursoPorIDUseCase) *GetRecursoPorIDController {
	return &GetRecursoPorIDController{usecase: usecase}
}

// Handle maneja la petición HTTP
func (c *GetRecursoPorIDController) Handle(w http.ResponseWriter, r *http.Request) {
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
