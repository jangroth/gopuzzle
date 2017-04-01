package main

import (
	"testing"
)

func TestTrivialSolve(t *testing.T) {
	p1 := NewPiece(2, Point{0, 0}, Point{1, 0})
	puzzle := NewPuzzle(4, 3, simpleBorder, *p1)
	ok := Solve(puzzle, Point{0, 0})
	if !ok {
		t.Error("This should have been solved")
	}
}

func TestTwoPiedeSolve(t *testing.T) {
	p1 := NewPiece(2, Point{0, 0}, Point{1, 0}, Point{1, 1})
	p2 := NewPiece(3, Point{0, 0}, Point{1, 0}, Point{1, 1})
	puzzle := NewPuzzle(5, 4, simpleBorder, *p1, *p2)
	ok := Solve(puzzle, Point{0, 0})
	if !ok {
		t.Error("This should have been solved")
	}
}

// func TistNiftyFifty(t *testing.T) {
// 	p1 := NewPiece(2, Point{0, 0}, Point{0, 1}, Point{0, 2}, Point{0, 3},
// 		Point{1, 3}, Point{2, 3}, Point{3, 3}, Point{4, 3}, Point{4, 4},
// 		Point{4, 5}, Point{4, 6}, Point{5, 6}, Point{6, 6}, Point{6, 7},
// 		Point{6, 8}, Point{7, 8}, Point{7, 9}, Point{7, 10}, Point{8, 10},
// 		Point{8, 11}, Point{9, 11}, Point{10, 11}, Point{11, 11}, Point{11, 12},
// 		Point{11, 13}, Point{12, 13}, Point{12, 14}, Point{12, 15})
// 	p1.matrix.dump()
// 	p2 := NewPiece(3, Point{0, 0}, Point{0, 1}, Point{0, 2}, Point{1, 2},
// 		Point{2, 2}, Point{2, 3}, Point{2, 4}, Point{2, 5}, Point{3, 5},
// 		Point{4, 5}, Point{4, 6}, Point{4, 7}, Point{4, 8}, Point{5, 8},
// 		Point{5, 9}, Point{5, 10}, Point{6, 10}, Point{7, 10}, Point{8, 10},
// 		Point{8, 11}, Point{9, 11}, Point{10, 11}, Point{11, 11}, Point{11, 12},
// 		Point{11, 13}, Point{12, 13}, Point{12, 14}, Point{12, 15}, Point{12, 16})
// 	p2.matrix.dump()
// 	p3 := NewPiece(4, Point{0, 0}, Point{0, 1}, Point{0, 2}, Point{0, 3},
// 		Point{1, 3}, Point{2, 3}, Point{2, 4}, Point{2, 5}, Point{2, 6},
// 		Point{3, 6}, Point{4, 6}, Point{5, 6}, Point{5, 7}, Point{6, 7},
// 		Point{7, 7}, Point{7, 8}, Point{7, 9}, Point{7, 10}, Point{7, 11},
// 		Point{8, 11}, Point{9, 11}, Point{9, 12}, Point{10, 12}, Point{11, 12},
// 		Point{12, 12}, Point{12, 13}, Point{12, 14}, Point{12, 15})
// 	p3.matrix.dump()
// 	p4 := NewPiece(5, Point{0, 0}, Point{0, 1}, Point{0, 2}, Point{0, 3},
// 		Point{0, 4}, Point{0, 5}, Point{1, 5}, Point{2, 5}, Point{3, 5},
// 		Point{3, 6}, Point{3, 7}, Point{4, 7}, Point{4, 8}, Point{4, 9},
// 		Point{5, 9}, Point{6, 9}, Point{7, 9}, Point{8, 9}, Point{8, 10}, Point{8, 11},
// 		Point{8, 12}, Point{9, 12}, Point{10, 12}, Point{11, 12}, Point{11, 13},
// 		Point{11, 14}, Point{11, 15}, Point{11, 16})
// 	p4.matrix.dump()
//
// 	t.Error("d")
//
// }
