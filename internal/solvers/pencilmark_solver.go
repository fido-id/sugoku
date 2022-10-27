package solvers

import (
	"fmt"
	"github.com/mauroartizzu/sugoku/internal/board"
	"sync"
)

type Allowed struct {
	values map[int]map[int][]int
}

func printAllowed(allowed Allowed) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d,%d,%v\n", i, j, allowed.values[i][j])
		}
	}
}

// findCenter returns the coordinates of the center of the quadrant containing the cell at (x,y)
func findCenter(x, y int) (int, int) {
	return x/3*3 + 1, y/3*3 + 1
}

func addAllowedValues(wg *sync.WaitGroup, mux *sync.Mutex, allowed *Allowed, i, j, k int) {
	defer wg.Done()
	mux.Lock()
	allowed.values[i][j] = append(allowed.values[i][j], k)
	mux.Unlock()
}

func PencilMarkSolver(b board.Board) (board.Board, error) {

	// inizializza i valori possibili per ogni cella
	allowed := Allowed{
		values: map[int]map[int][]int{},
	}
	allowed.values = map[int]map[int][]int{}
	for i := 0; i < 9; i++ {
		allowed.values[i] = map[int][]int{}
		for j := 0; j < 9; j++ {
			allowed.values[i][j] = make([]int, 0)
		}
	}

	// aggiunge i valori possibili per ogni cella
	var wg sync.WaitGroup
	mux := sync.Mutex{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.Get(i, j) == 0 {

				// single wg for each cell
				for k := 1; k <= 9; k++ {
					if b.IsAllowedValue(k, i, j) {
						wg.Add(1)
						go addAllowedValues(&wg, &mux, &allowed, i, j, k)
					}
				}

				// procedural with struct
				//for k := 1; k <= 9; k++ {
				//	if b.IsAllowedValue(k, i, j) {
				//		allowed.values[i][j] = append(allowed.values[i][j], k)
				//	}
				//}

				// procedural without struct
				//alloweds := make([]int, 0)
				//for k := 1; k <= 9; k++ {
				//	if b.IsAllowedValue(k, i, j) {
				//		alloweds = append(alloweds, k)
				//	}
				//}
			}
		}
	}
	wg.Wait()
	printAllowed(allowed)

	return b, nil
}

// removeValueFromSlice removes the value from the slice, foreach key in slice if value == key remove key
//func removeValueFromSlice(s []int, i int) []int {
//	for k, v := range s {
//		if v == i {
//			// exchange the value to remove with the last value
//			s[len(s)-1], s[k] = s[k], s[len(s)-1]
//
//			//s[len(s)-1] = 0
//			//return s
//
//			// FIXME this: causes a memory leak
//			return s[:len(s)-1]
//		}
//	}
//	return s
//}

//func iterateTable(b *board.Board, allowed *Allowed) {
//	for x := 0; x < 9; x++ {
//		for y := 0; y < 9; y++ {
//			// if cell is empty
//			if b.Get(x, y) == 0 {
//				fmt.Println("x:", x, "y:", y)
//				// for each possible value
//				for i := 1; i <= 9; i++ {
//					//	// if value is allowed
//					if b.IsAllowedValue(i, x, y) {
//						fmt.Println("i:", i, "is allowed")
//						// add value to allowed values
//						allowed.values[x][y] = append(allowed.values[x][y], i)
//					}
//				}
//				if len(allowed.values[x][y]) == 1 {
//					b.Set(x, y, allowed.values[x][y][0])
//
//					// remove value from row and column
//					for i := 0; i < 9; i++ {
//						allowed.values[x][i] = removeValueFromSlice(allowed.values[x][i], allowed.values[x][y][0])
//						allowed.values[i][y] = removeValueFromSlice(allowed.values[i][y], allowed.values[x][y][0])
//					}
//
//					// remove value from quadrant
//					cx, cy := findCenter(x, y)
//					for i := -1; i <= 1; i++ {
//						allowed.values[cx+i][cy-i] = removeValueFromSlice(allowed.values[cx+i][cy-i], allowed.values[x][y][0]) // quadrant
//					}
//
//					allowed.values[x][y] = make([]int, 9)
//				}
//				go iterateTable(b, allowed)
//			}
//		}
//	}
//}
