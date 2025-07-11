package main

import (
  "log"
  "github.com/adibbelel/gator/internal/config"
  "os"
)

func main()  {
  cfg, err := config.Read()
  if err != nil {
    log.Fatalf("Error reading config", err)
  }

  newState := &state{
    cfg: &cfg,
  }

  cmds := commands{
    registeredCommands: make(map[string]func(*state, command)error),
  }

  cmds.register("login", handlerLogin)

  if len(os.Args) < 2 {
    log.Fatal("usage: cli <command>")
  }

  cmdName := os.Args[1]
  cmdArgs := os.Args[2:]

  err = cmds.run(newState, command{name: cmdName, inputs: cmdArgs})
  if err != nil {
    log.Fatal(err)
  }
  
}
