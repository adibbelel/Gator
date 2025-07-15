package main

import (
  "github.com/adibbelel/gator/internal/database"
  "github.com/adibbelel/gator/internal/config"
  "context"
  "errors"
  "github.com/google/uuid"
  "time"
  "fmt"
)

type state struct {
  db *database.Queries
  cfg *config.Config 
}

type command struct {
  name string
  inputs []string
}

type commands struct {
  registeredCommands map[string]func(*state, command) error
}

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

func (c *commands) run(s *state, cmd command) error {
  function, ok := c.registeredCommands[cmd.name]
  if !ok {
    return errors.New("command not found")
  }
  return function(s, cmd)
}

func handlerReset (s *state, cmd command) error {
  err := s.db.ResetTable(context.Background())
  if err != nil {
    return errors.New("Could not reset database state")
  }
  fmt.Println("Table has been successfully reset")

  return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
  c.registeredCommands[name] = f
}
