package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gmclean3107/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *State, cmd Command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <feed_name> <feed_url>", cmd.command)
	}

	feedName := cmd.args[0]
	feedUrl := cmd.args[1]
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)

	if err != nil {
		return fmt.Errorf("error fetching current user: %v", err)
	}

	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    currentUser.ID,
	}

	dbFeed, err := s.db.CreateFeed(context.Background(), feed)

	if err != nil {
		return fmt.Errorf("error creating new feed: %v", err)
	}

	fmt.Println(dbFeed)

	return nil
}

func handlerGetFeeds(s *State, cmd Command) error {

	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %v", cmd.command)
	}

	feeds, err := s.db.GetFeeds(context.Background())

	if err != nil {
		return fmt.Errorf("error getting list of feeds: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %v\n", feed.Name)
		fmt.Printf("Url: %v\n", feed.Url)
		fmt.Printf("User: %v\n", feed.UserName)
	}

	return nil
}
