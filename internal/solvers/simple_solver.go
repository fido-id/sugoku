package solvers

import (
	"github.com/mauroartizzu/sugoku/internal/board"
)

func SimpleSolver(b board.Board) (board.Board, error) {

START:
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if b.IsEmptyCell(i, j) {
				allowedValues := make([]int, 0)
				for x := 1; x < 10; x++ {
					if b.IsAllowedValue(x, i, j) {
						allowedValues = append(allowedValues, x)
					}
				}

				if len(allowedValues) == 1 {
					b.Set(i, j, allowedValues[0])
					goto START
				}
			}
		}
	}

	return b, nil
}
