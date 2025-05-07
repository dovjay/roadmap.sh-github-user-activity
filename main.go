package main

import (
	"flag"
	"fmt"

	"github.com/dovjay/roadmap.sh-github-user-activity/display"
	"github.com/dovjay/roadmap.sh-github-user-activity/fetcher"
)

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
	events, err := fetcher.FetchEvents(username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(events) == 0 {
		fmt.Println("No recent activity found.")
		return
	}

	display.DisplayEvents(events, eventType)
}
