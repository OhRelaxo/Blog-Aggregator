package main

import (
	"context"
	"fmt"

	"github.com/OhRelaxo/Blog-Aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		username := s.config.CurrentUserName
		if username == nil {
			return fmt.Errorf("no user loggedin")
		}

		user, err := s.db.GetUser(context.Background(), *username)
		if err != nil {
			return fmt.Errorf("error while getting a user: %v", err)
		}

		err = handler(s, cmd, user)
		if err != nil {
			return err
		}

		return nil
	}
}
