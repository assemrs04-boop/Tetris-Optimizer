package solver

import (
	"testing"
	"tetris-optimizer/tetromino"
)

func TestSolveAndBacktrack(t *testing.T) {
	tetros := []tetromino.Tetromino{
		{Shape: []tetromino.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}}},
	}

	grid := [][]rune{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
	}

	if !backtrack(grid, tetros, 0, 4) {
		t.Errorf("Failed to solve grid")
	}

	if grid[0][0] != 'A' || grid[0][1] != 'A' || grid[0][2] != 'A' || grid[0][3] != 'A' {
		t.Errorf("Tetromino placed incorrectly")
	}
}
