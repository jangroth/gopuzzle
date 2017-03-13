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

func (matrix *Matrix) dimensions() (maxX, maxY int) {
	return len(*matrix), len((*matrix)[0])
}

func (matrix *Matrix) toString() string {
	maxX, maxY := matrix.dimensions()
	var result string
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			result += fmt.Sprintf("%d ", (*matrix)[x][y])
		}
		result += "\n"
	}
	return result
}

func (matrix *Matrix) dump() {
	maxX, maxY := matrix.dimensions()
	fmt.Printf("%ssize: x:%d, y:%d\n", matrix.toString(), maxX, maxY)
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
	maxX, maxY := piece.matrix.dimensions()
	mirrored := NewMatrix(maxX, maxY)
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			(*mirrored)[x][y] = (*piece).matrix[x][maxY-y-1]
		}
	}
	return &Piece{matrix: *mirrored}
}

func (piece *Piece) rotate() *Piece {
	maxX, maxY := piece.matrix.dimensions()
	rotated := NewMatrix(maxY, maxX)
	for x := 0; x < maxY; x++ {
		for y := 0; y < maxX; y++ {
			(*rotated)[maxY-x-1][y] = (*piece).matrix[y][x]
		}
	}
	return &Piece{matrix: *rotated}
}

func (piece *Piece) permutate() []*Piece {
	resultMap := make(map[string]*Piece)
	resultMap[piece.matrix.toString()] = piece
	for i := 0; i < 7; i++ {
		if i == 3 {
			piece = piece.mirror()
		}
		piece = piece.rotate()
		resultMap[piece.matrix.toString()] = piece
	}
	var results []*Piece
	for _, result := range resultMap {
		results = append(results, result)
	}
	return results
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
	maxX, maxY := (*p).matrix.dimensions()
	for y := pnt.y; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
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

func place(matrix *Matrix, piece *Piece, point Point) (modified_matrix *Matrix, success bool) {
	pieceX, pieceY := (*piece).matrix.dimensions()
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
