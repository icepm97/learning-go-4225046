package main

import (
	"fmt"
	"strconv"
	"strings"
)

// printBoard prints the Sudoku board with proper formatting.
func printBoard(board [][]int) {
	for i := 0; i < 9; i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("-" + strings.Repeat(" ", 21))
		}
		row := ""
		for j := 0; j < 9; j++ {
			if j%3 == 0 && j != 0 {
				row += "| "
			}
			if board[i][j] == 0 {
				row += " "
			} else {
				row += fmt.Sprintf("%d", board[i][j])
			}
		}
		fmt.Println(row)
	}
}

// isValidMove checks if placing a value at (row, col) is valid in Sudoku rules.
func isValidMove(board [][]int, row int, col int, val int) bool {
	// Check the row
	for j := 0; j < 9; j++ {
		if board[row][j] == val && j != col {
			return false
		}
	}

	// Check the column
	for i := 0; i < 9; i++ {
		if board[i][col] == val && i != row {
			return false
		}
	}

	// Check the 3x3 subgrid
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == val && (startRow+i != row || startCol+j != col) {
				return false
			}
		}
	}

	return true
}

// isSolved checks if the board has a valid solution.
func isSolved(board [][]int) bool {
	// Check if all cells are filled
	for _, row := range board {
		for _, num := range row {
			if num == 0 {
				return false
			}
		}
	}

	// Check rows for uniqueness (1-9)
	for _, row := range board {
		seen := make(map[int]bool)
		for _, num := range row {
			if num == 0 {
				continue
			}
			if seen[num] {
				return false
			}
			seen[num] = true
		}
	}

	// Check columns for uniqueness (1-9)
	for col := 0; col < 9; col++ {
		seen := make(map[int]bool)
		for row := 0; row < 9; row++ {
			num := board[row][col]
			if num == 0 {
				continue
			}
			if seen[num] {
				return false
			}
			seen[num] = true
		}
	}

	// Check subgrids for uniqueness (1-9)
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			seen := make(map[int]bool)
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					num := board[i+x][j+y]
					if num == 0 {
						continue
					}
					if seen[num] {
						return false
					}
					seen[num] = true
				}
			}
		}
	}

	return true
}

// main function to run the Sudoku game.
func main() {
	// Initial Sudoku board (0s represent empty cells)
	initialBoard := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 5, 0, 0, 9},
		{0, 0, 0, 0, 0, 0, 0, 7, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	// Make a deep copy of the board to avoid reference issues
	board := make([][]int, 9)
	for i := range initialBoard {
		board[i] = make([]int, 9)
		copy(board[i], initialBoard[i])
	}

	for {
		printBoard(board)

		if isSolved(board) {
			fmt.Println("Congratulations! You solved the puzzle!")
			break
		}

		fmt.Print("Enter your move (row col value) or 'quit': ")
		var input string
		fmt.Scanln(&input)

		if input == "quit" {
			fmt.Println("Game over!")
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Invalid input. Please enter row, column, value.")
			continue
		}

		rowStr := parts[0]
		colStr := parts[1]
		valStr := parts[2]

		var row, col, val int

		if err := strconv.Atoi(rowStr); err != nil {
			fmt.Println("Invalid input. Please enter integers.")
			continue
		} else if err := strconv.Atoi(colStr); err != nil {
			fmt.Println("Invalid input. Please enter integers.")
			continue
		} else if err := strconv.Atoi(valStr); err != nil {
			fmt.Println("Invalid input. Please enter integers.")
			continue
		} else {
			row, col, val = rowStr, colStr, valStr
		}

		// Validate input range (1-9)
		if row < 1 || row > 9 || col < 1 || col > 9 {
			fmt.Println("Row and column must be between 1-9.")
			continue
		}

		row-- // Convert to 0-based index
		col--

		if board[row][col] != 0 {
			fmt.Println("Cannot change existing clue.")
			continue
		}

		if val < 1 || val > 9 {
			fmt.Println("Value must be between 1-9.")
			continue
		}

		if isValidMove(board, row, col, val) {
			board[row][col] = val
		} else {
			fmt.Println("Invalid move: duplicate in row, column, or subgrid.")
		}
	}
}
