package board

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	//log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestEmptyBoard(t *testing.T) {
	b := Board{}
	b.New()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.board[i][j] != 0 {
				t.Fatalf("Board not empty")
			}
		}
	}
}

// test correct row
func TestCheckRowAndCol(t *testing.T) {
	b := Board{}
	b.New([][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 4},
		{7, 8, 9, 1, 2, 3, 4, 5, 5},
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{9, 2, 1, 3, 4, 5, 6, 7, 8},
	})
	if !b.checkRow(0) {
		t.Fatalf("Row not correct")
	}
	if b.checkRow(1) {
		t.Fatalf("Row not correct")
	}
	if b.checkRow(2) {
		t.Fatalf("Row not correct")
	}
	if !b.checkCol(0) {
		t.Fatalf("Row not correct")
	}
	if b.checkCol(1) {
		t.Fatalf("Row not correct")
	}
	if b.checkCol(2) {
		t.Fatalf("Row not correct")
	}
}

func TestCheckQuadrant(t *testing.T) {
	b := Board{}
	b.New([][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 0, 9, 1, 0, 3},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
	})
	if !b.checkQuadrant(0, 0) {
		t.Fatalf("Quadrant not correct")
	}
	if b.checkQuadrant(1, 0) {
		t.Fatalf("Quadrant not correct")
	}
	if b.checkQuadrant(2, 0) {
		t.Fatalf("Quadrant not correct")
	}
}

// check set value
func TestSetValue(t *testing.T) {
	b := Board{}
	b.New()
	b.Set(0, 0, 1)
	if b.Get(0, 0) != 1 {
		t.Fatalf("Value not set")
	}
}

func TestIsSolved(t *testing.T) {
	b := Board{}
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
	if !b.Solved() {
		t.Fatalf("Board should be solved")
	}
}

func TestIsValid(t *testing.T) {
	b := Board{}
	b.New([][]int{
		{0, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 4},
		{7, 8, 9, 1, 2, 3, 4, 5, 5},
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{9, 2, 1, 3, 4, 5, 6, 7, 8},
	})

	if !b.IsAllowedValue(1, 0, 0) {
		t.Fatalf("Should be valid")
	}
}

func TestShouldNotBeValidByRow(t *testing.T) {
	b := Board{}
	b.New([][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{5, 6, 7, 8, 0, 1, 2, 3, 4},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	})

	if b.IsAllowedValue(1, 4, 4) {
		t.Fatalf("Should not be valid")
	}
}

func TestShouldNotBeValidByCol(t *testing.T) {
	b := Board{}
	b.New([][]int{
		{0, 0, 0, 0, 5, 0, 0, 0, 0},
		{0, 0, 0, 0, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 0, 0, 6, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 0},
	})

	if b.IsAllowedValue(1, 4, 4) {
		t.Fatalf("Should not be valid")
	}
}

func TestShouldNotBeValidByQuadrant(t *testing.T) {
	b := Board{}
	b.New([][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 5, 6, 7, 0, 0, 0},
		{0, 0, 0, 8, 0, 1, 0, 0, 0},
		{0, 0, 0, 2, 3, 4, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	})

	if b.IsAllowedValue(2, 4, 4) {
		t.Fatalf("Should not be valid")
	}
	if b.IsAllowedValue(4, 4, 4) {
		t.Fatalf("Should not be valid")
	}
	if b.IsAllowedValue(5, 4, 4) {
		t.Fatalf("Should not be valid")
	}
	if b.IsAllowedValue(7, 4, 4) {
		t.Fatalf("Should not be valid")
	}
}

func TestIsEmptyCell(t *testing.T) {

	b := Board{}
	b.New([][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 5, 6, 7, 0, 0, 0},
		{0, 0, 0, 8, 0, 1, 0, 0, 0},
		{0, 0, 0, 2, 3, 4, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	})

	if !b.IsEmptyCell(0, 0) {
		t.Fatalf("Should be empty")
	}

	if b.IsEmptyCell(3, 3) {
		t.Fatalf("Should not be empty")
	}
}
