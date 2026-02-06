package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"biblioteca-api/prestamos/application"
)

type GetPrestamoPorIDController struct {
	usecase *application.GetPrestamoPorIDUseCase
}

func NewGetPrestamoPorIDController(usecase *application.GetPrestamoPorIDUseCase) *GetPrestamoPorIDController {
	return &GetPrestamoPorIDController{usecase: usecase}
}

func (c *GetPrestamoPorIDController) Handle(w http.ResponseWriter, r *http.Request) {
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
