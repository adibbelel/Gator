package main

import _ "github.com/lib/pq"

import (
  "context"
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
  cmds.register("users", handlerGetUsers)
  cmds.register("agg", handlerAgg)
  cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
  cmds.register("feeds", handlerFeeds)
  cmds.register("follow", middlewareLoggedIn(handlerFollow))
  cmds.register("following", middlewareLoggedIn(handlerFollowing))
  cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
  cmds.register("browse", middlewareLoggedIn(handlerBrowse))

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

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
