package gocheckers

import (
	"fmt"
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

	hasJump := board.TurnHasJump()

	moves := board.generateMovesForPiece(11, false, hasJump)

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

	position = [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0},
	}

	board = NewCheckersBoardFromPosition(position, black, make([][]int, 0))

	numMoves = len(board.GenerateMoves())

	if numMoves != 2 {
		t.Errorf("Expected 2 moves, got %d", numMoves)
	}
}

func TestGenerateMovesWithKing(t *testing.T) {

	position := [][]int{
		{0, 0, 0, 3, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	board := NewCheckersBoardFromPosition(position, black, make([][]int, 0))

	moves := board.GenerateMoves()

	fmt.Println(moves)

	if len(moves) != 1 {
		t.Errorf("Expected 1 move, got %d", len(moves))
		t.FailNow()
	}

	if len(moves[0]) != 2 {
		t.Errorf("Expected 2 squares in the sequence, got %d", len(moves[0]))
		t.FailNow()
	}

	if moves[0][0] != 31 {
		t.Errorf("Expected move 1 to be 31, got %d", moves[0][0])
	}

	if moves[0][1] != 24 {
		t.Errorf("Expected move 2 to be 24, got %d", moves[0][1])
	}

	position = [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 2, 0, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	board = NewCheckersBoardFromPosition(position, black, make([][]int, 0))

	moves = board.GenerateMoves()

	fmt.Println(moves)

	numMoves := len(moves)

	if numMoves != 1 {
		t.Errorf("Expected 1 move, got %d", numMoves)
	}

	if len(moves[0]) != 3 {
		t.Errorf("Expected a sequence length 3, got %d", len(moves[0]))
		t.FailNow()
	}

	if moves[0][0] != 22 {
		t.Errorf("Expected 22 in position 0, got %d", moves[0][0])
	}

	if moves[0][1] != 31 {
		t.Errorf("Expected 31 in position 1, got %d", moves[0][1])
	}

	if moves[0][2] != 24 {
		t.Errorf("Expected 24 in position 2, got %d", moves[0][2])
	}

}
