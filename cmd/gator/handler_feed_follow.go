package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rhafaelc/blog-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("usage %v", cmd.Name)
	}

	feedFollowsForUser, err := s.db.ListFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follow list: %w", err)
	}

	for _, feedFollow := range feedFollowsForUser {
		printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
		fmt.Println()
		fmt.Println("==================================")
	}

	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage %v <url>", cmd.Name)
	}

	url := cmd.Arguments[0]

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

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage %v <url>", cmd.Name)
	}

	url := cmd.Arguments[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't delete feed follow: %w", err)
	}

	fmt.Println("feed unfollowed")
	printFeedFollow(user.Name, feed.Url)
	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("User Name: %s\n", username)
	fmt.Printf("Feed Name: %s\n", feedname)
}
