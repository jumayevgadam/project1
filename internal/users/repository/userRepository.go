package repository

import (
	"Project1/internal/users/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

const (
	Users    = "users"
	GetUsers = "users"
)

func NewUserRepository(DB *sqlx.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (r *UserRepository) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (name, username, password) VALUES ($1, $2, $3) RETURNING id`, Users)

	row := r.DB.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf(`SELECT * FROM %s WHERE username = $1 and password = $2`, Users)

	err := r.DB.Get(&user, query, username, password)

	if err != nil {
		return model.User{}, err
		fmt.Errorf("error getting user: %w", err)
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(userId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s WHERE id = $1`, Users)
	_, err := r.DB.Exec(query, userId)
	return err
}
