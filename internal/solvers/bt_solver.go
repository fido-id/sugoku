package solvers

import (
	"github.com/mauroartizzu/sugoku/internal/board"
)

func backTracking(b *board.Board) bool {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// scansiono tutte le celle e cerco la prima con valore zero (non scritta)
			if b.Get(i, j) != 0 {
				continue
			}

			for v := 1; v < 10; v++ {
				if !b.IsAllowedValue(v, i, j) {
					continue
				}
				b.Set(i, j, v)
				if backTracking(b) {
					return true
				}
			}
			//  se non ci sono valori disponibil1 setto zero
			b.Set(i, j, 0)
			return false
		}
	}

	return true
}

func BackTrackingSolver(b board.Board) (board.Board, error) {

	// scansiono tutte le celle e cerco la prima con valore zero (non scritta)
	// se non ci sono valori disponibile setto zero e torno alla cella precedente
	// scelgo il successivo valore disponibile rispetto a quello della cella
	// lo scrivo nella cella
	// passo alla prossima cella

	backTracking(&b)

	return b, nil
}
