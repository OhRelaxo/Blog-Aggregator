package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) < 0 {
		return fmt.Errorf("the login command expects at least one argument")
	}
	user, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Printf("The User: %v has been logedin", user.Name)
	return nil
}
