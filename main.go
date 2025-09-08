package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/OhRelaxo/Blog-Aggregator/internal/config"
	"github.com/OhRelaxo/Blog-Aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

func main() {
	configFile, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", *configFile.DbUrl)
	if err != nil {
		log.Fatalf("error while establishing a databse connection: %v", err)
	}
	dbQueries := database.New(db)

	programState := &state{config: configFile, db: dbQueries}

	coms := commands{regComs: make(map[string]func(*state, command) error)}
	coms.register("login", handlerLogin)
	coms.register("register", handlerRegister)
	coms.register("reset", handlerReset)
	coms.register("users", handlerUsers)
	coms.register("addfeed", handlerAddFeed)

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
