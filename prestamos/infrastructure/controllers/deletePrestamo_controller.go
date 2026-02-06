package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"biblioteca-api/prestamos/application"
)

type DeletePrestamoController struct {
	usecase *application.DeletePrestamoUseCase
}

func NewDeletePrestamoController(usecase *application.DeletePrestamoUseCase) *DeletePrestamoController {
	return &DeletePrestamoController{usecase: usecase}
}

func (c *DeletePrestamoController) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := c.usecase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
