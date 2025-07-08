package config

import (
  "os"
  "encoding/json"
  "io/ioutil"
  "fmt"
)

type Config struct {
  DbURL           string `json:"db_url"`
  CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    return nil, fmt.Errorf("Error accessing home directory")
  }
  data :=  + configFileName
  jsonData, err := ioutil.ReadFile(data)
  if err != nil {
    return nil, fmt.Errorf("Error reading file")
  }

  var config Config
  err := json.Unmarshal([]byte(jsonData), &config)
  if err != nil {
    return nil, fmt.Errorf("could not unmarshal jsonData")
  }

  return config, nil
}

func (c *Config) SetUser(username string) {
  c.CurrentUserName := username
  updatedData, err := json.Marshal(c)
  if err != nil {
    fmt.Errorf("Could not Marshal Data")
  }

  err = json.WriteFile((os.UserHomeDir() + configFileName), updatedData, 0644)
  if err != nil {
    fmt.Errorf("Could not update data in json file")
    return 
  }
}
