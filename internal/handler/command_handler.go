package commandHandler

import (
	"fmt"

	config "github.com/OhRelaxo/Blog-Aggregator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name      string
	arguments []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("the login command expects at least one argument")
	}
	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("The User: %v has been set", cmd.arguments[0])
	return nil
}
