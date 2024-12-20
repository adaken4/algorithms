package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 9 {
		fmt.Println("ERROR")
		return
	}
	board := [][]string{}
	for _, row := range args {
		if len(row) != 9 {
			fmt.Println("ERROR")
			return
		}
		board = append(board, strings.Split(row, ""))
	}
	if solve(0, 0, board) {
		printBoard(board)
	} else {
		fmt.Println("No solution found!")
	}
}

func solve(row, col int, board [][]string) bool {
	if row == 9 {
		return true
	}
	if col == 9 {
		return solve(row+1, 0, board)
	}
	if board[row][col] != "." {
		return solve(row, col+1, board)
	}
	for val := '1'; val <= '9'; val++ {
		board[row][col] = string(val)
		// fmt.Println("\033[38;5;22m", board, "\u001b[0m")
		// time.Sleep(20 * time.Millisecond)
		if validPlacement(row, col, board) {
			if solve(row, col+1, board) {
				return true
			}
		}
		board[row][col] = "."
	}
	return false
}

func validPlacement(row, col int, board [][]string) bool {
	num := board[row][col]

	for c := 0; c <= 8; c++ {
		if board[row][c] == num && c != col {
			return false
		}
	}
	for r := 0; r <= 8; r++ {
		if board[r][col] == num && r != row {
			return false
		}
	}
	rowStart := (row / 3) * 3
	colStart := (col / 3) * 3
	for x := rowStart; x < rowStart+3; x++ {
		for y := colStart; y < colStart+3; y++ {
			if board[x][y] == num && x != row && y != col {
				return false
			}
		}
	}
	return true
}

func printBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(strings.Join(row, " "))
	}
}
