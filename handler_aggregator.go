package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gmclean3107/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <time_between_requests>", cmd.command)
	}

	timeVal, err := time.ParseDuration(cmd.args[0])

	if err != nil {
		return fmt.Errorf("error parsing time argument: %v", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeVal)

	ticker := time.NewTicker(timeVal)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func scrapeFeeds(s *State) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())

	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %v", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)

	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %v", err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)

	if err != nil {
		return fmt.Errorf("error fetching feed from url: %v", err)
	}

	for _, item := range rssFeed.Channel.Item {
		pub, err := time.Parse(item.PubDate, item.PubDate)

		if err != nil {
			fmt.Printf("error parsing published date: ", err)
		}

		post := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         feed.Url,
			Description: item.Description,
			PublishedAt: pub,
			FeedID:      feed.ID,
		}
		s.db.CreatePost(context.Background(), post)
	}

	return nil
}
