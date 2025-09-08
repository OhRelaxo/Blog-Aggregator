package main

import (
	"context"
	"fmt"
	"log"
)

func handlerAgg(_ *state, _ command) error {
	/*
		if len(cmd.arguments) < 1 {
			return fmt.Errorf("u need to pass a link to agg")
		}
	*/
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		log.Fatal(fmt.Sprintf("error while fetching the feed: %v\n", err))
	}
	fmt.Println(*feed)
	return nil
}
