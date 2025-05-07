package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dovjay/roadmap.sh-github-user-activity/cache"
	"github.com/dovjay/roadmap.sh-github-user-activity/dto"
)

// fetchEvents fetches user activity from GitHub or from cache
func FetchEvents(username string) ([]dto.GitHubEvent, error) {
	if events, ok := cache.LoadFromCache(username); ok {
		fmt.Println("✅ Loaded from cache")
		return events, nil
	}

	url := fmt.Sprintf("https://api.github.com/users/%s/events?per_page=100", username)
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API error: %s\n%s", resp.Status, string(body))
	}

	var events []dto.GitHubEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}

	cache.SaveToCache(username, events)
	fmt.Println("📡 Fetched from GitHub and cached")
	return events, nil
}
