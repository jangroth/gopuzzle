// puzzle solver
package main

import (
	"fmt"
)

// TODO: Point struct?

type Puzzle struct {
	Matrix [][]int
	Pieces [][]int
}

func (p Puzzle) dump() {
	for y := 0; y < len(p.Matrix); y++ {
		for x := 0; x < len(p.Matrix[y]); x++ {
			fmt.Printf("%d ", p.Matrix[y][x])
		}
		fmt.Println()
	}
	fmt.Printf("(x:%d, y:%d)\n", len(p.Matrix[0]), len(p.Matrix))
}

func NewPuzzle(maxX int, maxY int) *Puzzle {
	var matrix [][]int
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
	return &Puzzle{Matrix: matrix}
}

func main() {
	puzzle := NewPuzzle(15, 10)
	puzzle.dump()
}

func nextFreeCell(matrix [][]int, startX int, startY int) (int, int) {
	for x := startX; x < len(matrix[0]); x++ {
		for y := startY; y < len(matrix); y++ {
			fmt.Printf("x: %d, y:%d, matrix %d\n", x, y, matrix[x][y])
			if matrix[x][y] == 0 && (x != startX || y != startY) {
				return x, y
			}
		}
	}
	return -1, -1
}
