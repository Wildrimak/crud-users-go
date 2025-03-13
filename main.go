package main

import (
	"cadastro-usuarios-go/config"
	"cadastro-usuarios-go/controllers"
	"cadastro-usuarios-go/repositories"
	"cadastro-usuarios-go/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	repo := repositories.NewUserRepository(config.DB)
	service := services.NewUserService(repo)
	controller := controllers.NewUserController(service)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.CreateUser(w, r)
		case http.MethodGet:
			controller.GetUsers(w, r)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			controller.UpdateUser(w, r)
		case http.MethodDelete:
			controller.DeleteUser(w, r)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
