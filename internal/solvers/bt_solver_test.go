package solvers

import (
	"github.com/mauroartizzu/sugoku/internal/board"
	"testing"
)

// TestDummySolver tests the dummy solver
func TestBackTrackingSolver(t *testing.T) {
	b := board.Board{}
	b.New([][]int{
		{1, 0, 0, 4, 0, 0, 0, 0, 0},
		{4, 5, 6, 0, 0, 9, 1, 2, 3},
		{7, 8, 9, 0, 0, 3, 4, 5, 0},
		{2, 3, 4, 0, 0, 7, 8, 9, 1},
		{5, 6, 7, 0, 0, 1, 2, 3, 4},
		{8, 9, 1, 0, 0, 4, 5, 6, 7},
		{3, 4, 5, 0, 0, 8, 0, 1, 2},
		{6, 7, 8, 0, 0, 2, 3, 0, 5},
		{9, 1, 2, 0, 0, 0, 6, 7, 0},
	})

	solvedBoard, _ := BackTrackingSolver(b)

	if !solvedBoard.Solved() {
		t.Fatalf("Board should be solved")
	}
}
