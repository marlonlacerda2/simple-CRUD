package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/marlonlacerda2/simple-CRUD/internal/service"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetUsers()
	if err != nil {
		http.Error(w, "Error getting users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Println("Erro ao converter os usuários para JSON", err)
		http.Error(w, "Error converting users to JSON", http.StatusInternalServerError)
	}
}
