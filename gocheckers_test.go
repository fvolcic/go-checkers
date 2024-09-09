package gocheckers

import (
	"testing"
)

func TestCreateCheckersBoard(t *testing.T) {
	board := NewCheckersBoard()
	if board == nil {
		t.Error("Expected a non-nil board")
		t.FailNow()
	}

	if board.board == nil {
		t.Error("Expected a non-nil board.board")
	}

	// Check the board is the right size
	if len(board.GetBoardData()) != 8 {
		t.Error("Expected board to have 8 rows")
	}

	for i := 0; i < 8; i++ {
		if len(board.GetBoardData()[i]) != 8 {
			t.Error("Expected board to have 8 columns")
		}
	}

	if board.turn != Black {
		t.Error("Expected black to go first")
	}

}

func TestGenerateMovesSimple1(t *testing.T) {

	board := NewCheckersBoard()

	hasJump := board.turnHasJump()

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

	board.turn = White

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

	board := NewCheckersBoardFromPosition(position, Black, make([][]int, 0))

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

	board = NewCheckersBoardFromPosition(position, Black, make([][]int, 0))

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

	board := NewCheckersBoardFromPosition(position, Black, make([][]int, 0))

	moves := board.GenerateMoves()

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

	board = NewCheckersBoardFromPosition(position, Black, make([][]int, 0))

	moves = board.GenerateMoves()

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

func TestSingleOptionMoves(t *testing.T) {

	type boardInfo struct {
		position [][]int
		turn     int
	}

	positions := []boardInfo{
		{[][]int{
			{0, 0, 0, 3, 0, 0, 0, 0},
			{0, 0, 2, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
			Black,
		},

		{[][]int{
			{0, 0, 0, 3, 0, 0, 0, 0},
			{0, 0, 0, 0, 2, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
			Black,
		},

		{[][]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0, 0, 0, 0},
			{4, 0, 0, 0, 0, 0, 0, 0},
		},
			White,
		},

		{[][]int{
			{0, 0, 0, 3, 0, 0, 0, 0},
			{0, 0, 2, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 4, 0, 0, 0, 0, 0},
		},
			White,
		},
	}

	for i := 0; i < len(positions); i++ {
		board := NewCheckersBoardFromPosition(positions[i].position, positions[i].turn, make([][]int, 0))

		moves := board.GenerateMoves()

		if len(moves) != 1 {
			t.Errorf("Expected 1 move, but got %d on board index %d", len(moves), i)
		}

		if len(moves[0]) != 2 {
			t.Errorf("Expected a single move, but got %d on board index %d", len(moves[0]), i)
		}
	}
}

func TestDoubleJumpToKing(t *testing.T) {

	type boardInfo struct {
		position [][]int
		turn     int
	}

	positions := []boardInfo{
		{[][]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 2, 0, 2, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
			Black,
		},

		{[][]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 2, 0, 2, 0, 0, 0},
			{0, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
			Black,
		},

		{[][]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{2, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
			White,
		},

		{[][]int{
			{0, 0, 0, 3, 0, 0, 0, 0},
			{0, 0, 2, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 2, 0, 0, 0},
			{0, 1, 0, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
			White,
		},
	}

	for i := 0; i < len(positions); i++ {
		board := NewCheckersBoardFromPosition(positions[i].position, positions[i].turn, make([][]int, 0))

		moves := board.GenerateMoves()

		if len(moves) != 1 {
			t.Errorf("Expected 1 move, but got %d on board index %d", len(moves), i)
		}

		if len(moves[0]) != 3 {
			t.Errorf("Expected a two moves, but got %d on board index %d", len(moves[0]), i)
		}
	}
}

func TestMakeMove(t *testing.T) {
	position := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0},
	}

	board := NewCheckersBoardFromPosition(position, Black, make([][]int, 0))

	if !board.MakeMove([]int{4, 8}) {
		t.Errorf("Expected sucesssful move")
	}

	result := board.GetBoardData()

	if result[7][0] != Empty {
		t.Error("Expected starting square to be empty")
	}

	if result[6][1] != Black {
		t.Error("Expected next square to be black")
	}

}

func TestPromote(t *testing.T) {
	position := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0},
		{0, 4, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 2, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	board := NewCheckersBoardFromPosition(position, Black, make([][]int, 0))

	success := board.MakeMove([]int{28, 32})

	if !success {
		t.Error("Expected successful move")
	}

	result := board.GetBoardData()

	if result[0][1] != BlackKing {
		t.Errorf("Expected black king, got %d", result[0][1])
	}

	// check the piece count to make sure all is good
	if board.GetPieceCount(White) != 3 {
		t.Errorf("Expected 4 white pieces, got %d", board.GetPieceCount(White))
	}

	if board.GetPieceCount(WhiteKing) != 1 {
		t.Errorf("Expected 1 white king, got %d", board.GetPieceCount(White))
	}

	if board.GetPieceCount(BlackKing) != 1 {
		t.Errorf("Expected 1 black king, got %d", board.GetPieceCount(White))
	}
}

func TestGameIsOver(t *testing.T) {
	position := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 0, 2, 0, 0},
		{0, 0, 2, 0, 2, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	board := NewCheckersBoardFromPosition(position, Black, make([][]int, 0))

	if !board.GameOver() {
		t.Error("Expected game to be over")
	}
}

func TestMakeMoveAndGetMoveHistory(t *testing.T) {
	madeMoves := make([][]int, 0)

	board := NewCheckersBoard()

	for i := 0; i < 5; i++ {
		moves := board.GenerateMoves()

		if len(moves) == 0 {
			break
		}

		move := moves[0]

		madeMoves = append(madeMoves, move)

		board.MakeMove(move)
	}

	if len(madeMoves) != len(board.GetGameMoveHistory()) {
		t.Errorf("Expected %d moves, got %d", len(madeMoves), len(board.GetGameMoveHistory()))
	}

	for i, move := range madeMoves {
		for j, square := range move {
			if board.GetGameMoveHistory()[i][j] != square {
				t.Errorf("Expected %d, got %d", square, board.GetGameMoveHistory()[i][j])
			}
		}
	}

}
