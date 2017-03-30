package main

import (
	"testing"
)

func TestTrivialSolve(t *testing.T) {
	p1 := NewPiece(2, Point{0, 0}, Point{1, 0})
	puzzle := NewPuzzle(4, 3, *p1)
	ok := Solve(puzzle, Point{0, 0})
	if !ok {
		t.Error("This should have been solved")
	}
}

func TestTwoPieceSolve(t *testing.T) {
	p1 := NewPiece(2, Point{0, 0}, Point{1, 0}, Point{1, 1})
	p2 := NewPiece(3, Point{0, 0}, Point{1, 0}, Point{1, 1})
	puzzle := NewPuzzle(5, 4, *p1, *p2)
	ok := Solve(puzzle, Point{0, 0})
	t.Error("output")
	if !ok {
		t.Error("This should have been solved")
	}
}
