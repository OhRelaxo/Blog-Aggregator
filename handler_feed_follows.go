package main

import (
	"context"
	"fmt"
	"time"

	"github.com/OhRelaxo/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("you need a url that you want to follow")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("error while geting feed: %v", err)
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error while executing CreateFeedFollow: %v", err)
	}
	fmt.Printf("- username: %v\n- feedname: %v\n", feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handlerFollowing(s *state, _ command, _ database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), *s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error while executing the GetFeedFollowsForUser command: %v", err)
	}
	for _, feedFollow := range feedFollows {
		fmt.Printf("- feedname: %v\n", feedFollow.FeedName.String)
	}
	return nil
}
