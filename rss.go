package main

import (
  "context"
  "errors"
  "net/http"
  "fmt"
  "encoding/xml"
  "html"
  "io"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
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
  for i, items := range rssfeed.Channel.Item {
    items.Title = html.UnescapeString(items.Title)
    items.Description = html.UnescapeString(items.Description)
    rssfeed.Channel.Item[i] = items
  }

  return &rssfeed, nil
}
