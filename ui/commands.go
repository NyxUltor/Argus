package ui

// commands.go — built-in command dispatch and output generation.
// Adding a new portfolio command: add a case here and a constant/var in data/portfolio.go.
// update.go never needs to be touched for new commands.

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"argus/data"
)

// rng is a package-level random source for flavour-text selection.
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// sudoQuotes are returned randomly when a user attempts a sudo command.
var sudoQuotes = []string{
	"Who decided that?",
	"That level of genjutsu doesn't work on me.",
	"Since when did you fall under the illusion that you could command me?",
}

// getCommandOutput returns the rendered output string for a recognised built-in
// command. Returns an empty string for unknown commands (caller handles fallback).
func getCommandOutput(cmd string) string {
	store := data.GetPortfolioData()
	switch cmd {
	case "about":
		return store.AboutText

	case "skills":
		var sb strings.Builder
		sb.WriteString("--- RECOGNIZED SYSTEM CAPABILITIES ---\n")
		for _, skill := range store.Skills {
			sb.WriteString(fmt.Sprintf("* %s\n", skill))
		}
		return sb.String()

	case "projects":
		var sb strings.Builder
		sb.WriteString("--- ACTIVE BUILD PIPELINES ---\n")
		var keys []string
		for k := range store.Projects {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			proj := store.Projects[k]
			sb.WriteString(fmt.Sprintf("\n> %s [%s]\n  %s\n", proj.Name, strings.Join(proj.Stack, ", "), proj.Description))
		}
		return sb.String()

	case "contact":
		return store.ContactText

	case "stats":
		return store.StatsText

	case "fun":
		return `--- FUN COMMANDS ---

Overlay  (type 'over' to dismiss, or just run another command):
  summon ball -n     spawn n bouncing balls, max 20, stacks across calls

Takeover  (press any key to return):
  summon matrix      matrix rain
  summon starfield   starfield warp

Easter egg:
  ...                you'll know it when you see it`

	case "help":
		return "Available system commands: about, projects, skills, stats, contact, fun, clear, exit, quit"

	default:
		return ""
	}
}

// isSudoAttempt returns true when the raw command text is a sudo invocation.
func isSudoAttempt(cmdClean, cmdText string) bool {
	return cmdClean == "sudo" || strings.HasPrefix(cmdText, "sudo ")
}

// randomSudoResponse picks a random sudo rejection quote.
func randomSudoResponse() string {
	return sudoQuotes[rng.Intn(len(sudoQuotes))]
}
