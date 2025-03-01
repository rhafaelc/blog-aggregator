package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rhafaelc/blog-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) == 0 {
		cmd.Arguments = append(cmd.Arguments, "2")
	}
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage %v <limit (default 2)>", cmd.Name)
	}

	limit, err := strconv.ParseInt(cmd.Arguments[0], 10, 32)
	if err != nil {
		return fmt.Errorf("couldn't parse int: %w", err)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
