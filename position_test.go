package chessgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStartingPosition(t *testing.T) {
	want := testingPositions["starting"]
	got := NewStartingPosition()
	comparePositions(t, want, got)
}

func TestPosition_PieceAt(t *testing.T) {
	startingposition := NewStartingPosition()
	assert.Equal(t, startingposition.PieceAt(0), BlackRook)
	assert.Equal(t, startingposition.PieceAt(63), WhiteRook)
	assert.Equal(t, startingposition.PieceAt(32), Empty)
}
