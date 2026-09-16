package main

import (
	"fmt"
	"os"
	"tetris-optimizer/tetris-solver"
	"tetris-optimizer/tetromino"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	tetros, err := tetromino.ReadAndValidate(os.Args[1])
	if err != nil || tetros == nil {
		fmt.Println("ERROR")
		return
	}

	solver.Solve(tetros)
}
