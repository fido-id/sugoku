package board

import (
	"fmt"
	"github.com/buger/goterm"
	"math"
	"time"
)

// Board is a sudoku board
type Board struct {
	board [][]int
}

// New creates a new sudoku board
func (b *Board) New(initialBoard ...[][]int) {

	b.board = make([][]int, 9)
	for i := 0; i < 9; i++ {
		b.board[i] = make([]int, 9)
		if len(initialBoard) != 0 {
			for j := 0; j < 9; j++ {
				b.board[i][j] = initialBoard[0][i][j]
			}
		}
	}

}

// Print the sudoku board
func (b *Board) Print() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", b.board[i][j])
			if (j+1)%3 == 0 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
}

// PrettyPrintInPlace prints the sudoku board in place
func (b *Board) PrettyPrintInPlace() {
	//goterm.MoveCursor(1, 1)
	goterm.Println("┌───────┬───────┬───────┐")
	for i := 0; i < 9; i++ {
		goterm.Print("│")
		for j := 0; j < 9; j++ {
			switch j {
			case 0, 3, 6:
				goterm.Printf(" %d", b.board[i][j])
			case 2, 5, 8:
				goterm.Printf("%d ", b.board[i][j])
			default:
				goterm.Printf(" %d ", b.board[i][j])
			}
			if (j+1)%3 == 0 {
				goterm.Print("│")
			}
		}
		goterm.Println()
		if (i+1)%3 == 0 && i+1 != 9 {
			goterm.Println("├───────┼───────┼───────┤")
		}
	}
	goterm.Println("└───────┴───────┴───────┘")
	goterm.Flush()
	time.Sleep(time.Second)
	// move cursor up 9 left 9
}

// Set a value in the sudoku board
func (b *Board) Set(row int, col int, value int) {
	b.board[row][col] = value
}

// Get a value from the sudoku board
func (b *Board) Get(row int, col int) int {
	return b.board[row][col]
}

// checkRow checks if a row equals 45
func (b *Board) checkRow(row int) bool {
	sum := 0
	for i := 0; i < 9; i++ {
		sum += b.board[row][i]
	}
	return sum == 45
}

// checkCol checks if a column equals 45
func (b *Board) checkCol(col int) bool {
	sum := 0
	for i := 0; i < 9; i++ {
		sum += b.board[i][col]
	}
	return sum == 45
}

// checkQuadrant checks if a quadrant equals 45
func (b *Board) checkQuadrant(row int, col int) bool {
	sum := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += b.board[row+i][col+j]
		}
	}
	return sum == 45
}

// IsEmptyCell checks if a cell is empty
func (b *Board) IsEmptyCell(rowIndex int, colIndex int) bool {
	if b.Get(rowIndex, colIndex) == 0 {
		return true
	}
	return false
}

// IsAllowedValue checks if a value is allowed in a cell
func (b *Board) IsAllowedValue(v int, rowIndex int, colIndex int) bool {
	// find quadrant top left cell coordinates
	x := int(math.Floor(float64(rowIndex/3)) * 3)
	y := int(math.Floor(float64(colIndex/3)) * 3)

	for i := 0; i < 9; i++ {
		if b.Get(rowIndex, i) == v {
			return false
		}
		if b.Get(i, colIndex) == v {
			return false
		}
		if b.Get(x+int(math.Floor(float64(i/3))), y+(i%3)) == v {
			return false
		}
	}

	return true
}

// Solved checks if the sudoku board is solved
func (b *Board) Solved() bool {
	for i := 0; i < 9; i++ {
		if !b.checkRow(i) {
			return false
		}
		if !b.checkCol(i) {
			return false
		}
		if !b.checkQuadrant(int(math.Floor(float64(i/3))*3), i%3*3) {
			return false
		}
	}
	return true
}
