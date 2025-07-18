package main

import (
  "github.com/adibbelel/gator/internal/database"
  "context"
  "errors"
  "github.com/google/uuid"
  "time"
  "fmt"
)

func handlerLogin(s *state, cmd command) error {
  if len(cmd.inputs) != 1 {
     return fmt.Errorf("wrong usage")
  }

  name := cmd.inputs[0]

  _, err := s.db.GetUser(context.Background(), name)
  if err != nil {
    return fmt.Errorf("Could not get User: ", err)
  }

  err = s.cfg.SetUser(name)
  if err != nil {
    return fmt.Errorf("Failed to set user") 
  }
  fmt.Println("User has been set to -", s.cfg.CurrentUserName)
  return nil
}

func handlerRegister(s *state, cmd command) error {
  if len(cmd.inputs) != 1 {
     return fmt.Errorf("wrong usage")
  }

  name := cmd.inputs[0]

  user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name})
  if err != nil {
    return fmt.Errorf("failed to create new user: ", err)
  }

  err = s.cfg.SetUser(user.Name)
  if err != nil {
    return fmt.Errorf("Failed to set user") 
  }
  fmt.Println("User has been set to -", s.cfg.CurrentUserName)
  return nil
}

func handlerReset (s *state, cmd command) error {
  err := s.db.ResetFeeds(context.Background())
  if err != nil {
    return errors.New("Could not reset Feeds table state")
  }

  err = s.db.ResetTable(context.Background())
  if err != nil {
    return errors.New("Could not reset Users table state")
  }

  fmt.Println("Table has been successfully reset")
  return nil
}

func handlerGetUsers (s *state, cmd command) error {
  users, err := s.db.GetUsers(context.Background())
  if err != nil {
    return errors.New("Could not get user data")
  }

  for _, user := range users {
    if user.Name == s.cfg.CurrentUserName {
      fmt.Printf("%s (current) \n", user.Name)
    } else {
      fmt.Printf("%s\n", user.Name)
    }
  }

  return nil
}

