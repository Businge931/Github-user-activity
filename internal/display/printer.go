package display

import (
	"encoding/json"
	"fmt"
)

// PrintUserActivity pretty-prints the user activity (raw or filtered)
func PrintUserActivity(activity any) {
	activityJSON, err := json.MarshalIndent(activity, "", "  ")
	if err != nil {
		fmt.Println("Error formatting user activity:", err)
		return
	}
	fmt.Println(string(activityJSON))
}

// PrintError pretty-prints errors, including custom API errors
func PrintError(err error) {
	activityJSON, jerr := json.MarshalIndent(err, "", "  ")
	if jerr != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(activityJSON))
}

// PrintEventTypes prints the available event types
func PrintEventTypes(eventTypes []string) {
	fmt.Println("Available event types:")
	for _, et := range eventTypes {
		fmt.Println("-", et)
	}
}
