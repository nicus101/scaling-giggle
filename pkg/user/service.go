package user

import (
	"context"
	"log"
	"math/rand"
)

type CreateCommand struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

type AddPersonalDataCommand struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Pesel  int    `json:"pesel"`
}

func CreateUser(ctx context.Context, command CreateCommand) (int, error) {
	log.Println("Creating user", command)

	return rand.Int(), nil
}

func AddPersonalData(ctx context.Context, command AddPersonalDataCommand) error {
	log.Println("Adding personal data", command)

	return nil
}
