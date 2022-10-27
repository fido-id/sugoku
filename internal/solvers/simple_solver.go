package solvers

import "github.com/mauroartizzu/sugoku/internal/board"

func SimpleSolver(b board.Board) (board.Board, error) {

CYCLE:
	// controllo tutti i valori validi per la cella, se ce n'è uno solo lo inserisco e "ricomincio da capo", altrimenti passo alla prossima
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.Get(j, i) == 0 {
				allowedValues := make([]int, 0)
				for k := 1; k <= 9; k++ {
					if b.IsAllowedValue(k, j, i) {
						allowedValues = append(allowedValues, k)
					}
				}
				if len(allowedValues) == 1 {
					b.Set(j, i, allowedValues[0])
					goto CYCLE
				}
			}
		}
	}
	return b, nil
}
