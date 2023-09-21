package user

import (
	"context"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/nicus101/scaling-giggle/db"
)

type CreateCommand struct {
	Login    string `json:"login" validate:"required"`
	Name     string `json:"name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AddPersonalDataCommand struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Pesel  int    `json:"pesel"`
}

func CreateUser(ctx context.Context, command CreateCommand) (int, error) {
	log.Println("Creating user", command)
	if err := validator.New().Struct(command); err != nil {
		return 0, fmt.Errorf("input validation failed: %w", err)
	}

	result, err := db.GetConnection().ExecContext(ctx,
		`INSERT INTO user (login, name, last_name, password) 
		VALUES (?, ?, ?, ?)`, // <3

		command.Login,    // "           "
		command.Password, // `"; DROP DATABASE *;`
		command.Name,
		command.LastName,
	)
	if err != nil {
		return 0, fmt.Errorf("database insert failed: %w", err)
	}

	lastId, err := result.LastInsertId()
	return int(lastId), err
}

func AddPersonalData(ctx context.Context, command AddPersonalDataCommand) error {
	log.Println("Adding personal data", command)

	return nil
}
