package main

import (
  "github.com/adibbelel/gator/internal/database"
  "context"
  "github.com/google/uuid"
  "time"
  "fmt"
)

func handlerAddFeed (s *state, cmd command, user database.User) error {
  if len(cmd.inputs) != 2 {
     return fmt.Errorf("wrong usage")
  }

  name := cmd.inputs[0]
  url := cmd.inputs[1]

  feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name, Url: url, UserID: user.ID})
  if err != nil {
    return fmt.Errorf("Error creating feed: %w", err)
  }

  _, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID})
  if err != nil {
    return fmt.Errorf("Error creating feed follow: %w", err)
  }

  fmt.Printf("feed name: %s, feed URL: %s\n", feed.Name, feed.Url)
  return nil
}

func handlerFeeds (s *state, cmd command) error {
  feeds, err := s.db.GetFeeds(context.Background())
  if err != nil {
    return fmt.Errorf("Error getting feeds from database: %w", err)
  }

  for _, feed := range feeds {
    fmt.Printf("%s\n", feed.Name)
    fmt.Printf("%s\n", feed.Url)
    fmt.Printf("%s\n", feed.Name_2)
  }
  return nil
}

