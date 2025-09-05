package main

import (
	"log"
	"os"

	"github.com/OhRelaxo/Blog-Aggregator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	configFile, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	programState := &state{config: configFile}

	coms := commands{regComs: make(map[string]func(*state, command) error)}
	coms.register("login", handlerLogin)

	input := os.Args
	if len(input) < 2 {
		log.Fatal("you need to input a command name")
	}

	name := input[1]
	args := input[2:]
	cmd := command{name: name, arguments: args}
	err = coms.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
