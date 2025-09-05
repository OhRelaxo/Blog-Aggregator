package main

import (
	"context"
	"fmt"
	"time"

	"github.com/OhRelaxo/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("u need to pass a name to register")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.arguments[0]})
	if err != nil {
		return err
	}
	err = s.config.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("the user: %v was successfuly created\n", user.Name)
	fmt.Printf("LOG in handlerRegiser: ID: %v, CreatedAt: %v, UpdatedAt: %v, Name: %v", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)
	return nil
}
