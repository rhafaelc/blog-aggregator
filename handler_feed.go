package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rhafaelc/blog-aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Arguments) != 2 {
		return fmt.Errorf("usage %v <title> <url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}

	name := cmd.Arguments[0]
	url := cmd.Arguments[1]
	timeNow := time.Now().UTC()

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't add feed: %w", err)
	}
	fmt.Println("feed added")
	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:      %v\n", feed.ID)
	fmt.Printf(" * Created:      %v\n", feed.CreatedAt)
	fmt.Printf(" * Updated:      %v\n", feed.UpdatedAt)
	fmt.Printf(" * Name:      %v\n", feed.Name)
	fmt.Printf(" * URL:				%v\n", feed.Url)
	fmt.Printf(" * UserID:		%v\n", feed.UserID)
}
