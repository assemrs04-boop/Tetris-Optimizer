package solver

import (
	"fmt"
	"math"
	"tetris-optimizer/tetromino"
)

func Solve(tetros []tetromino.Tetromino) {
	size := int(math.Ceil(math.Sqrt(float64(len(tetros) * 4))))

	for {
		grid := make([][]rune, size)
		for i := range grid {
			grid[i] = make([]rune, size)
			for j := range grid[i] {
				grid[i][j] = '.'
			}
		}

		if backtrack(grid, tetros, 0, size) {
			printGrid(grid)
			return
		}
		size++
	}
}

func backtrack(grid [][]rune, tetros []tetromino.Tetromino, index, size int) bool {
	if index == len(tetros) {
		return true
	}

	letter := rune('A' + index)
	t := tetros[index]

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if canPlace(grid, t, x, y, size) {
				place(grid, t, x, y, letter)
				if backtrack(grid, tetros, index+1, size) {
					return true
				}
				place(grid, t, x, y, '.')
			}
		}
	}
	return false
}

func canPlace(grid [][]rune, t tetromino.Tetromino, x, y, size int) bool {
	for _, p := range t.Shape {
		nx, ny := x+p.X, y+p.Y
		if nx >= size || ny >= size || grid[ny][nx] != '.' {
			return false
		}
	}
	return true
}

func place(grid [][]rune, t tetromino.Tetromino, x, y int, char rune) {
	for _, p := range t.Shape {
		grid[y+p.Y][x+p.X] = char
	}
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
