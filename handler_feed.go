package main

import (
	"context"
	"fmt"
	"time"

	"github.com/OhRelaxo/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("you need at two arguments to add a feed")
	}
	user, err := s.db.GetUser(context.Background(), *s.config.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(),
		Name: cmd.arguments[0], Url: cmd.arguments[1], UserID: user.ID})
	if err != nil {
		return err
	}
	fmt.Printf("Feed: %v\n", feed)
	return nil
}
