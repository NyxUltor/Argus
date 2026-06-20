package data

import "time"

type Project struct {
	Name        string
	Description string
	Stack       []string
}

type PortfolioStore struct {
	AboutText    string
	Skills       []string
	Projects     map[string]Project
	Commands     []string
	ProjectStart time.Time
	ContactText  string
}

func GetPortfolioData() PortfolioStore {
	// Start time anchor: June 19, 2026, 12:30 PM (Local Time)
	startTime := time.Date(2026, time.June, 19, 12, 30, 0, 0, time.Local)

	return PortfolioStore{
		Commands:     []string{"about", "projects", "skills", "contact", "clear", "exit", "quit"},
		ProjectStart: startTime,

		AboutText: `--- ABOUT ME ---
Self-taught developer operating under strict engineering discipline.
Focusing on low-overhead system execution, custom developer tooling, 
and highly reliable localized processing architectures.`,

		ContactText: `--- CONTACT INFORMATION ---
Email:    veryl@example.com
GitHub:   github.com/veryl
Terminal: tty.argus.sh`,

		Skills: []string{
			"Languages:  Go, Python, Shell",
			"Platforms:  Arch Linux (i3/KDE), Docker",
			"Frameworks: Bubble Tea, Prompt Toolkit",
		},

		Projects: map[string]Project{
			"hercules": {
				Name:        "Hercules",
				Description: "A tailored, high-discipline workout logging application written to optimize personal physical metrics without tracking bloat.",
				Stack:       []string{"Android", "Custom Core Arrays"},
			},
			"atlantis": {
				Name:        "Atlantis / Poseidon",
				Description: "Conceptualized industrial AI verification framework engineered for edge-compute deployment within high-stakes offshore environments.",
				Stack:       []string{"Go", "Local Processing Engine"},
			},
		},
	}
}
