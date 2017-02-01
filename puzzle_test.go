package main

import (
	"testing"
)

func TestNextFreeCell(t *testing.T) {
	testPuzzle := NewPuzzle(5, 3)
	testPuzzle.dump()

	x, y := testPuzzle.nextFreeCell(1, 1)
	if !(x == 2 && y == 1) {
		t.Errorf("Failed, returned (%d,%d) instead of (2,1)", x, y)
	}

	x, y = testPuzzle.nextFreeCell(0, 0)
	if !(x == 1 && y == 1) {
		t.Errorf("Failed, returned (%d,%d) instead of (1,1)", x, y)
	}

	x, y = testPuzzle.nextFreeCell(2, 1)
	if !(x == 3 && y == 1) {
		t.Errorf("Failed, returned (%d,%d) instead of (3,1)", x, y)
	}

	x, y = testPuzzle.nextFreeCell(4, 1)
	if !(x == -1 && y == -1) {
		t.Errorf("Failed, returned (%d,%d) instead of (-1,-1)", x, y)
	}

	x, y = testPuzzle.nextFreeCell(2, 0)
	if !(x == 1 && y == 1) {
		t.Errorf("Failed, returned (%d,%d) instead of (1,1)", x, y)
	}
}
