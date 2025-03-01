package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAggregator(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage %v <time_between_request>", cmd.Name)
	}

	time_between_request, err := time.ParseDuration(cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("couldn't parse duration for time between request: %w", err)
	}
	fmt.Printf("collecting feeds every %v\n", time_between_request)
	ticker := time.NewTicker(time_between_request)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feed to be fetched: %w", err)
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	fetchedFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	for _, item := range fetchedFeed.Channel.Item {
		fmt.Printf("Title:		%v\n", item.Title)
		fmt.Printf("==============================\n")
	}
	return nil
}
