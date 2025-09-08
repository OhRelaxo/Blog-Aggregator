package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/OhRelaxo/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) < 1 || len(cmd.arguments) > 2 {
		return fmt.Errorf("please use a time like 10s as your argument")
	}

	duration, err := time.ParseDuration(cmd.arguments[0])
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
	for _, item := range feed.Channel.Item {
		timeStamp, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			log.Fatalf("time has defited me :(: %v", err)
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Title: item.Description, Url: item.Link, Description: item.Description, PublishedAt: timeStamp, FeedID: nextFeed.ID})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	return nil
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int
	var err error
	if len(cmd.arguments) != 1 {
		limit = 2
	} else {
		limit, err = strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return err
		}
	}
	posts, err := s.db.GetPosts(context.Background(), database.GetPostsParams{UserID: user.ID, Limit: int32(limit)})
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Printf("- Title: %v\n- Description: %v\n- Link %v\n", post.Title, post.Description, post.Url)
	}
	return nil
}
