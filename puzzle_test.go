package main

import (
	"testing"
)

func TestNextFreeCell(t *testing.T) {
	testPuzzle := NewPuzzle(3, 5)
	testPuzzle.dump()

	x, y := testPuzzle.nextFreeCell(1, 1)
	if !(x == 1 && y == 2) {
		t.Errorf("Failed, returned (%d,%d) instead of (1,2)", x, y)
	}

	x, y = testPuzzle.nextFreeCell(1, 1)
	if !(x == 1 && y == 2) {
		t.Errorf("Failed, returned (%d,%d) instead of (1,2)", x, y)
	}

	x, y = testPuzzle.nextFreeCell(0, 0)
	if !(x == 1 && y == 1) {
		t.Errorf("Failed, returned (%d,%d) instead of (1,1)", x, y)
	}

	x, y = testPuzzle.nextFreeCell(2, 2)
	if !(x == -1 && y == -1) {
		t.Errorf("Failed, returned (%d,%d) instead of (-1,-1)", x, y)
	}

	// 	x, y = testPuzzle.nextFreeCell(2, 0)
	// 	if !(x == -1 && y == -1) {
	// 		t.Errorf("Failed, returned (%d,%d) instead of (-1,-1)", x, y)
	// 	}
}
