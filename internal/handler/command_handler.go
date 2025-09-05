package commandHandler

import (
	"fmt"

	"github.com/OhRelaxo/Blog-Aggregator/internal/config"
)

type State struct {
	Config *config.Config
}

type Command struct {
	Name      string
	Arguments []string
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("the login command expects at least one argument")
	}
	err := s.Config.SetUser(cmd.Arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("The User: %v has been set", cmd.Arguments[0])
	return nil
}
