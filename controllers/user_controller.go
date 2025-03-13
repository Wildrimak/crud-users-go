package controllers

import (
	"cadastro-usuarios-go/models"
	"cadastro-usuarios-go/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Entrando no post users", r)
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Erro ao decodificar JSON:", err)
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	id, err := c.Service.CreateUser(user)
	if err != nil {
		log.Println("Erro ao criar usuário:", err)
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Entrando no get users", r)
	users, err := c.Service.GetUsers()
	if err != nil {
		log.Println("Erro ao buscar usuários:", err)
		http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	log.Println("Entrando no update user: ", r)

	id, err := c.extractIDFromURL(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	log.Println("ID extraído:", id)

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Erro ao decodificar JSON:", err)
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	user.ID = id
	if err := c.Service.UpdateUser(user); err != nil {
		log.Println("Erro ao atualizar usuário:", err)
		http.Error(w, "Erro ao atualizar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário atualizado com sucesso"})
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Entrando no delete user", r)

	id, err := c.extractIDFromURL(r)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	log.Println("ID extraído:", id)

	if err := c.Service.DeleteUser(id); err != nil {
		log.Println("Erro ao excluir usuário:", err)
		http.Error(w, "Erro ao excluir usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário excluído com sucesso"})

}

func (c *UserController) extractIDFromURL(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {
		log.Println("ID não encontrado na URL")
		return 0, http.ErrMissingFile
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Println("ID inválido:", err)
		return 0, err
	}

	return id, nil

}
