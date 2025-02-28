package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rhafaelc/blog-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("usage %v", cmd.Name)
	}
	username := s.cfg.CurrentUserName

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}

	feedFollowsForUser, err := s.db.ListFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follow list: %w", err)
	}

	for _, feedFollow:= range feedFollowsForUser {
		printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
		fmt.Println()
		fmt.Println("==================================")
	}

	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage %v <url>", cmd.Name)
	}

	username := s.cfg.CurrentUserName
	url := cmd.Arguments[0]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	timeNow := time.Now().UTC()
	feedFollow, err := s.db.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
			UserID:    user.ID,
			FeedID:    feed.ID,
		})
	if err != nil {
		return fmt.Errorf("couldn't add feed follow: %w", err)
	}

	fmt.Println("feed follow added")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("============================")

	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("User Name: %s\n", username)
	fmt.Printf("Feed Name: %s\n", feedname)
}
