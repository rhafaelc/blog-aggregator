package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rhafaelc/blog-aggregator/internal/database"
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
	return scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) error {
	err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark feed %s fetched: %w", feed.Name, err)
	}

	fetchedFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't get feed %s: %w", feed.Name, err)
	}
	for _, item := range fetchedFeed.Channel.Item {
		timeNow := time.Now().UTC()

		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		post, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			return fmt.Errorf("couldn't add post %s: %w", item.Title, err)
		}

		fmt.Printf("Found post: %s\n", post.Title)
	}
	fmt.Printf("Feed %s collected, %v posts found", feed.Name, len(fetchedFeed.Channel.Item))
	return nil
}

