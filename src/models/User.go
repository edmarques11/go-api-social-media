package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

const (
	SaveUser   = "SaveUser"
	UpdateUser = "UpdateUser"
)

// User represents an user
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare validate data and format
func (user *User) Prepare(step string) error {
	if err := user.validator(step); err != nil {
		return err
	}
	if err := user.formatData(step); err != nil {
		return err
	}
	return nil
}

func (user *User) validator(step string) error {
	var msg string

	switch {
	case user.Name == "":
		msg = "O campo 'name' é obrigatório e não pode estar em branco"
	case user.Email == "":
		msg = "O campo 'email' é obrigatório e não pode estar em branco"
	case user.Nick == "":
		msg = "O campo 'nick' é obrigatório e não pode estar em branco"
	case step == SaveUser && user.Password == "":
		msg = "O campo 'password' é obrigatório e não pode estar em branco"
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		msg = "O email inserido é inválido"
	}

	if msg != "" {
		return errors.New(msg)
	}

	return nil
}

func (user *User) formatData(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)

	if step == SaveUser {
		passwordByte, err := security.CreateHash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passwordByte)
	}

	return nil
}
