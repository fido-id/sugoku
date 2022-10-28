package solvers

import (
	"github.com/mauroartizzu/sugoku/internal/board"
	"sync"
)

var allowedMap map[int]map[int][]int
var mux sync.Mutex
var wg sync.WaitGroup

func PencilMarkSolver(b board.Board) (board.Board, error) {
	allowedMap = map[int]map[int][]int{}

	for i := 0; i < 9; i++ {
		allowedMap[i] = map[int][]int{}
		for j := 0; j < 9; j++ {
			allowedMap[i][j] = make([]int, 0)
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// scansiono tutte le celle e cerco la prima con valore zero (non scritta)
			if b.Get(i, j) == 0 {
				wg.Add(1)
				//parte una routine
				go markAllowed(b, i, j)
			}
		}
	}

	wg.Wait()

	for i := 0; i < 9; i++ {
		if len(allowedMap[i]) == 0 {
			continue
		}

		for j := 0; j < 9; j++ {
			if len(allowedMap[i][j]) == 0 {
				continue
			}

			jojo(b, i, j)
		}
	}

	return b, nil
}

func markAllowed(b board.Board, rowIndex int, colIndex int) {
	defer wg.Done()
	for i := 1; i <= 9; i++ {
		if b.IsAllowedValue(i, rowIndex, colIndex) {
			mux.Lock()
			allowedMap[rowIndex][colIndex] = append(allowedMap[rowIndex][colIndex], i)
			mux.Unlock()
		}
	}
}

func iona(v, r, c int) bool {
	//quante volte compare v in allowedMap[i][c]?
	counter := 0
	for d := range allowedMap[r][c] {
		if d == v {
			counter++
		}

		if counter > 1 {
			return false
		}
	}

	return true
}

func jojo(b board.Board, i, j int) {
	for _, v := range allowedMap[i][j] {
		for c := 0; c < 9; c++ {
			//check row
			if len(allowedMap[i][c]) == 0 {
				continue
			}
			if iona(v, i, c) {
				b.Set(i, j, v)
				return

			}

			//check col
			if len(allowedMap[c][j]) == 0 {
				continue
			}
			if iona(v, c, j) {
				b.Set(i, j, v)
				return
			}

			// todo check quadrant

		}
	}
}
