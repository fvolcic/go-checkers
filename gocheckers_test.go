package gocheckers

import (
	"testing"
)

func TestCreateCheckersBoard(t *testing.T) {
	board := NewCheckersBoard()
	if board == nil {
		t.Error("Expected a non-nil board")
	}

	if board.board == nil {
		t.Error("Expected a non-nil board.board")
	}

	// Check the board is the right size
	if len(board.board) != 8 {
		t.Error("Expected board to have 8 rows")
	}

	for i := 0; i < 8; i++ {
		if len(board.board[i]) != 8 {
			t.Error("Expected board to have 8 columns")
		}
	}

	if board.turn != black {
		t.Error("Expected black to go first")
	}

}
