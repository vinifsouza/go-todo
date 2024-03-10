package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vinifsouza/go-todo/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Houve um erro ao buscar todos os registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
