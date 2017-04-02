// puzzle solver
package main

import (
	"fmt"
	"strings"
)

// types

type Borderfun func(x, y, maxX, maxY int) (hasBorder bool)

type Point struct {
	x, y int
}

type Matrix [][]int

type Piece struct {
	value  int
	matrix Matrix
}

type Puzzle struct {
	matrix           Matrix
	permutatedPieces [][]*Piece
	solution         map[int]*Point
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
			val := (*matrix)[x][y]
			if val >= 2 {
				result += fmt.Sprintf("%d ", (*matrix)[x][y])
			} else if val == 1 {
				result += fmt.Sprintf("X ")
			} else {
				result += fmt.Sprintf("_ ")
			}
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

func (matrix *Matrix) testAndPlace(piece *Piece, point Point) (success bool) {
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

func (matrix *Matrix) remove(piece *Piece, point Point) {
	(*matrix).dump()
	pieceX, pieceY := (*piece).matrix.dimensions()
	for x := 0; x < pieceX; x++ {
		for y := 0; y < pieceY; y++ {
			if (*piece).matrix[x][y] != 0 {
				(*matrix)[x+point.x][y+point.y] = 0
			}
		}
	}
	matrixX, matrixY := (*matrix).dimensions()
	for y := 0; y < matrixY; y++ {
		for x := 0; x < matrixX; x++ {
			if (*matrix)[x][y] == (*piece).value {
				fmt.Printf("sanity check failed for %d\n", (*piece).value)
				(*matrix).dump()
				return
			}
		}
	}
	fmt.Printf("sanity check passed for %d\n", (*piece).value)
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
	return &Piece{value: value, matrix: *matrix}
}

func (piece *Piece) mirror() *Piece {
	maxX, maxY := piece.matrix.dimensions()
	mirrored := NewMatrix(maxX, maxY)
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			(*mirrored)[x][y] = (*piece).matrix[x][maxY-y-1]
		}
	}
	return &Piece{matrix: *mirrored, value: piece.value}
}

func (piece *Piece) rotate() *Piece {
	maxX, maxY := piece.matrix.dimensions()
	rotated := NewMatrix(maxY, maxX)
	for x := 0; x < maxY; x++ {
		for y := 0; y < maxX; y++ {
			(*rotated)[maxY-x-1][y] = (*piece).matrix[y][x]
		}
	}
	return &Piece{matrix: *rotated, value: piece.value}
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

func simpleBorder(x, y, maxX, maxY int) bool {
	return (y == 0 || y == maxY-1) || (x == 0 || x == maxX-1)
}

func niftyFiftyBorder(x, y, maxX, maxY int) bool {
	if maxX != 21 || maxY != 21 {
		return false
	} else {
		var result bool
		switch {
		case y == 1:
			result = x >= 8
		case y == 2 || y == 3:
			result = x >= 9
		case y == 4 || y == 5:
			result = x >= 12
		case y == 6:
			result = x >= 15
		case y == 7 || y == 8 || y == 9:
			result = x >= 16
		case y == 10:
			result = x <= 2 || x >= 17
		case y == 11:
			result = x <= 5 || x >= 17
		case y == 12 || y == 13:
			result = x <= 5
		case y == 14:
			result = x <= 6
		case y == 15 || y == 16:
			result = x <= 7
		case y == 17 || y == 18:
			result = x <= 9
		case y == 19:
			result = x <= 10
		}
		return result || simpleBorder(x, y, maxX, maxY)
	}
}

func NewPuzzle(maxX, maxY int, hasBorder Borderfun, pieces ...Piece) Puzzle {
	matrix := NewMatrix(maxX, maxY)
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if hasBorder(x, y, maxX, maxY) {
				(*matrix)[x][y] = 1
			}
		}
	}
	var permutatedPieces [][]*Piece
	for _, piece := range pieces {
		permutatedPieces = append(permutatedPieces, piece.permutate())
	}
	return Puzzle{matrix: *matrix, permutatedPieces: permutatedPieces, solution: make(map[int]*Point)}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p *Puzzle) dump() {
	fmt.Printf("Dump puzzle (%p):\n", &p)
	p.matrix.dump()
	for pp_index, permutatedPiece := range p.permutatedPieces {
		fmt.Printf("Piece #%d (%d permutations)\n", pp_index, len(permutatedPiece))
		for p_index, piece := range permutatedPiece {
			fmt.Printf("index %d, value %d\n", p_index, piece.value)
			piece.matrix.dump()
		}
	}
}

// functions

func Solve(p Puzzle, startingPnt Point) (success bool) {
	if len(p.permutatedPieces) == len(p.solution) {
		p.dump()
		fmt.Println("Solved!")
		return true
	} else {
		//  not solved yet. Try remaining pieces
		for pp_index, permutatedPiece := range p.permutatedPieces {
			_, alreadyInSolution := p.solution[pp_index]
			if !alreadyInSolution {
				for p_index, piece := range permutatedPiece {
					if p.matrix.testAndPlace(piece, startingPnt) {
						fmt.Printf("\n%s ", strings.Repeat(">", len(p.solution)+1))
						fmt.Printf("Placing piece #%d(%d)/%d at %s!\n", piece.value, piece.matrix[0][0], p_index, startingPnt)
						p.solution[pp_index] = &startingPnt
						found_solution := Solve(p, Point{0, 0})
						if found_solution {
							// found solution, no need to try anything else
							return true
						} else {
							// no solution found, remove and try next permutation / piece
							fmt.Printf("%s Removing piece #%d/%d at %s!\n", strings.Repeat(">", len(p.solution)+1), pp_index, p_index, startingPnt)
							p.matrix.remove(piece, startingPnt)
							delete(p.solution, pp_index)
						}
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
	//
}
