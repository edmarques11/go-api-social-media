package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/helpers/responses"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err = user.Prepare(models.SaveUser); err != nil {
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

	user.Password = ""
	responses.ToJson(w, http.StatusCreated, user)
}

// GetUsers get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, err := database.MakeConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	users, err := repository.GetUsers(nameOrNick)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.ToJson(w, http.StatusOK, users)
}

// GetUserById get an user by id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseInt(params["userId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.MakeConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user, err := repository.GetById(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.ToJson(w, http.StatusOK, user)
}

// UpdateUser update an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseInt(params["userId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userTokenId, err := auth.GetUserIdToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userId != int64(userTokenId) {
		responses.Error(w, http.StatusForbidden, errors.New("usuário sem permisão"))
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.ToJson(w, http.StatusBadRequest, err)
	}

	if err = user.Prepare(models.UpdateUser); err != nil {
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
	if err = repository.UpdateUser(userId, user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.ToJson(w, http.StatusNoContent, nil)
}

// DeleteUser delete an usser
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseInt(params["userId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userTokenId, err := auth.GetUserIdToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userId != int64(userTokenId) {
		responses.Error(w, http.StatusForbidden, errors.New("usuário sem permisão"))
		return
	}

	db, err := database.MakeConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	err = repository.DeleteUser(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.ToJson(w, http.StatusNoContent, nil)
}

// UserFollow follow an user
func UserFollow(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.GetUserIdToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if userId == followerId {
		message := "Não é possível seguir você mesmo, seu arigó!"
		responses.Error(w, http.StatusForbidden, errors.New(message))
		return
	}

	db, err := database.MakeConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err := repository.Follow(userId, followerId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.ToJson(w, http.StatusNoContent, nil)
}
