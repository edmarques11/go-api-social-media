package models

import (
	"errors"
	"strings"
	"time"
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
func (user *User) Prepare() error {
	if err := user.validator(); err != nil {
		return err
	}
	user.formatData()
	return nil
}

func (user *User) validator() error {
	var msg string

	switch {
	case user.Name == "":
		msg = "o campo 'name' é obrigatório e não pode estar em branco"
	case user.Email == "":
		msg = "o campo 'email' é obrigatório e não pode estar em branco"
	case user.Nick == "":
		msg = "o campo 'nick' é obrigatório e não pode estar em branco"
	case user.Password == "":
		msg = "o campo 'password' é obrigatório e não pode estar em branco"
	}

	if msg != "" {
		return errors.New(msg)
	}

	return nil
}

func (user *User) formatData() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)
}
