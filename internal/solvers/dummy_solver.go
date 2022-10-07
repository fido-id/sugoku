package solvers

import (
	"github.com/mauroartizzu/sugoku/internal/board"
)

func dummy(b *board.Board) bool {
	return false
}

func DummySolver(b board.Board) (board.Board, error) {

	dummy(&b)

	return b, nil
}
