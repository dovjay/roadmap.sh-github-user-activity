package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type GitHubEvent struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`
	} `json:"payload"`
}

func fetchEvents(username string) ([]GitHubEvent, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var events []GitHubEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}
	return events, nil
}

func displayEvents(events []GitHubEvent, filter string) {
	for _, event := range events {
		if filter != "" && !strings.EqualFold(event.Type, filter) {
			continue
		}

		switch event.Type {
		case "PushEvent":
			count := len(event.Payload.Commits)
			fmt.Printf("Pushed %d commit(s) to %s\n", count, event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("Interacted with an issue in %s\n", event.Repo.Name)
		case "IssueCommentEvent":
			fmt.Printf("Commented on an issue in %s\n", event.Repo.Name)
		case "PullRequestEvent":
			fmt.Printf("Interacted with a pull request in %s\n", event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("Starred %s\n", event.Repo.Name)
		default:
			fmt.Printf("Did %s in %s\n", event.Type, event.Repo.Name)
		}
	}
}

func main() {
	var eventType string
	flag.StringVar(&eventType, "type", "", "Filter events by type (e.g., PushEvent)")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: github-activity <username> [--type=EventType]")
		return
	}

	username := args[0]

	events, err := fetchEvents(username)
	if err != nil {
		fmt.Println("Error fetching events:", err)
		return
	}

	if len(events) == 0 {
		fmt.Println("No recent activity found.")
		return
	}

	displayEvents(events, eventType)
}
