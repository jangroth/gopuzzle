package main

import (
	"testing"
)

func TestNextFreeCell(t *testing.T) {
	testPuzzle := NewPuzzle(5, 3)
	testPuzzle.dump()

	pnt := testPuzzle.nextFreeCell(Point{1, 1})
	if !(pnt.x == 2 && pnt.y == 1) {
		t.Errorf("Failed - (1, 1) returned %s instead of (2,1)", pnt)
	}

	pnt = testPuzzle.nextFreeCell(Point{0, 0})
	if !(pnt.x == 1 && pnt.y == 1) {
		t.Errorf("Failed - (0, 0) returned %s instead of (1,1)", pnt)
	}

	pnt = testPuzzle.nextFreeCell(Point{2, 1})
	if !(pnt.x == 3 && pnt.y == 1) {
		t.Errorf("Failed - (2, 1) returned %s instead of (3,1)", pnt)
	}

	pnt = testPuzzle.nextFreeCell(Point{4, 1})
	if !(pnt.x == -1 && pnt.y == -1) {
		t.Errorf("Failed - (4, 1) returned %s instead of (-1,-1)", pnt)
	}

	pnt = testPuzzle.nextFreeCell(Point{2, 0})
	if !(pnt.x == 1 && pnt.y == 2) {
		t.Errorf("Failed - (2, 0) returned %s instead of (1,1)", pnt)
	}
}
