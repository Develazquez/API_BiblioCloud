package controllers

import (
	"encoding/json"
	"net/http"

	"biblioteca-api/prestamos/application"
	"biblioteca-api/prestamos/domain/entities"
)

type UpdatePrestamoController struct {
	usecase *application.UpdatePrestamoUseCase
}

func NewUpdatePrestamoController(usecase *application.UpdatePrestamoUseCase) *UpdatePrestamoController {
	return &UpdatePrestamoController{usecase: usecase}
}

func (c *UpdatePrestamoController) Handle(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(resultado)
}
