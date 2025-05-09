# GitHub User Activity CLI

[Project on roadmap.sh](https://roadmap.sh/projects/github-user-activity)

## Overview

The **GitHub User Activity CLI** is a command-line tool that allows you to fetch and view recent public activity for any GitHub user. It provides a user-friendly interactive experience, supports filtering by event type, and handles common errors gracefully. The tool is designed for developers, recruiters, and anyone interested in tracking GitHub user activity from the terminal.

## Features

- Fetches recent public events for any GitHub username using the GitHub API
- Filter activity by event type (e.g., PushEvent, PullRequestEvent, IssueEvent, etc.)
- User-friendly prompts and error messages
- Input validation and the ability to exit at any prompt
- In-memory caching to avoid redundant API requests for the same user
- Structured and readable terminal output

## Usage

### 1. Build the CLI

```sh
cd /path/to/Github\ User\ Activity
# Ensure Go is installed (1.18+ recommended)
go build -o github-activity ./cmd
```

### 2. Run the CLI

```sh
./github-activity
```

You will be prompted to enter:
- A GitHub username (type `exit` to quit at any prompt)
- An event type to filter by (choose from the displayed list)

The tool will fetch and display the user's recent activity, filtered by your selection.

### 3. Example Session

```
Enter GitHub username: octocat
Available event types:
- PushEvent
- PullRequestEvent
- ...
Enter event type: PushEvent

Recent activity for octocat (PushEvent):
- [repo1] ...
- [repo2] ...
```

### 4. Error Handling
- If the user is not found, you will see a clear error message.
- If you exceed the GitHub API rate limit, you will be notified and given a link to documentation.

## Development & Testing

- The codebase is organized with clear separation between fetching, display, and input logic.
- Unit tests for the GitHub API logic use mocked HTTP servers (no real API calls in tests).
- To run tests:
  ```sh
  go test ./internal/github
  ```

## Project URL

For more details, see the project page: [https://roadmap.sh/projects/github-user-activity](https://roadmap.sh/projects/github-user-activity)

## License

MIT License (or specify your license here)
