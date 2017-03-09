package main

import (
	"fmt"
	"testing"
)

func TestPlacement(t *testing.T) {
	var matrix *Matrix
	puzzle := NewPuzzle(5, 3)
	piece := NewPiece(2, Point{1, 0}, Point{1, 1})
	_, success := place(&puzzle.matrix, piece, Point{1, 1})
	if success {
		t.Error("This piece should not fit into the matrix")
	}
	puzzle = NewPuzzle(5, 4)
	matrix, success = place(&puzzle.matrix, piece, Point{1, 1})
	fmt.Println(matrix)
	matrix.dump()
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
}

func TestNewPiece(t *testing.T) {
	piece := NewPiece(2, Point{1, 0}, Point{1, 1})
	piece.piecetrix.dump()
	if !(len(piece.piecetrix) == 2 && len(piece.piecetrix[0]) == 2) {
		t.Error("Testpiece doesn't have the right size")
	}
	if !(piece.piecetrix[1][0] == 2 && piece.piecetrix[0][1] == 0 && piece.piecetrix[1][1] == 2 && piece.piecetrix[0][0] == 0) {
		t.Errorf("Testpiece doesn't look right %s", piece)
	}
}

func TestNextFreeCell(t *testing.T) {
	puzzle := NewPuzzle(5, 3)
	puzzle.matrix.dump()

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
			fmt.Printf("%d:m%d-r%d \t", x, (*matrix)[x][index], row[x])
			if (*matrix)[x][index] != row[x] {
				return false
			}
		}
		fmt.Println()
	}
	return true
}
