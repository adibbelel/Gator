package main

import (
  "github.com/adibbelel/gator/internal/database"
  "github.com/adibbelel/gator/internal/config"
  "errors"
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

