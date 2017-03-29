package main

import (
	"testing"
)

func TestTrivialSolve(t *testing.T) {
	piece := NewPiece(2, Point{0, 0}, Point{1, 0})
	piece2 := NewPiece(2, Point{3, 0}, Point{1, 0})
	puzzle := NewPuzzle(4, 3, *piece, *piece2)
	puzzle.dump()
	t.Error("yes")
	//	success := puzzle.Solve()
	//	if !success {
	//		t.Error("this puzzle is solvable")
	//	}
}
