package ui

import (
	"math"
	"testing"
	"time"
)

func TestLogoHitmapExclusions(t *testing.T) {
	hitmap := getLogoHitmap()
	if len(hitmap) != len(EyeArt) {
		t.Fatalf("Expected hitmap height %d, got %d", len(EyeArt), len(hitmap))
	}

	// Verify that space and U+2800 are not marked as hit, but other braille runes are.
	for r, line := range EyeArt {
		runes := []rune(line)
		for c, ch := range runes {
			isHit := hitmap[r][c]
			if ch == ' ' || ch == '\u2800' {
				if isHit {
					t.Errorf("At row %d, col %d: character %q should not be marked as solid in hitmap", r, c, ch)
				}
			} else {
				if !isHit {
					t.Errorf("At row %d, col %d: character %q should be marked as solid in hitmap", r, c, ch)
				}
			}
		}
	}
}

func TestClampVelocity(t *testing.T) {
	// Test normal magnitude
	vx, vy := clampVelocity(0.5, 0.5, 0.2, 1.2)
	mag := math.Sqrt(vx*vx + vy*vy)
	if mag < 0.2 || mag > 1.2 {
		t.Errorf("Expected mag in [0.2, 1.2], got %f (vx=%f, vy=%f)", mag, vx, vy)
	}

	// Test zero vector
	vx, vy = clampVelocity(0, 0, 0.2, 1.2)
	mag = math.Sqrt(vx*vx + vy*vy)
	if math.Abs(mag-0.2) > 1e-9 {
		t.Errorf("Expected mag to be clamped to min 0.2, got %f", mag)
	}

	// Test too small vector
	vx, vy = clampVelocity(0.05, 0.05, 0.2, 1.2)
	mag = math.Sqrt(vx*vx + vy*vy)
	if math.Abs(mag-0.2) > 1e-9 {
		t.Errorf("Expected mag to be clamped to min 0.2, got %f", mag)
	}

	// Test too large vector
	vx, vy = clampVelocity(2.0, 2.0, 0.2, 1.2)
	mag = math.Sqrt(vx*vx + vy*vy)
	if math.Abs(mag-1.2) > 1e-9 {
		t.Errorf("Expected mag to be clamped to max 1.2, got %f", mag)
	}
}

func TestUpdateFloatingEyeCollisionAndDazed(t *testing.T) {
	// Let's create a model with terminal size enough to contain eye
	m := Model{
		TerminalWidth:        80,
		TerminalHeight:       24,
		EyeCooldownRemaining: 0,
	}

	// Manually set FloatingEyeState to trigger collision
	// Let's put the eye sprite right on the logo: logo starts at row 0, col 2.
	// The EyeArt has solid characters in its first few lines/columns.
	// Let's place it at X: 2, Y: 0. This should overlap the logo.
	m.FloatingEye = &FloatingEyeState{
		X:   2,
		Y:   0,
		VX:  0.5,
		VY:  0.5,
		TTL: 10 * time.Second,
	}

	m = m.updateFloatingEye()

	// Since X:2, Y:0 overlaps the solid logo parts, the eye should now be dazed
	if m.FloatingEye == nil {
		t.Fatal("Expected floating eye to not despawn from collision")
	}
	if m.FloatingEye.DazedRemaining <= 0 {
		t.Errorf("Expected eye to enter DazedState (DazedRemaining > 0), got %v", m.FloatingEye.DazedRemaining)
	}
}
