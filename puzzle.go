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
	matrix           Matrix
	permutatedPieces [][]*Piece
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

func (matrix *Matrix) nextCell(pnt Point) (nextPoint Point, ok bool) {
	maxX, maxY := (*matrix).dimensions()
	switch {
	case pnt.x < maxX-1:
		return Point{pnt.x + 1, pnt.y}, true
	case pnt.x == maxX-1 && pnt.y < maxY-1:
		return Point{0, pnt.y + 1}, true
	default:
		return nextPoint, false
	}
}

func (matrix *Matrix) place(piece *Piece, point Point) (success bool) {
	matrixX, matrixY := (*matrix).dimensions()
	pieceX, pieceY := (*piece).matrix.dimensions()
	if point.x+pieceX > matrixX || point.y+pieceY > matrixY {
		return false
	}
	for x := 0; x < pieceX; x++ {
		for y := 0; y < pieceY; y++ {
			if (*piece).matrix[x][y] != 0 && ((*matrix)[x+point.x][y+point.y] != 0) {
				return false
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
	return true
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

func NewPuzzle(maxX, maxY int, pieces ...Piece) Puzzle {
	matrix := NewMatrix(maxX, maxY)
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if (y == 0 || y == maxY-1) || (x == 0 || x == maxX-1) {
				(*matrix)[x][y] = 1
			}
		}
	}
	var permutatedPieces [][]*Piece
	for _, piece := range pieces {
		permutatedPieces = append(permutatedPieces, piece.permutate())
	}

	return Puzzle{matrix: *matrix, permutatedPieces: permutatedPieces}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p *Puzzle) removePermuatedPiece(index int) {
	p.permutatedPieces = append(p.permutatedPieces[:index], p.permutatedPieces[index+1:]...)
}

func (p *Puzzle) dump() {
	fmt.Printf("Dump puzzle (%p):\n", &p)
	p.matrix.dump()
	for index, val := range p.permutatedPieces {
		fmt.Printf("Piece #%d (%d permutations)\n", index, len(val))
		//val[0].matrix.dump()
	}
}

// functions

func Solve(p Puzzle, startingPnt Point) (success bool) {
	fmt.Printf("---> Call for %s: ", startingPnt)
	p.dump()
	if len(p.permutatedPieces) == 0 {
		fmt.Println("Solved!")
		// solved
		return true
	} else {
		//  not solved yet. Try remaining pieces
		for pp_index, permutatedPiece := range p.permutatedPieces {
			for p_index, piece := range permutatedPiece {
				if p.matrix.place(piece, startingPnt) {
					fmt.Printf("Managed to place piece #%d, %d. permutation!\n", pp_index, p_index)
					p.removePermuatedPiece(pp_index)
					found_solution := Solve(p, startingPnt)
					if found_solution {
						// found solution, no need to try anything else
						return true
					} else {
						break
					}
				}
			}
		}
		nextPoint, hasNext := p.matrix.nextCell(startingPnt)
		if hasNext {
			return Solve(p, nextPoint)
		}
	}
	return false
}

func main() {
	puzzle := NewPuzzle(15, 10)
	puzzle.matrix.dump()
}
