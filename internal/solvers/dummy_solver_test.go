package solvers

import (
	"bufio"
	"fmt"
	"github.com/mauroartizzu/sugoku/internal/board"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func getRandomSudoku() board.Board {
	files, _ := os.ReadDir("../../samples/")

	rand.Seed(time.Now().Unix())
	fileIndex := rand.Intn(len(files))
	file, err := os.Open(fmt.Sprintf("../../samples/%s", files[fileIndex].Name()))
	if err != nil {
		fmt.Println("Impossibile aprire il file:", os.Args[1])
		panic(err)
	}

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

// TestDummySolver tests the dummy solver
func TestDummySolver(t *testing.T) {
	b := board.Board{}
	b.New([][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
	})

	solvedBoard, _ := DummySolver(b)

	if !solvedBoard.Solved() {
		t.Fatalf("Board should be solved")
	}
}
