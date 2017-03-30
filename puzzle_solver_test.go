package main

import (
	"testing"
)

func TestTrivialSolve(t *testing.T) {
	piece := NewPiece(2, Point{0, 0}, Point{1, 0})
	puzzle := NewPuzzle(4, 3, *piece)
	ok := Solve(puzzle, Point{0, 0})
	if !ok {
		t.Error("This should have been solved")
	}
}
