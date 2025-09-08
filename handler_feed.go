package main

import (
	"context"
	"fmt"
	"time"

	"github.com/OhRelaxo/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("you need at two arguments to add a feed")
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(),
		Name: cmd.arguments[0], Url: cmd.arguments[1], UserID: user.ID})
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed follow: %v", err)
	}
	fmt.Println("created feed successfully")
	fmt.Printf("- Feed: %v\n", feed)
	fmt.Println("created feed follow successfully")
	fmt.Printf("- Feed Follow: %v\n", feedFollow)
	return nil
}

func handlerFeeds(s *state, _ command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("- feed name: %v\n- feed url: %v\n- user name: %v\n", feed.Name, feed.Url, feed.UserName.String)
	}
	return nil
}
