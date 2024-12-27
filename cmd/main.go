package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marlonlacerda2/simple-CRUD/internal/handler"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/", handler.HelloHandler).Methods("GET")
	log.Println("Iniciando o servidor na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
