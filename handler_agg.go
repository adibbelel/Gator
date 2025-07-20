package main 

import (
  "github.com/adibbelel/gator/internal/database"
  "context"
  "github.com/google/uuid"
  "strings"
  "time"
  "fmt"
  "log"
  "database/sql"
)

func handlerAgg (s *state, cmd command) error {
  if len(cmd.inputs) != 1 {
    return fmt.Errorf("wrong usage: %v <time_between_reqs>", cmd.name)
  }

  time_between_reqs, err := time.ParseDuration(cmd.inputs[0])
  if err != nil {
    return fmt.Errorf("invalid duration: %w\n", err)
  }
  log.Printf("Collecting feeds every %s\n", time_between_reqs)

  ticker := time.NewTicker(time_between_reqs)
  for ; ; <-ticker.C {
    scrapeFeeds(s)
  }

  return nil
}

func scrapeFeeds (s *state) {
  feed, err := s.db.GetNextFeedToFetch(context.Background())
  if err != nil {
    log.Printf("Error fetching next feed: %w\n", err)
  }

  log.Println("Found a feed")

  _, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
  if err != nil {
    log.Printf("Error marking feed: %w\n", err)
  }

  rssFeed, err := fetchFeed(context.Background(), feed.Url)
  if err != nil {
    log.Printf("Error fetching RSS Feed: %w\n", err)
  }

  for _, item := range rssFeed.Channel.Item {
    publishedAt := sql.NullTime{}
    if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
      publishedAt = sql.NullTime{
        Time: t,
        Valid: true,
      }
    }
    _, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
      ID: uuid.New(),
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
      Title: item.Title,
      Url: item.Link,
      Description: item.Description,
      FeedID: feed.ID,
      PublishedAt: publishedAt,
    })
    
    if err != nil {
      if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
        continue
      }
      log.Printf("Error creating post: %w\n", err)
    }
  }

  log.Println("Feed %s collected", feed.Name)
}


