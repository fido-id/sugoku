package solvers

import (
	"github.com/mauroartizzu/sugoku/internal/board"
	"testing"
)

// TestBacktrackingSolver tests the backtracking solver
func TestBackTrackingSolver(t *testing.T) {
	b := board.Board{}
	b.New([][]int{
		{0, 0, 0, 1, 0, 5, 0, 6, 8},
		{0, 0, 0, 0, 0, 0, 7, 0, 1},
		{9, 0, 1, 0, 0, 0, 0, 3, 0},
		{0, 0, 7, 0, 2, 6, 0, 0, 0},
		{5, 0, 0, 0, 0, 0, 0, 0, 3},
		{0, 0, 0, 8, 7, 0, 4, 0, 0},
		{0, 3, 0, 0, 0, 0, 8, 0, 5},
		{1, 0, 5, 0, 0, 0, 0, 0, 0},
		{7, 9, 0, 4, 0, 1, 0, 0, 0},
	})

	solvedBoard, _ := BackTrackingSolver(b)

	if !solvedBoard.Solved() {
		t.Fatalf("Board should be solved")
	}
}

func BenchmarkBackTrackingSolver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := getRandomSudoku()

		BackTrackingSolver(s)
	}
}
