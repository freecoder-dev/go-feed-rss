package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Feed represents an RSS feed item
type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// Channel represents the channel of an RSS feed
type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

// Item represents an item in an RSS feed
type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
}

func main() {
	// Example RSS feed URL
	url := "https://freecoder.dev/feed/"

	// Fetch the RSS feed
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to fetch RSS feed:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		os.Exit(1)
	}

	// Parse the RSS feed
	var feed Feed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		fmt.Println("Failed to parse RSS feed:", err)
		os.Exit(1)
	}

	// Print the feed title and description
	fmt.Println("************************************")
	fmt.Println("Channel Title:", feed.Channel.Title)
	fmt.Println("Description:", feed.Channel.Description)
	fmt.Println("************************************")
	fmt.Println()

	// Print the latest items
	for _, item := range feed.Channel.Items {
		fmt.Println("Title:", item.Title)
		fmt.Println("Description:", item.Description)
		fmt.Println("Link:", item.Link)
		fmt.Println("--------")
	}
}
