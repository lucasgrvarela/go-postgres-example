package user_repository

import "github.com/lucasgrvarela/go-postgres-example/user"

type UserRepository interface {
	GetByID(id int) (*user.User, error)
	Create(user *user.User) error
	Update(user *user.User) error
	Delete(id int) error
}
