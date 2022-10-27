package main

import (
	"bufio"
	"fmt"
	"github.com/buger/goterm"
	"github.com/mauroartizzu/sugoku/internal/board"
	"github.com/mauroartizzu/sugoku/internal/solvers"
	"os"
	"strconv"
	"time"
)

// getSudoku reads a sudoku from a file
func getSudoku(file *os.File) board.Board {

	scan := bufio.NewScanner(file)

	scan.Split(bufio.ScanLines)

	matrix := make([][]int, 0)

	for scan.Scan() {
		row := make([]int, 0)
		for _, c := range scan.Text() {
			result, _ := strconv.Atoi(string(c))
			row = append(row, result)
		}
		matrix = append(matrix, row)
	}

	b := board.Board{}

	b.New(matrix)

	return b
}

// main is the entry point of the program
func main() {

	if len(os.Args) < 2 {
		fmt.Println("Specificare un file col sudoku!")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Impossibile aprire il file:", os.Args[1])
		panic(err)
	}
	defer file.Close()

	b := getSudoku(file)
	goterm.Flush()

	b.PrettyPrintInPlace()

	fmt.Println("Scegliere il solver da usare:")
	fmt.Println("1. DummySolver")
	fmt.Println("2. SimpleSolver")
	fmt.Println("3. BackTrackingSolver")
	//fmt.Println("4. PencilMarkSolver")

	// read the choice
	var choice int
	fmt.Scanln(&choice)

	var start time.Time
	// solve the sudoku
	switch choice {
	case 1:
		start = time.Now()
		b, err = solvers.DummySolver(b)
		if err != nil {
			panic(err)
		}
	case 2:
		start = time.Now()
		b, err = solvers.SimpleSolver(b)
		if err != nil {
			panic(err)
		}
	case 3:
		start = time.Now()
		b, err = solvers.BackTrackingSolver(b)
		if err != nil {
			panic(err)
		}
	//case 4:
	//	start = time.Now()
	//	b, err = solvers.PencilMarkSolver(b)
	//	if err != nil {
	//		panic(err)
	//	}
	default:
		goterm.Println("Scelta non valida!")
		return
	}

	goterm.Flush()
	// start the timer

	// stop the timer
	elapsed := time.Since(start)

	// print the result
	b.PrettyPrintInPlace()

	if b.Solved() {
		fmt.Println("Sudoku risolto!")
	} else {
		fmt.Println("Sudoku non risolto!")
	}

	// print the time elapsed
	fmt.Println("Tempo impiegato:", elapsed)
}
