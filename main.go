package main

import (
  "fmt"
  "github.com/adibbelel/gator/internal/config"
)

func main()  {
  username := "adib"

  newConfig, err := config.Read()
  if err != nil {
    fmt.Errorf("Error creating new config")
  }
  newConfig.SetUser(username)
  newConfig, err = config.Read()
  if err != nil {
    fmt.Errorf("Error rereading config")
  }
  
  fmt.Printf("%s, %s\n", newConfig.DbURL, newConfig.CurrentUserName)
}
