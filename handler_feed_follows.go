package main

import (
  "github.com/adibbelel/gator/internal/database"
  "context"
  "github.com/google/uuid"
  "time"
  "fmt"
)

func handlerFollow (s *state, cmd command, user database.User) error {
  if len(cmd.inputs) != 1 {
     return fmt.Errorf("wrong usage")
  }

  Url := cmd.inputs[0]

  feed, err := s.db.GetFeed(context.Background(), Url)
  if err != nil {
    return fmt.Errorf("Error getting feed: %w\n", err)
  }

  follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID})
  if err != nil {
    return fmt.Errorf("Error creating feed: %w", err)
  }

  fmt.Printf("feed name: %s, User: %s\n", follow.FeedName, follow.UserName)
  return nil
}

func handlerFollowing (s *state, cmd command, user database.User) error {
  ffu, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
  if err != nil {
    return fmt.Errorf("Error getting feed follows: %w\n", err)
  }

  for _, followUser := range ffu {
    fmt.Printf("%s", followUser.FeedName)
  }
  
  return nil
}

 
func handlerUnfollow (s *state, cmd command, user database.User) error {
  if len(cmd.inputs) != 1 {
     return fmt.Errorf("wrong usage")
  }

  Url := cmd.inputs[0]

  err := s.db.DeleteFeedFollow(context.Background(), Url)
  if err != nil {
    return fmt.Errorf("Error Unfollowing feed: ", err)
  }

  return nil
}
