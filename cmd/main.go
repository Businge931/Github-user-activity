package main

import (
	"fmt"
	"strings"

	display "github.com/Businge931/github-user-activity/internal/display"
	"github.com/Businge931/github-user-activity/internal/github"
)

func main() {

	eventTypes := make([]string, len(github.AllEventTypes))
	for i, et := range github.AllEventTypes {
		eventTypes[i] = string(et)
	}
	display.PrintEventTypes(eventTypes)

	var username string
	for {
		username = display.GetUserInput("Enter github username (or type 'exit'): ")
		if username == "" {
			fmt.Println("No username provided. Please try again.")
			continue
		}
		if strings.Contains(username, " ") {
			fmt.Println("Username cannot contain spaces. Please try again.")
			continue
		}
		break
	}

	fmt.Println("Fetching user activity for", username)

	userActivity, err := github.FetchUserActivity(username)
	if err != nil {
		if apiErr, ok := err.(github.ErrMessage); ok {
			switch apiErr.Status {
			case "404":
				fmt.Println("GitHub user not found. Did you enter the username correctly?")
			case "403":
				if strings.Contains(strings.ToLower(apiErr.Message), "rate limit") {
					fmt.Println("You have hit the GitHub API rate limit. Please try again later or authenticate with a token.")
				} else {
					fmt.Println("Access forbidden. You may not have permission to view this user's activity.")
				}
			default:
				display.PrintError(err)
			}
			return
		}
		display.PrintError(err)
		return
	}
	display.PrintUserActivity(userActivity)

	display.PrintEventTypes(eventTypes)
	var eventType string
	for {
		eventType = display.GetUserInput("Enter event type to filter by (or type 'exit'): ")
		if eventType == "" {
			fmt.Println("No event type provided. Please try again.")
			continue
		}
		valid := false
		for _, et := range github.AllEventTypes {
			if strings.EqualFold(string(et), eventType) {
				valid = true
				break
			}
		}
		if !valid {
			fmt.Println("Invalid event type. Please choose from the list.")
			continue
		}
		break
	}

	filteredActivity := github.FilterUserActivity(userActivity, github.EventType(eventType))
	display.PrintUserActivity(filteredActivity)

}
