package controllers

import (
	"encoding/json"
	"net/http"

	"biblioteca-api/prestamos/application"
)

type GetTodosPrestamosController struct {
	usecase *application.GetTodosPrestamosUseCase
}

func NewGetTodosPrestamosController(usecase *application.GetTodosPrestamosUseCase) *GetTodosPrestamosController {
	return &GetTodosPrestamosController{usecase: usecase}
}

func (c *GetTodosPrestamosController) Handle(w http.ResponseWriter, r *http.Request) {
	resultado, err := c.usecase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultado)
}
