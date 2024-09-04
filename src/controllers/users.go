package controllers

import "net/http"

// CreateUser create an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!"))
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
