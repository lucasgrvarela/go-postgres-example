package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(connString string) (*PostgresUserRepository, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return &PostgresUserRepository{db: db}, nil
}

func (r *PostgresUserRepository) GetByID(id int) (*User, error) {
	query := "SELECT id, username, email FROM users WHERE id=$1"
	row := r.db.QueryRow(query, id)

	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) Create(user *User) error {
	query := "INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, user.Username, user.Email).Scan(&user.ID)
	return err
}

func (r *PostgresUserRepository) Update(user *User) error {
	query := "UPDATE users SET username=$1, email=$2 WHERE id=$3"
	_, err := r.db.Exec(query, user.Username, user.Email, user.ID)
	return err
}

func (r *PostgresUserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.db.Exec(query, id)
	return err
}
