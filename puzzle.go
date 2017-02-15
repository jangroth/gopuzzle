// puzzle solver
package main

import (
	"fmt"
)

type Piece struct {
	piectrix [][]int
}

type Point struct {
	x, y int
}

type Puzzle struct {
	matrix     [][]int
	pieces     []Piece
	maxX, maxY int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p *Puzzle) dump() {
	for y := 0; y < p.maxY; y++ {
		for x := 0; x < p.maxX; x++ {
			fmt.Printf("%d ", p.matrix[x][y])
		}
		fmt.Println()
	}
	fmt.Printf("size: x:%d, y:%d\n", p.maxX, p.maxY)
}

func (p *Puzzle) nextFreeCell(pnt Point) Point {
	fmt.Printf("\nnext free cell for: %s\n", pnt)
	for y := pnt.y; y < p.maxY; y++ {
		for x := 0; x < p.maxX; x++ {
			fmt.Printf("%s:%d\n", Point{x, y}, p.matrix[x][y])
			if (y == pnt.y && x > pnt.x) || y != pnt.y {
				if p.matrix[x][y] == 0 {
					return Point{x, y}
				}
			}
		}
	}
	return Point{-1, -1}
}

func NewPuzzle(maxX, maxY int) *Puzzle {
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
	return &Puzzle{matrix: matrix, maxX: maxX, maxY: maxY}
}

func main() {
	puzzle := NewPuzzle(15, 10)
	puzzle.dump()
}
