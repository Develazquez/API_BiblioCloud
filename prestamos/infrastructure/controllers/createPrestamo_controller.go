package controllers

import (
	"encoding/json"
	"net/http"

	"biblioteca-api/prestamos/application"
	"biblioteca-api/prestamos/domain/entities"
)

type CreatePrestamoController struct {
	usecase *application.CreatePrestamoUseCase
}

func NewCreatePrestamoController(usecase *application.CreatePrestamoUseCase) *CreatePrestamoController {
	return &CreatePrestamoController{usecase: usecase}
}

func (c *CreatePrestamoController) Handle(w http.ResponseWriter, r *http.Request) {
	var prestamo entities.Prestamo
	err := json.NewDecoder(r.Body).Decode(&prestamo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultado, err := c.usecase.Execute(&prestamo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}
