package main

import (
	"fmt"
	"log"
)

func main() {
	board := make([][]byte, 9)
	fmt.Println("Type in a sudoku board:")
	for i := 0; i < 9; i++ {
		var s string
		fmt.Scanln(&s)
		board[i] = []byte(s)
		if len(board[i]) != 9 {
			log.Fatal("Wrong number of cells")
		}
	}

	solveSudoku(board)
}

func solveSudoku(board [][]byte) {
	solved := solve(board, 0, 0)

	if solved {
		fmt.Printf("\nSolved:\n")
		for _, boardRow := range board {
			fmt.Println(string(boardRow))
		}
	} else {
		fmt.Printf("\nUnsolvable\n")
	}
}

func solve(b [][]byte, x, y int) bool {
	nX, nY := nextXY(x, y)

	if x > 8 {
		return true
	}

	if b[x][y] != '.' {
		return solve(b, nX, nY)
	}

	sqx := (x / 3) * 3
	sqy := (y / 3) * 3

	var possible int16 = 1022
	for i := 0; i < 9; i++ {
		if b[i][y] != '.' {
			possible &= ^(1 << (b[i][y] - '0'))
		}
		if b[x][i] != '.' {
			possible &= ^(1 << (b[x][i] - '0'))
		}
		if b[sqx+i/3][sqy+i%3] != '.' {
			possible &= ^(1 << (b[sqx+i/3][sqy+i%3] - '0'))
		}
	}
	if possible == 0 {
		return false
	}

	for i := byte(1); i <= 9; i++ {
		if (possible & (1 << i)) > 0 {
			b[x][y] = '0' + i
			if solve(b, nX, nY) {
				return true
			}
		}
	}
	b[x][y] = '.'
	return false
}

func nextXY(x, y int) (int, int) {
	if y == 8 {
		return x + 1, 0
	}
	return x, y + 1
}
