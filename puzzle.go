/*
Solver for nifty-fifty puzzles.

Also works for other puzzles that have a two dimensional matrix structure of
blocked / unblocked fields and pieces to fit into the matrix that can also
be described similarly.
*/
package main

import (
	"fmt"
	"strings"
)

// types

type matrix [][]int

// Point defines a point in a 2-dimensional matrix
type Point struct {
	x, y int
}

// Piece defines a a 2-dimensional matrix that is a piece of the puzzle
type Piece struct {
	value  int
	matrix matrix
}

// Puzzle defines a a 2-dimensional matrix together with the pieces that should fit into it
type Puzzle struct {
	matrix           matrix
	permutatedPieces [][]*Piece
	solution         map[int]*Point
}

// BorderFun defines a function that describes the borders of a matrix
type BorderFun func(x, y, maxX, maxY int) (hasBorder bool)

// methods

func newMatrix(maxX, maxY int, borderFun BorderFun) *matrix {
	var matrix matrix
	for x := 0; x < maxX; x++ {
		column := make([]int, maxY)
		matrix = append(matrix, column)
		for y := 0; y < maxY; y++ {
			if borderFun(x, y, maxX, maxY) {
				matrix[x][y] = 1
			}
		}
	}
	return &matrix
}

func (matrix *matrix) dimensions() (maxX, maxY int) {
	return len(*matrix), len((*matrix)[0])
}

func (matrix *matrix) toString() string {
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

func (matrix *matrix) dump() {
	maxX, maxY := matrix.dimensions()
	fmt.Printf("%ssize: x:%d, y:%d\n", matrix.toString(), maxX, maxY)
}

func (matrix *matrix) nextCell(pnt Point) (nextPoint Point, ok bool) {
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

func (matrix *matrix) testAndPlace(piece *Piece, point Point) (success bool) {
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

func (matrix *matrix) remove(piece *Piece, point Point) {
	(*matrix).dump()
	pieceX, pieceY := (*piece).matrix.dimensions()
	for x := 0; x < pieceX; x++ {
		for y := 0; y < pieceY; y++ {
			if (*piece).matrix[x][y] != 0 {
				(*matrix)[x+point.x][y+point.y] = 0
			}
		}
	}
}

// NewPiece creates a piece of the puzzle
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
	matrix := newMatrix(maxX+1, maxY+1, noBorder)
	for _, val := range points {
		(*matrix)[val.x][val.y] = value
	}
	return &Piece{value: value, matrix: *matrix}
}

func (piece Piece) mirror() *Piece {
	maxX, maxY := piece.matrix.dimensions()
	mirrored := newMatrix(maxX, maxY, noBorder)
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			(*mirrored)[x][y] = piece.matrix[x][maxY-y-1]
		}
	}
	return &Piece{matrix: *mirrored, value: piece.value}
}

func (piece Piece) rotate() *Piece {
	maxX, maxY := piece.matrix.dimensions()
	rotated := newMatrix(maxY, maxX, noBorder)
	for x := 0; x < maxY; x++ {
		for y := 0; y < maxX; y++ {
			(*rotated)[maxY-x-1][y] = piece.matrix[y][x]
		}
	}
	return &Piece{matrix: *rotated, value: piece.value}
}

func (piece Piece) permutate() (resultList []*Piece) {
	resultMap := make(map[string]*Piece)
	workpiece := &piece
	resultMap[piece.matrix.toString()] = workpiece
	for i := 0; i < 3; i++ {
		workpiece = workpiece.rotate()
		resultMap[workpiece.matrix.toString()] = workpiece
	}
	workpiece = workpiece.mirror()
	resultMap[workpiece.matrix.toString()] = workpiece
	for i := 0; i < 3; i++ {
		workpiece = workpiece.rotate()
		resultMap[workpiece.matrix.toString()] = workpiece
	}
	for _, result := range resultMap {
		resultList = append(resultList, result)
	}
	return resultList
}

// NewPuzzle creates a puzzle
func NewPuzzle(maxX, maxY int, borderFun BorderFun, pieces ...Piece) Puzzle {
	matrix := newMatrix(maxX, maxY, borderFun)
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
	for ppIndex, permutatedPiece := range p.permutatedPieces {
		fmt.Printf("Piece #%d (%d permutations)\n", ppIndex, len(permutatedPiece))
		for pIndex, piece := range permutatedPiece {
			fmt.Printf("index %d, value %d\n", pIndex, piece.value)
			piece.matrix.dump()
		}
	}
}

// functions

func noBorder(x, y, maxX, maxY int) bool {
	return false
}

func simpleBorder(x, y, maxX, maxY int) bool {
	return (y == 0 || y == maxY-1) || (x == 0 || x == maxX-1)
}

func niftyFiftyBorder(x, y, maxX, maxY int) bool {
	if maxX != 21 || maxY != 21 {
		return false
	}
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

func solve(p Puzzle, startingPnt Point) (success bool) {
	fmt.Printf("Okay lets solve this. Solution so far:%v, startingPoint:%s\n", p.solution, startingPnt)
	if len(p.permutatedPieces) == len(p.solution) {
		p.dump()
		fmt.Println("Solved!")
		return true
	}
	//  not solved yet. Try remaining pieces
	for outerIndex, permutatedPiece := range p.permutatedPieces {
		fmt.Printf("Outer index %d, solution: %v\n", outerIndex, p.solution)
		_, alreadyInSolution := p.solution[outerIndex]
		if !alreadyInSolution {
			for innerIndex, piece := range permutatedPiece {
				if p.matrix.testAndPlace(piece, startingPnt) {
					fmt.Printf("%s Placing piece #%d/%d at %s!\n", strings.Repeat(">", len(p.solution)+1), outerIndex, innerIndex, startingPnt)
					p.solution[outerIndex] = &startingPnt
					foundSolution := solve(p, Point{0, 0})
					if foundSolution {
						return true
					}
					// no solution found, remove and try next permutation / piece
					fmt.Printf("%s Removing piece #%d/%d at %s!\n", strings.Repeat(">", len(p.solution)+1), outerIndex, innerIndex, startingPnt)
					p.matrix.remove(piece, startingPnt)
					delete(p.solution, outerIndex)
				}
			}
		}
	}
	nextPoint, hasNext := p.matrix.nextCell(startingPnt)
	if hasNext {
		return solve(p, nextPoint)
	}
	return false
}

// Solve solves the puzzle.
func Solve(p Puzzle) (success bool) {
	return solve(p, Point{0, 0})
}

func main() {
	//
}
