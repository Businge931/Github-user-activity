package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func FetchUserActivity(username string) ([]UserActivity, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var apiErr ErrMessage
		err := json.Unmarshal(body, &apiErr)
		if err != nil {
			apiErr.Message = "Not Found"
			apiErr.DocumentationURL = "https://docs.github.com/rest/activity/events#list-events-for-the-authenticated-user"
		}
		return nil, ErrMessage{
			Message:          apiErr.Message,
			DocumentationURL: apiErr.DocumentationURL,
			Status:           fmt.Sprintf("%d", resp.StatusCode),
		}
	}

	var userActivity []UserActivity
	if err := json.Unmarshal(body, &userActivity); err != nil {
		return nil, err
	}

	return userActivity, nil
}

func FilterUserActivity(userActivity []UserActivity, eventType EventType) []UserActivity {
	var filteredActivity []UserActivity
	for _, activity := range userActivity {
		if strings.EqualFold(string(activity.Type), string(eventType)) {
			filteredActivity = append(filteredActivity, activity)
		}
	}
	return filteredActivity
}
