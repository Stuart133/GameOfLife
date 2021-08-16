package main

import "fmt"

func main() {
	board := make_board(10)
	add_glider(board, 0, 0)
	draw(board)
}

// Draw board on console
func draw(board [][]bool) {
	fmt.Println("##########")
	for i := range board {
		for j := range board[i] {
			if board[i][j] {
				fmt.Print("x")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Print("\n")
	}
}

// Initialize empty board
func make_board(size int) [][]bool {
	board := make([][]bool, 10)
	for i := range board {
		board[i] = make([]bool, 10)
	}

	return board
}

// Create a glider at given coordinates
func add_glider(board [][]bool, x int, y int) {

	board[x][y+2] = true
	board[x+1][y] = true
	board[x+1][y+2] = true
	board[x+2][y+1] = true
	board[x+2][y+2] = true
}
