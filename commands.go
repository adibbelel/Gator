package main

import (
  "github.com/adibbelel/gator/internal/config"
  "errors"
  "fmt"
)

type state struct {
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

  err := s.cfg.SetUser(name)
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

func (c *commands) register(name string, f func(*state, command) error) {
  c.registeredCommands[name] = f
}
