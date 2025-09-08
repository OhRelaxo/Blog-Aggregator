package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func handlerAgg(s *state, _ command, timeBetweenReqs string) error {
	duration, err := time.ParseDuration(timeBetweenReqs)
	if err != nil {
		return err
	}
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error while GetNextFeedToFetch: %v", err)
	}
	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("error while MarkFeedFetched: %v", err)
	}
	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		log.Fatal(fmt.Sprintf("error while fetching the feed: %v\n", err))
	}
	fmt.Printf("========================%v========================\nDescription:%v\n", feed.Channel.Title, feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Printf("------------------------%v------------------------\nDescription:%v\n", item.Title, item.Description)
	}
	return nil
}
