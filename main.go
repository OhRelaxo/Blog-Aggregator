package main

import (
	"log"
	"os"

	"github.com/OhRelaxo/Blog-Aggregator/internal/config"
	"github.com/OhRelaxo/Blog-Aggregator/internal/handler"
)

//const userName = "Marcel"

//continue at step 6

func main() {
	configFile, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	state := commandHandler.State{Config: configFile}

	commands := commandHandler.Commands{Coms: make(map[string]func(*commandHandler.State, commandHandler.Command) error)}
	commands.Register("login", commandHandler.HandlerLogin)

	input := os.Args
	if len(input) < 2 {
		log.Fatal("you need to input a command name")
	}

	name := input[1]
	args := input[2:]
	cmd := commandHandler.Command{Name: name, Arguments: args}
	err = commands.Run(&state, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
