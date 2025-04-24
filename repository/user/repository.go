package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository struct {
	Conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{Conn: conn}
}

func (r UserRepository) List() ([]User, error) {
	rows, err := r.Conn.Query(context.Background(), "SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
