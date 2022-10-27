package solvers

import (
	"github.com/mauroartizzu/sugoku/internal/board"
	"testing"
)

func TestSimpleSolver(t *testing.T) {
	b := board.Board{}
	b.New([][]int{
		{0, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 0, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 0, 1, 2, 3, 4, 5, 6},
		{2, 3, 4, 0, 6, 7, 8, 9, 1},
		{5, 6, 7, 8, 0, 1, 2, 3, 4},
		{8, 9, 1, 2, 3, 0, 5, 6, 7},
		{3, 4, 5, 6, 7, 8, 0, 1, 2},
		{6, 7, 8, 9, 1, 2, 3, 0, 5},
		{9, 1, 2, 3, 4, 5, 6, 7, 0},
	})

	solvedBoard, _ := SimpleSolver(b)

	if !solvedBoard.Solved() {
		t.Fatalf("Board should be solved")
	}
}

func BenchmarkSimpleSolver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := getRandomSudoku()

		SimpleSolver(s)
	}
}
