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

const configFileName = "/.gatorconfig.json"

func Read() (Config, error) {
  var config Config
  homeDir, err := os.UserHomeDir()
  if err != nil {
    return config, fmt.Errorf("Error accessing home directory")
  }
  data := homeDir + configFileName
  jsonData, err := ioutil.ReadFile(data)
  if err != nil {
    return config, fmt.Errorf("Error reading file")
  }

  err = json.Unmarshal([]byte(jsonData), &config)
  if err != nil {
    return config, fmt.Errorf("could not unmarshal jsonData")
  }

  return config, nil
}

func (c *Config) SetUser(username string) error {
  c.CurrentUserName = username
  updatedData, err := json.Marshal(c)
  if err != nil {
    fmt.Errorf("Could not Marshal Data")
  }

  homeDir, err := os.UserHomeDir()
  if err != nil {
    return fmt.Errorf("Error accessing home directory")
  }
  data := homeDir + configFileName

  err = ioutil.WriteFile(data, updatedData, 0644)
  if err != nil {
    return fmt.Errorf("Could not update data in json file") 
  }

  return nil
}
