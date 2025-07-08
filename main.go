package main

import (
  "fmt"
  "github.com/adibbelel/gator/internal/config"
)

func main()  {
  username := "adib"

  newConfig := config.Read()
  newConfig.config.SetUser(username)
  
  fmt.Printf("%", config.Read)
}
