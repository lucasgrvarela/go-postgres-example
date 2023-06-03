package main

import "github.com/lucasgrvarela/models"

type UserRepository interface {
	GetByID(id int) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id int) error
}
