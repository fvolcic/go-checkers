package gocheckers

// define an enum type

// Define the different square types
const (
	invalid  = -1
	empty     = 0
	black     = 1
	white     = 2
	blackKing = 3
	whiteKing = 4
)

var (
	squareNumbers = [][]int{
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, 32, -1, 31, -1, 30, -1, 29, -1},
		{-1, 28, -1, 27, -1, 26, -1, 25, -1, -1},
		{-1, -1, 24, -1, 23, -1, 22, -1, 21, -1},
		{-1, 20, -1, 19, -1, 18, -1, 17, -1, -1},
		{-1, -1, 16, -1, 15, -1, 14, -1, 13, -1},
		{-1, 12, -1, 11, -1, 10, -1, 9, -1, -1},
		{-1, -1, 8, -1, 7, -1, 6, -1, 5, -1},
		{-1, 4, -1, 3, -1, 2, -1, 1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	}
)

func squareNumber(row, col int) int {
	return squareNumbers[row][col]
}

func pieceToIndex(piece int) (int, int) {
	for row := 1; row < 9; row++ {
		for col := 1; col < 9; col++ {
			if squareNumbers[row][col] == piece {
				return row, col
			}
		}
	}

	return -1, -1
}

type CheckersBoard struct {
	board [][]int

	turn int // black goes first
}

func NewCheckersBoard() *CheckersBoard {
	board :=
		[][]int{
			{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, 2, -1, 2, -1, 2, -1, 2, -1},
			{-1, 2, -1, 2, -1, 2, -1, 2, -1, -1},
			{-1, -1, 2, -1, 2, -1, 2, -1, 2, -1},
			{-1, 0, -1, 0, -1, 0, -1, 0, -1, -1},
			{-1, -1, 0, -1, 0, -1, 0, -1, 0, -1},
			{-1, 1, -1, 1, -1, 1, -1, 1, -1, -1},
			{-1, -1, 1, -1, 1, -1, 1, -1, 1, -1},
			{-1, 1, -1, 1, -1, 1, -1, 1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		}

	return &CheckersBoard{
		board,
		black,
	}
}

func (b * CheckersBoard) GenerateDeepCopy() *CheckersBoard {
	boardCopy := make([][]int, 10)

	for i := 0; i < 10; i++ {
		boardCopy[i] = make([]int, 10)
		copy(boardCopy[i], b.board[i])
	}

	return &CheckersBoard{
		boardCopy,
		b.turn,
	}
}

func (b *CheckersBoard) GetTurn() int {
	return b.turn
}

func (b * CheckersBoard) GetUnpaddedBoard() [][]int {
	unpaddedBoard := make([][]int, 8)

	for row := 1; row < 9; row++ {
		unpaddedBoard[row] = make([]int, 8)
		for col := 1; col < 9; col++ {
			unpaddedBoard[row][col] = b.board[row+1][col+1]
		}
	}

	return unpaddedBoard
}

func GenerateMovesForKing(piece int) [][]int {

}

func (b * CheckersBoard) GenerateMovesForBlackPiece(piece int) [][]int {
	var moves [][]int

	row, col := pieceToIndex(piece)

	if b.board[row][col] == blackKing {
		return GenerateMovesForKing(piece)
	}

	// move up and left
	square := b.board[row-1][col-1]
	if square != invalid {
		
		if square == empty {
			squareNumber = squareNumbers(row-1, col-1)
			currentSquare = piece
			move = []int{piece, squareNumber}
			moves = append(moves, move)
		} else if square == white || square == whiteKing {
			// check if we can jump
			if b.board[row-2][col-2] == empty {
				
				newBoard := b.GenerateDeepCopy() // make a copy of the board
				newBoard.board[row-2][col-2] = b.board[row][col]
				newBoard.board[row-1][col-1] = empty
				newBoard.board[row][col] = empty
				newBoardSquare = squareNumbers[row-2][col-2]
				
				// check if we need to king the piece
				if row-2 == 1 {
					newBoard.board[row-2][col-2] = blackKing
				}

				newBoardMoves := newBoard.GenerateMovesForBlackPiece(newBoardSquare)

				moves = append(moves, newBoardMoves...)
			}
		}

	}

	// move up and right
	square := b.board[row-1][col+1]
	if square != invalid {
		
		if square == empty {
			squareNumber = squareNumbers(row-1, col+1)
			currentSquare = piece
			move = []int{piece, squareNumber}
			moves = append(moves, move)
		} else if square == white || square == whiteKing {
			// check if we can jump
			if b.board[row-2][col+2] == empty {
				
				newBoard := b.GenerateDeepCopy() // make a copy of the board
				newBoard.board[row-2][col+2] = b.board[row][col]
				newBoard.board[row-1][col+1] = empty
				newBoard.board[row][col] = empty
				newBoardSquare = squareNumbers[row-2][col+2]

				// check if we need to king the piece
				if row-2 == 1 {
					newBoard.board[row-2][col+2] = blackKing
				}

				moves = append(moves, newBoardMoves...)
			}
		}

	}


	return moves

}

func (b *CheckersBoard) GenerateMovesForWhitePeice(piece int) [][]int {
	var moves [][]int

	row, col := pieceToIndex(piece)

	if b.board[row][col] == whiteKing {
		return GenerateMovesForKing(piece)
	}

	// move down and left
	
}

func (b *CheckersBoard) GenerateMovesForPiece(piece int) int[][] {
	var moves int[][]
}
