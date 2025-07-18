package main

import (
  "github.com/adibbelel/gator/internal/database"
  "context"
  "github.com/google/uuid"
  "time"
  "fmt"
)

func handlerAgg (s *state, cmd command) error {
  url := "https://www.wagslane.dev/index.xml"
  rssfeed, err := fetchFeed(context.Background(), url)
  if err != nil {
    return fmt.Errorf("Error fetching RSSFeed %w\n", err)
  }

  fmt.Printf("%v", rssfeed)
  return nil
}

func handlerAddFeed (s *state, cmd command) error {
  if len(cmd.inputs) != 2 {
     return fmt.Errorf("wrong usage")
  }

  user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
  if err != nil {
    return fmt.Errorf("Error getting user: %w\n", err)
  }

  name := cmd.inputs[0]
  url := cmd.inputs[1]

  feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name, Url: url, UserID: user.ID})
  if err != nil {
    return fmt.Errorf("Error creating feed: %w", err)
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

