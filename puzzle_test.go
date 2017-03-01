package main

import (
	"fmt"
	"testing"
)

func TestPlacement(t *testing.T) {
	puzzle := NewPuzzle(5, 3)
	matrix := &puzzle.matrix
	piece := NewPiece(2, Point{1, 0}, Point{1, 1})
	var success bool
	matrix, success = place(&puzzle.matrix, piece, Point{1, 1})
	fmt.Printf("%t", matrix)
	if !success {
		t.Error("This piece should fit into the matrix")
	}
}

func TestNewPiece(t *testing.T) {
	piece := NewPiece(2, Point{1, 0}, Point{1, 1})
	piece.dump()
	if !(piece.piectrix[1][0] == 2 && piece.piectrix[0][1] == 0 && piece.piectrix[1][1] == 2 && piece.piectrix[0][0] == 0) {
		t.Errorf("Testpiece doesn't look right %s", piece)
	}
}

func TestNextFreeCell(t *testing.T) {
	puzzle := NewPuzzle(5, 3)
	puzzle.dump()

	pnt := puzzle.nextFreeCell(Point{1, 1})
	if !(pnt.x == 2 && pnt.y == 1) {
		t.Errorf("Failed - (1, 1) returned %s instead of (2,1)", pnt)
	}

	pnt = puzzle.nextFreeCell(Point{0, 0})
	if !(pnt.x == 1 && pnt.y == 1) {
		t.Errorf("Failed - (0, 0) returned %s instead of (1,1)", pnt)
	}

	pnt = puzzle.nextFreeCell(Point{2, 1})
	if !(pnt.x == 3 && pnt.y == 1) {
		t.Errorf("Failed - (2, 1) returned %s instead of (3,1)", pnt)
	}

	pnt = puzzle.nextFreeCell(Point{4, 1})
	if !(pnt.x == -1 && pnt.y == -1) {
		t.Errorf("Failed - (4, 1) returned %s instead of (-1,-1)", pnt)
	}

	pnt = puzzle.nextFreeCell(Point{2, 0})
	if !(pnt.x == 1 && pnt.y == 1) {
		t.Errorf("Failed - (2, 0) returned %s instead of (1,1)", pnt)
	}
}
