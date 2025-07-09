package main

import (
  "log"
  "github.com/adibbelel/gator/internal/config"
  "os"
)

func main()  {
  username := "adib"
  var newState state

  newState.cfg, err := config.Read()
  if err != nil {
    log.Fatalf("Error creating new config", err)
  }

  cmds := commands{
    registeredCommands: make(map[string]func(*state, commmand)error),
  }

  cmds.register("login", handlerLogin)

  newConfig.SetUser(username)
  newConfig, err = config.Read()
  if err != nil {
    log.Fatalf("Error rereading config")
  }
  
  fmt.Printf("%s, %s\n", newConfig.DbURL, newConfig.CurrentUserName)
}
