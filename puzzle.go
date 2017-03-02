// puzzle solver
package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type Matrix [][]int

type Piece struct {
	piecetrix Matrix
}

type Puzzle struct {
	matrix Matrix
	pieces []Piece
}

func (p *Piece) dump() {
	dump(&p.piecetrix)
}

func NewPiece(value int, points ...Point) *Piece {
	maxX, maxY := 0, 0
	for _, val := range points {
		if val.x > maxX {
			maxX = val.x
		}
		if val.y > maxY {
			maxY = val.y
		}
	}
	matrix := createEmptyMatrix(maxX, maxY)
	for _, val := range points {
		matrix[val.x][val.y] = value
	}
	return &Piece{piecetrix: matrix}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p *Puzzle) dump() {
	dump(&p.matrix)
}

func NewPuzzle(maxX, maxY int, pieces ...Piece) *Puzzle {
	var matrix Matrix
	for x := 0; x < maxX; x++ {
		column := make([]int, maxY)
		matrix = append(matrix, column)
		for y := 0; y < maxY; y++ {
			switch {
			case y == 0 || y == maxY-1:
				column[y] = 1
			case x == 0 || x == maxX-1:
				column[y] = 1
			}
		}
	}
	return &Puzzle{matrix: matrix, pieces: pieces}
}

func (p *Puzzle) nextFreeCell(pnt Point) Point {
	fmt.Printf("\nnext free cell for: %s\n", pnt)
	for y := pnt.y; y < len((*p).matrix[0]); y++ {
		for x := 0; x < len((*p).matrix); x++ {
			fmt.Printf("%s:%d\t", Point{x, y}, (*p).matrix[x][y])
			if (y == pnt.y && x > pnt.x) || y != pnt.y {
				if p.matrix[x][y] == 0 {
					return Point{x, y}
				}
			}
		}
	}
	return Point{-1, -1}
}

func place(matrix *Matrix, piece *Piece, point Point) (*Matrix, bool) {
	success := true
	pieceX := len(((*piece).piecetrix)[0])
	pieceY := len((*piece).piecetrix)
	for x := 0; x < pieceX; x++ {
		for y := 0; y < pieceY; y++ {
			if (*matrix)[x+point.x][y+point.y] == 0 {
				(*matrix)[x+point.x][y+point.y] = (*piece).piecetrix[x][y]
			} else {
				success = false
			}
		}
	}
	dump(matrix)

	return nil, success
}

func dump(matrix *Matrix) {
	maxX, maxY := len(*matrix), len((*matrix)[0])
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			fmt.Printf("%d ", (*matrix)[x][y])
		}
		fmt.Println()
	}
	fmt.Printf("size: x:%d, y:%d\n", maxX, maxY)
}

func createEmptyMatrix(maxX, maxY int) Matrix {
	var matrix [][]int
	for x := 0; x < maxX+1; x++ {
		column := make([]int, maxY+1)
		matrix = append(matrix, column)
	}
	return matrix
}

func main() {
	puzzle := NewPuzzle(15, 10)
	puzzle.dump()
}
