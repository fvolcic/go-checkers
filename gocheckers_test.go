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
	if len(board.GetUnpaddedBoard()) != 8 {
		t.Error("Expected board to have 8 rows")
	}

	for i := 0; i < 8; i++ {
		if len(board.GetUnpaddedBoard()[i]) != 8 {
			t.Error("Expected board to have 8 columns")
		}
	}

	if board.turn != black {
		t.Error("Expected black to go first")
	}

}

func TestGenerateMovesSimple1(t *testing.T) {

	board := NewCheckersBoard()

	moves := board.generateMovesForPiece(11)

	if len(moves) != 2 {
		t.Errorf("Expected 2 movesm got %d", len(moves))
	}

}

func TestGenerateMovesSimple(t *testing.T) {

	board := NewCheckersBoard()

	moves := board.GenerateMoves()

	if len(moves) != 7 {
		t.Errorf("Expected 7 moves, got %d", len(moves))
	}

	board.turn = white

	moves = board.GenerateMoves()

	if len(moves) != 7 {
		t.Errorf("Expected 7 moves, got %d", len(moves))
	}

}

func TestGenerateMovesFromPosition(t *testing.T) {
	position := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0},
	}

	board := NewCheckersBoardFromPosition(position, black, make([][]int, 0))

	numMoves := len(board.GenerateMoves())

	if numMoves != 1 {
		t.Errorf("Expected 1 move, got %d", numMoves)
	}
}
