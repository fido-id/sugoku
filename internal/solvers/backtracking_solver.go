package solvers

import (
	"github.com/mauroartizzu/sugoku/internal/board"
)

func backTrack(b *board.Board) bool {
	// Ciclo la board
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			// Se la cella è a zero
			if b.Get(x, y) == 0 {
				// Provo tutti i numeri validi
				for i := 1; i <= 9; i++ {
					if b.IsAllowedValue(i, x, y) {
						// Se valido setto il numero, e richiamo ricorsivamente il backTrack
						b.Set(x, y, i)
						if backTrack(b) {
							return true
						}
						// Altrimenti riporto il valore a zero
						b.Set(x, y, 0)
					}
				}
				return false
			}
		}
	}
	return true
}
func BackTrackingSolver(b board.Board) (board.Board, error) {

	backTrack(&b)

	return b, nil
}
