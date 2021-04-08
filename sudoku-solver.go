package main

import (
	"flag"
	"fmt"
	"log"
)

var printPretty bool

func main() {
	flag.BoolVar(&printPretty, "pretty", false, "prints formatted output")
	flag.Parse()

	board := make([][]byte, 9)
	fmt.Println("Type in a sudoku board:")
	for i := 0; i < 9; i++ {
		var s string
		fmt.Scanln(&s)
		board[i] = []byte(s)
		if len(board[i]) != 9 {
			log.Fatal("Wrong number of cells")
		}
		for j := range board[i] {
			if '1' <= board[i][j] && board[i][j] <= '9' {
				board[i][j] -= '0'
			} else {
				board[i][j] = 0
			}
		}
	}

	solveSudoku(board)
}

func solveSudoku(board [][]byte) {
	solved := solve(board, 0, 0)

	if solved {
		printSudoku(board)
	} else {
		fmt.Printf("\nUnsolvable\n")
	}
}

func printSudoku(board [][]byte) {
	fmt.Printf("\nSolved:\n")
	fmt.Println("+---+---+---+")
	for i, boardRow := range board {
		fmt.Print("|")
		for j := range boardRow {
			fmt.Print(boardRow[j])
			if j%3 == 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i%3 == 2 {
			fmt.Println("+---+---+---+")
		}
	}
}

func solve(b [][]byte, x, y int) bool {
	nX, nY := nextXY(x, y)

	if x > 8 {
		return true
	}

	if b[x][y] != 0 {
		return solve(b, nX, nY)
	}

	sqx := (x / 3) * 3
	sqy := (y / 3) * 3

	var possible int16 = 2<<10 - 2
	for i := 0; i < 9; i++ {
		if b[i][y] != 0 {
			possible &= ^(1 << b[i][y])
		}
		if b[x][i] != 0 {
			possible &= ^(1 << b[x][i])
		}
		if b[sqx+i/3][sqy+i%3] != 0 {
			possible &= ^(1 << b[sqx+i/3][sqy+i%3])
		}
	}
	if possible == 0 {
		return false
	}

	for i := byte(1); i <= 9; i++ {
		if (possible & (1 << i)) > 0 {
			b[x][y] = i
			if solve(b, nX, nY) {
				return true
			}
		}
	}
	b[x][y] = 0
	return false
}

func nextXY(x, y int) (int, int) {
	if y == 8 {
		return x + 1, 0
	}
	return x, y + 1
}
