// puzzle solver
package main

import (
	"fmt"
)

// types

type Point struct {
	x, y int
}

type Matrix [][]int

type Piece struct {
	matrix Matrix
}

type Puzzle struct {
	matrix Matrix
	pieces []Piece
}

// methods

func NewMatrix(maxX, maxY int) *Matrix {
	var matrix Matrix
	for x := 0; x < maxX; x++ {
		column := make([]int, maxY)
		matrix = append(matrix, column)
	}
	return &matrix
}

func (matrix *Matrix) dump() {
	maxX, maxY := len(*matrix), len((*matrix)[0])
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			fmt.Printf("%d ", (*matrix)[x][y])
		}
		fmt.Println()
	}
	fmt.Printf("size: x:%d, y:%d\n", maxX, maxY)
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
	matrix := NewMatrix(maxX+1, maxY+1)
	for _, val := range points {
		(*matrix)[val.x][val.y] = value
	}
	return &Piece{matrix: *matrix}
}

func (piece *Piece) mirror() *Piece {
	maxX, maxY := len(piece.matrix), len((piece.matrix)[0])
	mirror := NewMatrix(maxX, maxY)
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			(*mirror)[x][y] = (*piece).matrix[x][maxY-y-1]
		}
	}
	return &Piece{matrix: *mirror}
}

func NewPuzzle(maxX, maxY int, pieces ...Piece) *Puzzle {
	matrix := NewMatrix(maxX, maxY)
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if (y == 0 || y == maxY-1) || (x == 0 || x == maxX-1) {
				(*matrix)[x][y] = 1
			}
		}
	}
	return &Puzzle{matrix: *matrix, pieces: pieces}
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

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

// functions

func place(matrix *Matrix, piece *Piece, point Point) (*Matrix, bool) {
	pieceX := len(((*piece).matrix)[0])
	pieceY := len((*piece).matrix)
	for x := 0; x < pieceX; x++ {
		for y := 0; y < pieceY; y++ {
			if (*piece).matrix[x][y] != 0 && ((*matrix)[x+point.x][y+point.y] != 0) {
				return matrix, false
			}
		}
	}
	for x := 0; x < pieceX; x++ {
		for y := 0; y < pieceY; y++ {
			if (*piece).matrix[x][y] != 0 {
				(*matrix)[x+point.x][y+point.y] = (*piece).matrix[x][y]
			}
		}
	}
	return matrix, true
}

func main() {
	puzzle := NewPuzzle(15, 10)
	puzzle.matrix.dump()
}
