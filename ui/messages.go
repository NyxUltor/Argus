package ui

// messages.go — all Bubble Tea message types for Argus.
// Add new async message types here as the command set grows.

// commandFinishedMsg is returned by async shell commands (exec_local.go)
// when execution completes. Index maps back to the History slice entry that
// showed "Running..." so Update() can overwrite it with the real output.
type commandFinishedMsg struct {
	Index  int
	Output string
}
