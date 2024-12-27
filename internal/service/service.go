package service

import "github.com/marlonlacerda2/simple-CRUD/internal/model"

func GetHelloMessage() model.HelloResponse {
	return model.HelloResponse{Message: "Hello from the Service!"}
}
