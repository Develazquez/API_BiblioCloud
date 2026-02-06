package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"biblioteca-api/prestamos/application"
)

type DevolverPrestamoController struct {
	usecase *application.DevolverPrestamoUseCase
}

func NewDevolverPrestamoController(usecase *application.DevolverPrestamoUseCase) *DevolverPrestamoController {
	return &DevolverPrestamoController{usecase: usecase}
}

func (c *DevolverPrestamoController) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := c.usecase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
