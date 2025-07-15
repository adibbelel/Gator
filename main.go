package main

import _ "github.com/lib/pq"

import (
  "log"
  "github.com/adibbelel/gator/internal/config"
  "github.com/adibbelel/gator/internal/database"
  "database/sql"
  "os"
)

func main()  {
  cfg, err := config.Read()
  if err != nil {
    log.Fatalf("Error reading config", err)
  }

  dbURL := cfg.DbURL 
  db, err := sql.Open("postgres", dbURL)
  dbQueries := database.New(db)


  newState := &state{
    db: dbQueries,
    cfg: &cfg,
  }
  

  cmds := commands{
    registeredCommands: make(map[string]func(*state, command)error),
  }

  cmds.register("login", handlerLogin)
  cmds.register("register", handlerRegister)
  cmds.register("reset", handlerReset)

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
