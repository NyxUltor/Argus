package ui

// eye_animation.go — iris tracking for the header eye art.
// The iris position is driven by:
//   1. The floating easter egg eye position (when active) — priority
//   2. Mouse cursor position (fallback)
//   3. Center/no offset (fallback when no mouse input)
//
// TODO: implement iris offset rendering once eye art coordinates are mapped.
// For now, applyIrisOffset returns the art unchanged.

// IrisTarget holds the current tracking target in terminal cell coordinates.
type IrisTarget struct {
	X, Y   int
	Active bool
}

// applyIrisOffset takes the base eye art lines and an iris target,
// and returns the art lines with the iris shifted toward the target.
// Currently a no-op placeholder.
func applyIrisOffset(art []string, target IrisTarget, termW, termH int) []string {
	// TODO: implement braille dot manipulation or frame-swap approach
	return art
}
