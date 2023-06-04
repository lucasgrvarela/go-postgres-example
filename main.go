package main

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	userdomain "github.com/lucasgrvarela/go-postgres-example/user"
	"github.com/lucasgrvarela/go-postgres-example/user/user_repository"
)

func main() {
	connString := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=postgres sslmode=disable"
	repo, err := user_repository.NewPostgresUserRepository(connString)
	if err != nil {
		fmt.Println("Failed to create UserRepository:", err)
		return
	}

	// Use the UserRepository methods
	newUser := &userdomain.User{
		Username: "john_doe",
		Email:    "john.doe@example.com",
	}
	err = repo.Create(newUser)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("User created with success")

	user, err := repo.GetByID(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("User:", user)

	user, err = repo.GetByID(2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("User:", user)

	newUser.Username = "johndoe"
	newUser.Email = "johndoe@example.com"
	err = repo.Update(newUser)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("User updated with success")

	err = repo.Delete(newUser.ID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("User deleted with success")
}
