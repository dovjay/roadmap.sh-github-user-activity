package display

import (
	"fmt"
	"strings"

	"github.com/dovjay/roadmap.sh-github-user-activity/dto"
)

// displayEvents prints GitHub events filtered by type
func DisplayEvents(events []dto.GitHubEvent, filter string) {
	for _, event := range events {
		if filter != "" && !strings.EqualFold(event.Type, filter) {
			continue
		}

		switch event.Type {
		case "PushEvent":
			fmt.Printf("Pushed to %s\n", event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("Interacted with issues in %s\n", event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("Starred %s\n", event.Repo.Name)
		case "CreateEvent":
			fmt.Printf("Created something in %s\n", event.Repo.Name)
		default:
			fmt.Printf("%s in %s\n", event.Type, event.Repo.Name)
		}
	}
}
