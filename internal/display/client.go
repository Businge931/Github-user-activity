package display

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username:", err)
		return ""
	}
	username = strings.TrimSpace(username)
	if strings.ToLower(username) == "exit" {
		fmt.Println("Goodbye!")
		os.Exit(0)
	}
	return username
}
