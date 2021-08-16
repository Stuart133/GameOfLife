package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const board_size = 20
const grid_size = board_size + 2 // Add a buffer layer so edge cells have correct neighbour calculations
const tick_delay = time.Millisecond * 100

type Board [][]bool

func main() {
	board := make_board(grid_size)
	board.add_lwss(2, 10)
	board.draw()
	time.Sleep(tick_delay)

	for i := 0; i < 50; i++ {
		board = board.compute_new_state()
		board.draw()
		time.Sleep(tick_delay)
	}
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Compute new board state
func (b Board) compute_new_state() Board {
	nb := make_board(grid_size)

	for i := range b {
		for j := range b[i] {
			num := b.num_neighbours(i, j)
			// Is the cell alive next round?
			if b[i][j] && num == 2 || num == 3 {
				nb[i][j] = true
			}
		}
	}

	return nb
}

func (board Board) num_neighbours(x int, y int) int {
	num := 0

	for _, x_offset := range []int{-1, 0, 1} {
		for _, y_offset := range []int{-1, 0, 1} {
			// Ignore the target cell
			if x_offset == 0 && y_offset == 0 {
				continue
			}

			// Check y overflow
			if y+y_offset < 0 || y+y_offset >= grid_size {
				continue
			}

			// Check x overflow
			if x+x_offset < 0 || x+x_offset >= grid_size {
				continue
			}

			if board[x+x_offset][y+y_offset] {
				num++
			}
		}
	}

	return num
}

// Draw board on console
func (board Board) draw() {
	clear()
	// There is a 2 layer buffer in all directions, don't draw this
	for i := 2; i < board_size; i++ {
		for j := 2; j < board_size; j++ {
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
func make_board(size int) Board {
	board := make(Board, size)
	for i := range board {
		board[i] = make([]bool, size)
	}

	return board
}

// Create a glider at given coordinates
func (board Board) add_glider(x int, y int) {
	board[x][y+2] = true
	board[x+1][y] = true
	board[x+1][y+2] = true
	board[x+2][y+1] = true
	board[x+2][y+2] = true
}

// Create a LWSS at given coordinates
func (board Board) add_lwss(x int, y int) {

	board[x][y+1] = true
	board[x][y+4] = true

	board[x+1][y] = true

	board[x+2][y] = true
	board[x+2][y+4] = true

	board[x+3][y] = true
	board[x+3][y+1] = true
	board[x+3][y+2] = true
	board[x+3][y+3] = true
}
