package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/helpers/responses"
	"api/src/models"
	"api/src/repositories"
	"api/src/security"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Login authenticate user
func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
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
	userDatabase, err := repository.GetByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if userDatabase.ID == 0 {
		message := "Ususário não existe"
		responses.Error(w, http.StatusUnauthorized, errors.New(message))
		return
	}

	if err = security.VerifyPassword(user.Password, userDatabase.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.GenerateToken(userDatabase.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.ToJson(w, http.StatusOK, map[string]string{"token": token})
}
