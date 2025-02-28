package main

import (
	"context"
	"fmt"
)



func handlerAggregator(s *state, cmd command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("usage %v", cmd.Name)
	}
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}
