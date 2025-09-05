package main

import "fmt"

type command struct {
	name      string
	arguments []string
}
type commands struct {
	regComs map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if f, ok := c.regComs[cmd.name]; ok {
		err := f(s, cmd)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("no command found: %v", cmd.name)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.regComs[name] = f
}
