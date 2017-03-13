package main

import (
	"testing"
)

func TestPermutate(t *testing.T) {
	piece := NewPiece(2, Point{0, 0}, Point{1, 0}, Point{2, 0}, Point{2, 1})
	permutatedPieces := piece.permutate()
	if !(len(permutatedPieces) == 8) {
		t.Error("Expected 8 different pieces")
	}
	piece = NewPiece(2, Point{0, 0}, Point{1, 0}, Point{1, 1})
	permutatedPieces = piece.permutate()
	if !(len(permutatedPieces) == 4) {
		t.Error("Expected 4 different pieces")
	}
}

func TestRotate(t *testing.T) {
	piece := NewPiece(2, Point{0, 0}, Point{1, 0}, Point{2, 0}, Point{2, 1})
	rotatedPiece := piece.rotate()
	if piece == rotatedPiece {
		t.Error("These should be two different objects")
	}
	if !compare(&rotatedPiece.matrix,
		[]int{0, 2},
		[]int{0, 2},
		[]int{2, 2}) {
		t.Error("This matrix doesn't look right")
	}
	rotatedPiece2 := rotatedPiece.rotate()
	if !compare(&rotatedPiece2.matrix,
		[]int{2, 0, 0},
		[]int{2, 2, 2}) {
		t.Error("This matrix doesn't look right")
	}
}

func TestMirror(t *testing.T) {
	piece := NewPiece(2, Point{0, 0}, Point{1, 0}, Point{2, 0}, Point{2, 1})
	mirroredPiece := piece.mirror()
	if piece == mirroredPiece {
		t.Error("These should be two different objects")
	}
	if !compare(&mirroredPiece.matrix,
		[]int{0, 0, 2},
		[]int{2, 2, 2}) {
		t.Error("This matrix doesn't look right")
	}
	piece = NewPiece(3, Point{0, 0}, Point{1, 0}, Point{2, 0}, Point{2, 1}, Point{3, 3})
	mirroredPiece = piece.mirror()
	if !compare(&mirroredPiece.matrix,
		[]int{0, 0, 0, 3},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 3, 0},
		[]int{3, 3, 3, 0}) {
		t.Error("This matrix doesn't look right")
	}
}

func TestPlacement(t *testing.T) {
	var matrix *Matrix
	puzzle := NewPuzzle(5, 3)
	piece := NewPiece(2, Point{1, 0}, Point{1, 1})
	matrix, success := place(&puzzle.matrix, piece, Point{1, 1})
	if success {
		t.Error("This piece should not fit into the matrix")
	}
	if !compare(matrix,
		[]int{1, 1, 1, 1, 1},
		[]int{1, 0, 0, 0, 1},
		[]int{1, 1, 1, 1, 1}) {
		t.Error("This matrix doesn't look right")
	}
	puzzle = NewPuzzle(5, 4)
	matrix, success = place(&puzzle.matrix, piece, Point{1, 1})
	if !success {
		t.Error("This piece should fit into the matrix")
	}
	if !compare(matrix,
		[]int{1, 1, 1, 1, 1},
		[]int{1, 0, 2, 0, 1},
		[]int{1, 0, 2, 0, 1},
		[]int{1, 1, 1, 1, 1}) {
		t.Error("This matrix doesn't look right")
	}
	pieceX1 := NewPiece(2, Point{0, 0}, Point{1, 1})
	pieceX2 := NewPiece(3, Point{0, 1}, Point{1, 0})
	puzzle = NewPuzzle(4, 4)
	matrix, _ = place(&puzzle.matrix, pieceX1, Point{1, 1})
	matrix, _ = place(&puzzle.matrix, pieceX2, Point{1, 1})
	if !compare(matrix,
		[]int{1, 1, 1, 1},
		[]int{1, 2, 3, 1},
		[]int{1, 3, 2, 1},
		[]int{1, 1, 1, 1}) {
		t.Error("This matrix doesn't look right")
	}

}

func TestNewPiece(t *testing.T) {
	piece := NewPiece(2, Point{1, 0}, Point{1, 1})
	if !(len(piece.matrix) == 2 && len(piece.matrix[0]) == 2) {
		t.Error("Testpiece doesn't have the right size")
	}
	if !(piece.matrix[1][0] == 2 && piece.matrix[0][1] == 0 && piece.matrix[1][1] == 2 && piece.matrix[0][0] == 0) {
		t.Errorf("Testpiece doesn't look right %s", piece)
	}
}

func TestNextFreeCell(t *testing.T) {
	puzzle := NewPuzzle(5, 3)

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

func compare(matrix *Matrix, rows ...[]int) bool {
	for index, row := range rows {
		for x := 0; x < len(row); x++ {
			if (*matrix)[x][index] != row[x] {
				return false
			}
		}
	}
	return true
}
