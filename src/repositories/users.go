package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// GetUsers search users by name or nick
func (userRepository users) GetUsers(nameOrNIck string) ([]models.User, error) {
	nameOrNIck = fmt.Sprintf("%%%s%%", nameOrNIck)
	rows, err := userRepository.db.Query(
		"select id, name, email, nick, created_at from tb_user where name like ? or nick like ?",
		nameOrNIck,
		nameOrNIck,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Nick,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetById search an user by id
func (userRepository users) GetById(userId int64) (models.User, error) {
	rows, err := userRepository.db.Query(
		"select id, name, email, nick, created_at from tb_user where id = ?",
		userId,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Nick,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// GetByEmail search an user by email
func (userRepository users) GetByEmail(userEmail string) (models.User, error) {
	rows, err := userRepository.db.Query(
		"select id, password from tb_user where email = ?",
		userEmail,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// UpdateUser update an user
func (userRepository users) UpdateUser(userId int64, user models.User) error {
	statment, err := userRepository.db.Prepare(
		"update tb_user set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statment.Close()

	if _, err = statment.Exec(user.Name, user.Nick, user.Email, userId); err != nil {
		return err
	}

	return nil
}

// DeleteUser delete an user
func (userRepository users) DeleteUser(userId int64) error {
	statment, err := userRepository.db.Prepare(
		"delete from tb_user where id = ?",
	)
	if err != nil {
		return err
	}
	defer statment.Close()

	if _, err = statment.Exec(userId); err != nil {
		return err
	}

	return nil
}
