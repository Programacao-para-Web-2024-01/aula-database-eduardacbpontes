package user

import (
	"database/sql"
	"errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Register(user User) (int64, error) {
	result, err := r.db.Exec(
		"INSERT INTO users (username, password) VALUES (?, SHA2(?, 256))",
		user.Username, user.Password)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) Login(username, password string) error {
	row := r.db.QueryRow(
		"SELECT COUNT(id) FROM users WHERE username = ? AND password = SHA2(?, 256)",
		username, password,
	)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("user/password does not match")
	}

	return nil
}
