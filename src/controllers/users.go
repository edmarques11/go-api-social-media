package controllers

import (
	"api/src/database"
	"api/src/helpers/responses"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io"
	"net/http"
)

// CreateUser create an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.MakeConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.ToJson(w, http.StatusCreated, user)
}

// GetUsers get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários!"))
}

// GetUserById get an user by id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário por id!"))
}

// UpdateUser update an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualiza usuário!"))
}

// DeleteUser delete an usser
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário!"))
}
