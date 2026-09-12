package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *State, cmd Command) error {
	// if len(cmd.args) != 1 {
	// 	return fmt.Errorf("usage: %v <feed_url>", cmd.command)
	// }

	// feed, err := rss.FetchFeed(context.Background(), cmd.args[0])

	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return fmt.Errorf("error getting rss feed: %v", err)
	}

	fmt.Println(feed)

	return nil
}
