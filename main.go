package main

import (
	"database/sql"
	"fmt"
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
		log.Fatal(fmt.Sprintf("error while establishing a databse connection: %v", err))
	}
	dbQueries := database.New(db)

	programState := &state{config: configFile, db: dbQueries}

	coms := commands{regComs: make(map[string]func(*state, command) error)}
	coms.register("register", handlerRegister)
	coms.register("login", handlerLogin)
	coms.register("reset", handlerReset)
	coms.register("users", handlerUsers)
	coms.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	coms.register("feeds", handlerFeeds)
	coms.register("follow", middlewareLoggedIn(handlerFollow))
	coms.register("following", middlewareLoggedIn(handlerFollowing))
	coms.register("agg", middlewareAgg(handlerAgg))
	coms.register("unfollow", middlewareLoggedIn(handlerUnfollow))

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
