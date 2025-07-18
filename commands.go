package main

import (
  "github.com/adibbelel/gator/internal/database"
  "github.com/adibbelel/gator/internal/config"
  "context"
  "errors"
  "net/http"
  "fmt"
  "encoding/xml"
  "html"
  "io"
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

func fetchFeed (ctx context.Context, feedURL string) (*RSSFeed, error) {
  var rssfeed RSSFeed

  req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
  if err != nil {
    return &rssfeed, errors.New("Error creating http request")
  }

  req.Header.Set("User-Agent", "gator")

  client := &http.Client{}
  res, err := client.Do(req)
  if err != nil {
    return &rssfeed, fmt.Errorf("Error sending request: %w", err)
  }
  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
    return &rssfeed, fmt.Errorf("Error reading response body %w", err)
  }

  if err := xml.Unmarshal(body, &rssfeed); err != nil {
    return &rssfeed, fmt.Errorf("Error unmarshalling body: %w\n", err)
  }

  rssfeed.Channel.Title = html.UnescapeString(rssfeed.Channel.Title)
  rssfeed.Channel.Description = html.UnescapeString(rssfeed.Channel.Description)
  for _, items := range rssfeed.Channel.Item {
    items.Title = html.UnescapeString(items.Title)
    items.Description = html.UnescapeString(items.Description)
  }

  return &rssfeed, nil
}
