package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dovjay/roadmap.sh-github-user-activity/dto"
)

// loadFromCache reads cache file if it's fresh (<5 min)
func LoadFromCache(username string) ([]dto.GitHubEvent, bool) {
	path := getCachePath(username)
	info, err := os.Stat(path)
	if err != nil || time.Since(info.ModTime()) > 5*time.Minute {
		return nil, false
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, false
	}

	var events []dto.GitHubEvent
	if err := json.Unmarshal(data, &events); err != nil {
		return nil, false
	}

	return events, true
}

// saveToCache saves events to a cache file
func SaveToCache(username string, events []dto.GitHubEvent) {
	os.MkdirAll("./cache", os.ModePerm)
	path := getCachePath(username)

	data, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		fmt.Println("Error encoding cache:", err)
		return
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		fmt.Println("Error writing cache:", err)
	}
}

// getCachePath returns the path to the user's cache file
func getCachePath(username string) string {
	return filepath.Join("cache", fmt.Sprintf("cache_%s.json", username))
}
