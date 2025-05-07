# GitHub Activity CLI

A simple Go command-line tool to fetch and display recent GitHub user activity using the GitHub Events API.

Project from: [roadmap.sh GitHub User Activity](https://roadmap.sh/projects/github-user-activity)

---

## ✨ Features

- 🔍 Fetches recent activity of any public GitHub user
- 📁 Caches API responses for 5 minutes to reduce API usage
- 🎯 Supports filtering by event type (e.g. PushEvent, WatchEvent)
- ⚡ Fast and terminal-friendly

---

## 📦 Prerequisites

- [Go (>=1.18)](https://golang.org/dl/)
- Internet connection

---

## 🚀 Installation

```bash
git clone https://github.com/yourusername/github-activity.git
cd github-activity
go build -o github-activity
```

---

## 🛠️ Usage

```bash
./github-activity <github_username> [--type=EventType]
```

---

## 📌 Examples

```bash
./github-activity vercel
./github-activity vercel --type=PushEvent
./github-activity torvalds --type=WatchEvent
```

---

## 🗃️ Caching

- Cached responses are stored in the ./cache directory
- Automatically reused for 5 minutes to improve performance
- If cache is stale or missing, the app fetches from the GitHub API
