// puzzle solver
package main

import (
	"fmt"
)

// TODO: Point struct?

type Puzzle struct {
	Matrix [][]int
	Pieces [][]int
	maxX   int
	maxY   int
}

func (p Puzzle) dump() {
	for y := 0; y < p.maxY; y++ {
		for x := 0; x < p.maxX; x++ {
			fmt.Printf("%d ", p.Matrix[x][y])
		}
		fmt.Println()
	}
	fmt.Printf("(x:%d, y:%d)\n", p.maxX, p.maxY)
}

func (p Puzzle) nextFreeCell(startX int, startY int) (int, int) {
	fmt.Printf("\nnext free cell %d - %d\n", startX, startY)
	for y := startY; y < p.maxY; y++ {
		for x := 0; x < p.maxX; x++ {
			fmt.Printf("x: %d, y:%d, matrix %d\n", x, y, p.Matrix[x][y])
			if (y == startY && x > startX) || y != startY {
				if p.Matrix[x][y] == 0 {
					return x, y
				}
			}
		}
	}
	return -1, -1
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
	return &Puzzle{Matrix: matrix, maxX: maxX, maxY: maxY}
}

func main() {
	puzzle := NewPuzzle(15, 10)
	puzzle.dump()
}
