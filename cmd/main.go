package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/marlonlacerda2/simple-CRUD/internal/handler"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	r := mux.NewRouter()

	r.HandleFunc("/api/users", handler.GetUsersHandler).Methods("GET")
	log.Println("Iniciando o servidor na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
