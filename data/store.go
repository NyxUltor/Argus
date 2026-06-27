package data

// =============================================================================
// ENGINE — do not edit for content changes. See portfolio.go instead.
// =============================================================================

import (
	"sync"
	"time"
)

// Project describes a single project entry shown by the `projects` command.
type Project struct {
	Name        string
	Description string
	Stack       []string
}

// PortfolioStore is the runtime-accessible snapshot of all portfolio data.
type PortfolioStore struct {
	AboutText    string
	ContactText  string
	Skills       []string
	Projects     map[string]Project
	Commands     []string
	ProjectStart time.Time
	OwnerHandle  string
	StatsText    string
}

var (
	storeOnce   sync.Once
	cachedStore PortfolioStore
)

// GetPortfolioData returns the singleton PortfolioStore.
// All content is sourced from portfolio.go — edit that file to customise.
func GetPortfolioData() PortfolioStore {
	storeOnce.Do(func() {
		cachedStore = PortfolioStore{
			Commands:     []string{"about", "projects", "skills", "contact", "stats", "fun", "clear", "exit", "quit"},
			AboutText:    AboutText,
			ContactText:  ContactText,
			Skills:       Skills,
			Projects:     Projects,
			ProjectStart: ProjectStart,
			OwnerHandle:  OwnerHandle,
			StatsText:    StatsText,
		}
	})
	return cachedStore
}
