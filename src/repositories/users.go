package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// NewUserRepository create a new user repository
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

// Create save a new user into database
func (userRepository users) Create(userModel models.User) (uint64, error) {
	statment, err := userRepository.db.Prepare(
		"insert into tb_user (name, nick, email, password) values (?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statment.Close()

	result, err := statment.Exec(
		userModel.Name,
		userModel.Nick,
		userModel.Email,
		userModel.Password,
	)
	if err != nil {
		return 0, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(userId), nil
}
