package data

// =============================================================================
// PORTFOLIO CONTENT — edit this file to customise your Argus portfolio terminal.
// The engine (store.go, ui/) never needs to be touched for content changes.
// =============================================================================

import "time"

// OwnerHandle is displayed in the stats panel (e.g. github username).
const OwnerHandle = "NyxUltor"

// ProjectStart is the anchor date used to calculate the live project uptime counter.
var ProjectStart = time.Date(2026, time.June, 19, 12, 30, 0, 0, time.Local)



// AboutText is rendered by the `about` command.
const AboutText = `--- ABOUT ---
Independent systems designer. I conceptualize, architect, and direct
software from specification to stable release — handling system design,
logic-flow, UX decisions, and quality control across the full build cycle.

Four months in. Two shipped projects. Several in the pipeline.
Not slowing down.`

// ContactText is rendered by the `contact` command.
const ContactText = `--- CONTACT ---
Email:    ultor.nyx@gmail.com
GitHub:   github.com/NyxUltor
LinkedIn: linkedin.com/in/krishna-bisht-985062373
Discord:  nyx_ultor`

// Skills is rendered by the `skills` command.
var Skills = []string{
	"Design:    System architecture, logic-flow specification, UX decisions",
	"Workflow:  Generative pipelines, multi-AI coordination, build direction",
	"Environ:   Arch Linux (primary), KDE, multi-boot, shell scripting, Git",
	"Tools:     Bubble Tea, Lipgloss, Python scripting, Kdenlive",
	"Other:     Local-first data models, file I/O, input validation patterns",
}

// StatsText is rendered by the `stats` command.
const StatsText = `--- OPERATOR STATS ---
Lifting:
  Deadlift     125kg  (2.27x BW)
  Bench        60kg   (1.09x BW)
  Bodyweight   55kg

Academic:
  CBSE Class 10          95%        2025
  SASMO National Rank 1  Bronze     2025

Languages:
  Proficient   English, Hindi, Kumaoni
  Learning     Russian, Japanese

Art:
  40+ works completed
  10+ brought to portfolio-grade

Misc:
  Chess, philosophy, reading, gaming`

// Projects is rendered by the `projects` command.
// Key: slug used for sub-command lookup (e.g. `projects hercules`).
var Projects = map[string]Project{
	"heracles": {
		Name:        "Heracles",
		Description: "Local-first Android fitness logger. Dual-mode input, collision-safe file indexing, progressive disclosure from casual presets to raw JSON config. Shipped.",
		Stack:       []string{"Android", "Generative Workflow"},
	},
	"argus": {
		Name:        "Argus",
		Description: "This terminal. A Bubble Tea portfolio shell — local execution, no tracking, full control. Named after the hundred-eyed giant.",
		Stack:       []string{"Go", "Bubble Tea", "Lipgloss"},
	},
	"contemplation": {
		Name:        "Contemplation",
		Description: "Web portfolio. Animations, shaders, textures. The visual counterpart to this terminal.",
		Stack:       []string{"Web", "GLSL"},
	},
	"selena": {
		Name:        "Selena",
		Description: "Local AI assistant — closer to JARVIS than Google Assistant. Full command access, high customizability, no cloud dependency.",
		Stack:       []string{"In Development"},
	},
}
