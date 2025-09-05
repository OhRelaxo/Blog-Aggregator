package commandHandler

import "fmt"

type Commands struct {
	Coms map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	if f, ok := c.Coms[cmd.Name]; ok {
		err := f(s, cmd)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("no command found: %v", cmd.Name)
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Coms[name] = f
}
