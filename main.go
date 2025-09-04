package main

import (
	"fmt"
	"log"

	config "github.com/OhRelaxo/Blog-Aggregator/internal/config"
)

const userName = "Marcel"

//continue at step 6

func main() {
	configFile, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = configFile.SetUser(userName)
	if err != nil {
		log.Fatal(err)
	}

	newConfigFile, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Db_url: %v\nCurrent_user_name: %v\n", *newConfigFile.Db_url, *newConfigFile.Current_user_name)
}
