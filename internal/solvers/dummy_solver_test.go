package solvers

import (
	"github.com/mauroartizzu/sugoku/internal/board"
	"testing"
)

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
