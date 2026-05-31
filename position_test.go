package chessgo

import (
	"testing"
)

func TestNewStartingPosition(t *testing.T) {
	want := testingPositions["starting"]
	got := NewStartingPosition()
	comparePositions(t, want, got)
}
