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

	fmt.Printf("successfully added feed: %v\n", dbFeed)

	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)

	if err != nil {
		return fmt.Errorf("error fetching current user: %v", err)
	}

	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	s.db.CreateFeedFollow(context.Background(), feedFollow)

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

func handlerFollow(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <feed_url>", cmd.command)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)

	if err != nil {
		return fmt.Errorf("error fetching current user: %v", err)
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])

	if err != nil {
		return fmt.Errorf("error fetching feed with supplied url: %v", err)
	}

	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	row, err := s.db.CreateFeedFollow(context.Background(), feedFollow)

	if err != nil {
		return fmt.Errorf("error following feed: %v", err)
	}

	fmt.Printf("successfully followed feed: %v", row)

	return nil
}

func handlerFollowing(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %v", cmd.command)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)

	if err != nil {
		return fmt.Errorf("error fetching user: %v", err)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return fmt.Errorf("error getting feeds for supplied user: %v", err)
	}

	fmt.Printf("%v Followed Feeds:\n", s.cfg.Current_user_name)

	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}
