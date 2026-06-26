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

// LogoLines is the ASCII/braille art displayed in the header panel.
// Each element is one terminal row. Keep all lines roughly the same display width
// so the stats column aligns neatly.
var LogoLines = []string{
	"⠀⠀⠀⠀ ⢀⠀⠀⠀⠀⠀⠀⠀⠄⠀⠀⠀⠤⢀⣁⣀⠠⣤⠌⢻⣛⣷⣖⣿⣿⣿⣗⣷⠖⣿⣿⣿⣿⣻⣿⠯⣡⠔⢁⠔⣠⠞⣡⣔⠼⣶⠟⠮⣃⠐⠉⠁⠀⠀⠀⠉⠁⠀⠀⠠⡠⠤⠲⠊⡀⠄⠁⠀⠀  ⡄⣲",
	"⠀⠀⠀ ⡠⠁⠀⠀⠀⠀⠀⠈⠑⠒⠈⢩⣶⣶⣶⣾⢲⣛⣿⢿⠟⣿⣿⣿⣿⡥⣶⡷⣟⣿⣿⣿⣿⡥⣚⣥⡾⣛⣛⣴⡿⠎⣫⠟⠉⠐⠂⢸⠐⠀⠀⠀⠀⠀⠀⠉⠁⠒⠑⠒⠀⠈⠀⠀  ⠀⢀⣴⣬⣾⣿⣋ ",
	"⠀⢀⠔⠁⠀⠀⠀⢀⠀⠀⢀⣐⣛⢻⣿⠿⠿⠿⣽⡷⣯⣿⣶⣞⣭⣯⠷⣻⣽⣷⣿⣿⣿⡿⢛⠁⢅⣋⣝⣶⠶⠏⠡⠄⠒⠀⢈⠄⠀⠀⡠⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀  ⠀⠀⠀⣀⢐⣴⣯⣿⣿⠟⣩⡿   ",
	"⠀⠀⡤⠀⠀⢪⡠⠤⠬⢭⣭⢿⣟⣷⣽⣛⣛⣭⡿⣟⡿⢿⢿⣿⡿⣾⠻⠍⢓⠩⠕⠏⠉⠋⠐⠀⠂⠀⠀⠀⠀⠠⠈⠀⠀⣴⠀⠀⢀⠀⠀⣀⢀⣀⣠⣠⣤⣤⣴⣶⣶⣶⣶⣿⣿⣿⣿⣿⣿⣿⣟⣠⣾⠓⠀   ",
	"⠀⣁⡤⠟⠙⢁⢔⠾⢫⢴⣯⣟⣿⠞⡝⠉⠉⠋⠃⠊⠁⠀⠀⣀⣅⣀⣀⣀⣀⣀⢠⣶⣠⣤⢤⣴⠶⣶⠾⠥⣴⡶⣶⣾⣭⣟⠺⡿⢟⠻⠟⠛⠻⣿⢿⣿⣟⣛⣭⣴⣾⣾⣿⣿⠿⠟⠋⠉⠁⠀       ⠀ ",
	"⠜⡁⠤⢒⢀⣄⣴⠺⠛⠛⠑⣉⢠⡴⣶⣒⣢⣦⣦⣴⣾⣷⣿⣯⣿⣿⣿⣿⣿⣿⣼⣥⣤⣾⣿⣶⣾⣿⣶⣾⣴⣿⣷⣾⣶⣿⣿⣾⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣋⠉⢀⡠⠆⣀⣀⣀⣀⣀⠠           ",
	"⡰⢖⣿⣻⠏⠛⠁⠀⠀⠠⢠⣶⣭⣾⢿⣛⣿⡯⣴⣷⣷⣿⣿⣿⣿⣿⣿⣿⣿⡟⢿⢿⣿⣿⣿⣿⣾⣿⣿⣿⡿⢿⣿⢾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣮⣙⣾⣿⡋⠉⠃⠉⠀⠀⠀⠀⠀           ",
	"⠋⠉⠀⠀⠀⠀⠀⢄⣠⣴⠿⢫⣿⡷⣿⡹⠑⣱⣿⣿⣿⣋⠝⢋⡛⢿⣿⣿⣽⣿⣿⣾⡿⠻⣿⡟⢩⢻⣉⣿⢋⢂⢀⠄⢈⠋⢟⠟⢿⣿⣿⣿⢿⣿⢿⡿⣿⠫⣈⠑⠊⠉⠉⠉⠉⠉⠉⠉⠉          ",
	"⡀⠀⠀⠀⠀⠀⣤⣾⢿⠿⣤⣾⡿⣱⣟⣷⣶⣿⡻⠹⣿⣏⠀⠨⠷⣼⣿⠿⣿⠿⠿⠿⠓⡖⢯⣣⡈⠨⣽⠏⠈⠈⠊⠀⠈⠂⠘⠆⣬⡿⠋⢩⡘⠟⠚⠓⠿⡦⣉⠂⡀⠀⠀⠀⠀⠀⠀⠀\t\t\t",
	"⠀⠀⠀⢀⡴⡾⡿⠃⣠⣼⣿⢿⡟⠛⠓⠒⠀⠀⠀⠀⠈⢷⣦⡄⠁⠈⠁⢀⠁⠀⠀⠐⠀⠐⢈⢊⢮⣿⠏⠀⠀⠀⠀⠀⠀⠀⢰⡽⠟⠀⣶⣿⡿⣇⠉⠛⢬⣧⠂⠑⠂⠁⠀⠀⠀⠀⠀⠀\t\t\t",
	"⠀⠀⢀⢽⡾⠋⢠⣾⢟⡾⠟⡟⡁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠻⢿⣷⣤⡤⢀⢀⠠⢀⣀⡀⢀⣠⠶⠋⠀⠀⠀⠀⠀⠀⣀⣤⠖⠋⠀⣠⣾⣯⣟⣛⣾⡆⠀⠀⠹⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀\t\t\t",
	"⠀⠀⡡⢉⣤⢞⡽⠒⠋⣴⠬⠼⠓⠒⠦⠤⠤⣀⣀⡀⠀⠀⠀⠀⠀⠛⠛⠛⠗⠿⠿⠛⠛⠋⠁⠀⠀⠀⠀⡀⠄⠐⠈⠀⢀⣠⣶⣿⣿⣷⡃⠈⢉⢼⠗⠀⠀⠰⠐⠀⠀⠀⠀⠀⠀⠀⠀⠀\t\t\t",
	"⠀⠀⢔⡿⢵⠟⠒⢉⣡⡴⠵⠒⠚⣲⠤⢖⡀⡀⠤⠈⠚⠛⠚⠶⠤⠤⠤⠄⠀⠀⠄⢀⠀⠀⠔⠒⠂⠉⠀⢀⣀⣤⣰⣵⣿⢿⣿⠿⣿⣿⣯⢊⠼⡋⣹⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\t\t\t",
	"⠀⠀⠡⠀⠈⠀⠐⠁⠀⠀⠀⠐⠾⣭⣶⠉⣼⡽⠛⡥⣢⢤⣤⣀⡀⠀⠀⠀⠀⠀⠀⡀⣀⢀⢀⣀⢠⣖⣽⡼⡧⣙⠿⢿⣷⡧⢉⣾⢠⡺⣿⠏⢈⠀⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\t\t\t",
	"⠀⠀⠁⡘⠀⠀⢄⠀⠠⢀⠀⢂⢼⠏⠀⣹⡟⠠⠀⣰⣘⣹⣡⡞⠛⣦⢱⢆⢑⢶⢤⡈⡖⣗⡷⣷⣥⡇⠈⣧⡇⢀⣠⣦⢿⣟⠋⠠⠳⠍⣼⠁⠀⠀⠀⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\t\t\t",
	"⠀⠀⠀⠀⠀⠀⠀⠀⠐⠀⡠⢚⠋⠀⢰⠋⡆⠀⠀⠓⠁⣿⡏⣀⠀⠘⣞⡾⠃⢸⢿⢾⣡⣿⠄⠐⢽⡗⡀⣻⣇⡧⠟⢛⢉⡗⠀⠀⠀⢀⠛⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\t\t\t",
}

// AboutText is rendered by the `about` command.
const AboutText = `--- ABOUT ME ---
Developer building efficient software and tailored developer tools.
Driven by an interest in low-overhead systems, Linux architectures, 
and practical engineering solutions. 

I prioritize reliability, simplicity, and measurable performance 
across every codebase.`

// ContactText is rendered by the `contact` command.
const ContactText = `--- CONTACT INFORMATION ---
Email:    ultor.nyx@gmail.com
GitHub:   https://github.com/NyxUltor
LinkedIn: https://www.linkedin.com/in/krishna-bisht-985062373/ 
Discord:  nyx_ultor `

// Skills is rendered by the `skills` command.this block really needs changes
var Skills = []string{
	"Languages:  Go, Python, Shell",
	"Platforms:  Arch Linux (i3/KDE), Docker",
	"Frameworks: Bubble Tea, Prompt Toolkit",
}

// Projects is rendered by the `projects` command.
// Key: slug used for sub-command lookup (e.g. `projects hercules`).
var Projects = map[string]Project{
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
	"argus": {
		Name:        "Argus",
		Description: "This terminal. A Bubble Tea portfolio shell built for sovereign self-presentation — no cloud, no tracking, full local execution.",
		Stack:       []string{"Go", "Bubble Tea", "Lipgloss"},
	},
}
