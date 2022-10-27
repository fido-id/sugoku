package solvers

import (
	"github.com/mauroartizzu/sugoku/internal/board"
	"testing"
)

// TestPencilMarkSolver tests the PencilMark solver
func TestPencilMarkSolver(t *testing.T) {
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
		{9, 1, 2, 3, 4, 5, 6, 7, 9},
	})

	solvedBoard, _ := PencilMarkSolver(b)

	if !solvedBoard.Solved() {
		t.Fatalf("Board should be solved")
	}
}

// TestPencilMarkSolverMedium tests the PencilMark solver
func TestPencilMarkSolverMedium(t *testing.T) {
	b := board.Board{}
	b.New([][]int{
		{6, 0, 5, 1, 3, 0, 0, 0, 0},
		{0, 3, 0, 0, 9, 0, 0, 2, 0},
		{0, 0, 0, 0, 0, 4, 0, 8, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 4, 2, 0, 1, 0, 0, 7, 0},
		{0, 0, 7, 6, 8, 0, 0, 0, 5},
		{9, 0, 0, 8, 7, 0, 3, 5, 0},
		{8, 0, 1, 3, 0, 0, 0, 0, 0},
		{0, 7, 3, 2, 0, 6, 0, 0, 9},
	})

	solvedBoard, _ := PencilMarkSolver(b)

	if !solvedBoard.Solved() {
		t.Fatalf("Board should be solved")
	}
}

// TestRemoveValueFromSlice tests the removeValueFromSlice function
//func TestRemoveValueFromSlice(t *testing.T) {
//	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	slice = removeValueFromSlice(slice, 1)
//	if len(slice) != 8 {
//		t.Fatalf("Slice should have 8 elements")
//	}
//	for _, v := range slice {
//		if v == 1 {
//			t.Fatalf("Slice should not contain 1")
//		}
//	}
//
//	slice = removeValueFromSlice(slice, 9)
//	if len(slice) != 7 {
//		t.Fatalf("Slice should have 7 elements")
//	}
//	for _, v := range slice {
//		if v == 9 {
//			t.Fatalf("Slice should not contain 1")
//		}
//	}
//	slice = removeValueFromSlice(slice, 5)
//	if len(slice) != 6 {
//		t.Fatalf("Slice should have 6 elements")
//	}
//	for _, v := range slice {
//		if v == 5 {
//			t.Fatalf("Slice should not contain 1")
//		}
//	}
//}

// TestFindCenter tests the findCenter function
func TestFindCenter(t *testing.T) {
	var cx, cy int
	for q := 0; q < 3; q++ {
		for p := 0; p < 3; p++ {
			for i := q * 3; i < q*3+3; i++ {
				for j := p * 3; j < p*3+3; j++ {
					cx, cy = findCenter(i, j)
					if cx != q*3+1 || cy != p*3+1 {
						t.Errorf("Center for %d,%d should be %d, %d", i, j, q*3+1, p*3+1)
					}
				}
			}

		}
	}
}

func BenchmarkPencilMarkSolver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := getRandomSudoku()

		PencilMarkSolver(s)
	}
}
