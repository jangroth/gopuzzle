package main

import (
	"testing"
)

func TestTrivialSolve(t *testing.T) {
	piece := NewPiece(2, Point{0, 0}, Point{1, 0})
	piece2 := NewPiece(3, Point{3, 0}, Point{1, 0})
	piece3 := NewPiece(4, Point{1, 0}, Point{1, 1}, Point{1, 2}, Point{0, 2})
	puzzle := NewPuzzle(4, 3, *piece, *piece2, *piece3)
	puzzle.dump()
	t.Error("yes")
}
