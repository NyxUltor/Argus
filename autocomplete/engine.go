package autocomplete

import (
	"argus/data"
	"strings"
)

func FilterCommands(input string) []string {
	if input == "" {
		return []string{}
	}

	store := data.GetPortfolioData()
	var matches []string

	for _, cmd := range store.Commands {
		if strings.HasPrefix(cmd, input) {
			matches = append(matches, cmd)
		}
	}

	if strings.HasPrefix(input, "projects ") {
		subInput := strings.TrimPrefix(input, "projects ")
		for projKey := range store.Projects {
			if subInput == "" || strings.HasPrefix(projKey, subInput) {
				matches = append(matches, "projects "+projKey)
			}
		}
	}

	return matches
}
