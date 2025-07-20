package main

import (
  "github.com/adibbelel/gator/internal/database"
  "context"
  "strconv"
  "fmt"
)

func handlerBrowse (s *state, cmd command, user database.User) error {
  limit := 2
  if len(cmd.inputs) == 1 {
    if newlimit, err := strconv.Atoi(cmd.inputs[0]); err == nil {
      limit = newlimit
    } else {
      return fmt.Errorf("invalid limit: %w\n", err)
    }
  }

  posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
    UserID: user.ID,
    Limit: int32(limit),
  })
  if err != nil {
    return fmt.Errorf("Error getting posts for %s: %w\n", user.Name, err)
  }

  fmt.Printf("Found %d posts for User: %s\n", len(posts), user.Name)

  for _, post := range posts {
    fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %s\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
    fmt.Println("=========================================================================")
  }

  return nil
}

