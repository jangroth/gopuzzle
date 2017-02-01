package main

import (
	"testing"
)

func TestNextFreeCell(t *testing.T) {
	testPuzzle := initPuzzle(3, 5)
	testPuzzle.dump()

	x, y := nextFreeCell(testPuzzle.Matrix, 1, 1)
	if !(x == 1 && y == 2) {
		t.Errorf("Failed, returned (%d,%d) instead of (1,2)", x, y)
	}

	x, y = nextFreeCell(testPuzzle.Matrix, 1, 1)
	if !(x == 1 && y == 2) {
		t.Errorf("Failed, returned (%d,%d) instead of (1,2)", x, y)
	}

	x, y = nextFreeCell(testPuzzle.Matrix, 0, 0)
	if !(x == 1 && y == 1) {
		t.Errorf("Failed, returned (%d,%d) instead of (1,1)", x, y)
	}

	//	x, y = nextFreeCell(testPuzzle.Matrix, 2, 0)
	//	if !(x == -1 && y == -1) {
	//		t.Errorf("Failed, returned (%d,%d) instead of (-1,-1)", x, y)
	//	}
}
