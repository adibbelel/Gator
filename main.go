package main

import (
  "fmt"
  "github.com/adibbelel/gator/internal/config"
)

func main()  {
  username := "adib"
  var newState state

  newState.cfg, err := config.Read()
  if err != nil {
    fmt.Errorf("Error creating new config")
  }

  cmds := commands{
    registeredCommands: make(map[string]func(*state, commmand))
  }
  newConfig.SetUser(username)
  newConfig, err = config.Read()
  if err != nil {
    fmt.Errorf("Error rereading config")
  }
  
  fmt.Printf("%s, %s\n", newConfig.DbURL, newConfig.CurrentUserName)
}
