package gocheckers

import (
	"math"
)

// define an enum type

// Define the different square types
const (
	Invalid   = -1
	Empty     = 0
	Black     = 1
	White     = 2
	BlackKing = 3
	WhiteKing = 4
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

	moves [][]int // all of the moves that have occured up to now
	turn  int     // black goes first
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
		make([][]int, 0),
		Black,
	}
}

func NewCheckersBoardFromPosition(position [][]int, turn int, moves [][]int) *CheckersBoard {

	paddedBoard := make([][]int, 10)

	paddedBoard[0] = []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	paddedBoard[9] = []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

	for row := 0; row < 8; row++ {
		paddedBoard[row+1] = make([]int, 10)
		for col := 0; col < 8; col++ {
			paddedBoard[row+1][0] = -1
			paddedBoard[row+1][9] = -1
			paddedBoard[row+1][col+1] = position[row][col]
		}
	}

	return &CheckersBoard{
		paddedBoard,
		moves,
		turn,
	}

}

func (b *CheckersBoard) GenerateDeepCopy() *CheckersBoard {
	boardCopy := make([][]int, 10)

	for i := 0; i < 10; i++ {
		boardCopy[i] = make([]int, 10)
		copy(boardCopy[i], b.board[i])
	}

	return &CheckersBoard{
		boardCopy,
		make([][]int, 0),
		b.turn,
	}
}

func (b *CheckersBoard) GetTurn() int {
	return b.turn
}

// Returns a 2D array of the checkers board
func (b *CheckersBoard) GetBoardData() [][]int {
	unpaddedBoard := make([][]int, 8)

	for row := 0; row < 8; row++ {
		unpaddedBoard[row] = make([]int, 8)
		for col := 0; col < 8; col++ {
			unpaddedBoard[row][col] = b.board[row+1][col+1]
		}
	}

	return unpaddedBoard
}

// Return number of pieces on the board that have the given value
// Note: Non-King and King pieces are counted separetly. For example, if
// there are 2 black pieces and 1 black king, then GetPieceCount(Black) will return 2
func (b *CheckersBoard) GetPieceCount(piece int) int {
	count := 0

	for row := 1; row < 9; row++ {
		for col := 1; col < 9; col++ {
			if b.board[row][col] == piece {
				count++
			}
		}
	}

	return count
}

// Returns true if the game is over. A game is over if there are no possible moves
func (b *CheckersBoard) GameOver() bool {
	return len(b.GenerateMoves()) == 0
}

func (b *CheckersBoard) pieceHasJump(piece int) bool {
	row, col := pieceToIndex(piece)

	pieceColor := Black
	opponentColor := White
	opponentKingColor := WhiteKing
	isKing := false

	if b.board[row][col] == WhiteKing || b.board[row][col] == BlackKing {
		isKing = true
	}

	if b.board[row][col] == White || b.board[row][col] == WhiteKing {
		pieceColor = White
		opponentColor = Black
		opponentKingColor = BlackKing
	}

	// check for jumps

	// up and left | piece color black or king
	if pieceColor == Black || isKing {
		pieceOnSquare := b.board[row-1][col-1]

		if (pieceOnSquare == opponentColor || pieceOnSquare == opponentKingColor) && b.board[row-2][col-2] == Empty {
			return true
		}
	}

	// up and right | piece color black or king
	if pieceColor == Black || isKing {
		pieceOnSquare := b.board[row-1][col+1]

		if (pieceOnSquare == opponentColor || pieceOnSquare == opponentKingColor) && b.board[row-2][col+2] == Empty {
			return true
		}
	}

	// down and left | piece color black or king
	if pieceColor == White || isKing {
		pieceOnSquare := b.board[row+1][col-1]

		if (pieceOnSquare == opponentColor || pieceOnSquare == opponentKingColor) && b.board[row+2][col-2] == Empty {
			return true
		}
	}

	// down and right | piece color black or king
	if pieceColor == White || isKing {
		pieceOnSquare := b.board[row+1][col+1]

		if (pieceOnSquare == opponentColor || pieceOnSquare == opponentKingColor) && b.board[row+2][col+2] == Empty {
			return true
		}
	}

	return false

}

func (b *CheckersBoard) turnHasJump() bool {
	turnColor := Black
	turnKingColor := BlackKing

	if b.turn == White {
		turnColor = White
		turnKingColor = WhiteKing
	}

	for row := 1; row < 9; row++ {
		for col := 1; col < 9; col++ {
			pieceOnSquare := b.board[row][col]
			if pieceOnSquare == turnColor || pieceOnSquare == turnKingColor {
				piece := squareNumbers[row][col]

				if b.pieceHasJump(piece) {
					return true
				}
			}
		}
	}

	return false
}

func (b *CheckersBoard) generateMovesForKing(piece int, isDoubleJump bool, hasJump bool) [][]int {
	moves := make([][]int, 0)

	row, col := pieceToIndex(piece)

	opponentColor := White
	opponentColorKing := WhiteKing

	if b.board[row][col] == White || b.board[row][col] == WhiteKing {
		opponentColor = Black
		opponentColorKing = BlackKing
	}

	// move up and left
	square := b.board[row-1][col-1]
	if square != Invalid && square == Empty && !isDoubleJump && !hasJump {
		squareNumber := squareNumbers[row-1][col-1]
		currentSquare := piece
		move := []int{currentSquare, squareNumber}
		moves = append(moves, move)
	} else if (square == opponentColor || square == opponentColorKing) && b.board[row-2][col-2] == Empty {
		newBoard := b.GenerateDeepCopy()

		newBoard.board[row-2][col-2] = b.board[row][col]
		newBoard.board[row][col] = Empty
		newBoard.board[row-1][col-1] = Empty

		newSquare := squareNumbers[row-2][col-2]

		newBoardMoves := newBoard.generateMovesForKing(newSquare, true, hasJump)

		// If there are no double jumps, then add the single jump to the list
		if len(newBoardMoves) == 0 {
			moves = append(moves, []int{piece, newSquare})
		}

		for moveIndex := 0; moveIndex < len(newBoardMoves); moveIndex++ {
			newBoardMoves[moveIndex] = append([]int{piece}, newBoardMoves[moveIndex]...)
		}

		moves = append(moves, newBoardMoves...)
	}

	// move up and right
	square = b.board[row-1][col+1]
	if square != Invalid && square == Empty && !isDoubleJump && !hasJump {
		squareNumber := squareNumbers[row-1][col+1]
		currentSquare := piece
		move := []int{currentSquare, squareNumber}
		moves = append(moves, move)
	} else if (square == opponentColor || square == opponentColorKing) && b.board[row-2][col+2] == Empty {
		newBoard := b.GenerateDeepCopy()

		newBoard.board[row-2][col+2] = b.board[row][col]
		newBoard.board[row][col] = Empty
		newBoard.board[row-1][col+1] = Empty

		newSquare := squareNumbers[row-2][col+2]

		newBoardMoves := newBoard.generateMovesForKing(newSquare, true, hasJump)

		// If there are no double jumps, then add the single jump to the list
		if len(newBoardMoves) == 0 {
			moves = append(moves, []int{piece, newSquare})
		}

		for moveIndex := 0; moveIndex < len(newBoardMoves); moveIndex++ {
			newBoardMoves[moveIndex] = append([]int{piece}, newBoardMoves[moveIndex]...)
		}

		moves = append(moves, newBoardMoves...)
	}

	// move down and left
	square = b.board[row+1][col-1]
	if square != Invalid && square == Empty && !isDoubleJump && !hasJump {
		squareNumber := squareNumbers[row+1][col-1]
		currentSquare := piece
		move := []int{currentSquare, squareNumber}
		moves = append(moves, move)
	} else if (square == opponentColor || square == opponentColorKing) && b.board[row+2][col-2] == Empty {
		newBoard := b.GenerateDeepCopy()

		newBoard.board[row+2][col-2] = b.board[row][col]
		newBoard.board[row][col] = Empty
		newBoard.board[row+1][col-1] = Empty

		newSquare := squareNumbers[row+2][col-2]

		newBoardMoves := newBoard.generateMovesForKing(newSquare, true, hasJump)

		// If there are no double jumps, then add the single jump to the list
		if len(newBoardMoves) == 0 {
			moves = append(moves, []int{piece, newSquare})
		}

		for moveIndex := 0; moveIndex < len(newBoardMoves); moveIndex++ {
			newBoardMoves[moveIndex] = append([]int{piece}, newBoardMoves[moveIndex]...)
		}

		moves = append(moves, newBoardMoves...)
	}

	// move down and right
	square = b.board[row+1][col+1]
	if square != Invalid && square == Empty && !isDoubleJump && !hasJump {
		squareNumber := squareNumbers[row+1][col+1]
		currentSquare := piece
		move := []int{currentSquare, squareNumber}
		moves = append(moves, move)
	} else if (square == opponentColor || square == opponentColorKing) && b.board[row+2][col+2] == Empty {
		newBoard := b.GenerateDeepCopy()

		newBoard.board[row+2][col+2] = b.board[row][col]
		newBoard.board[row][col] = Empty
		newBoard.board[row+1][col+1] = Empty

		newSquare := squareNumbers[row+2][col+2]

		newBoardMoves := newBoard.generateMovesForKing(newSquare, true, hasJump)

		// If there are no double jumps, then add the single jump to the list
		if len(newBoardMoves) == 0 {
			moves = append(moves, []int{piece, newSquare})
		}

		for moveIndex := 0; moveIndex < len(newBoardMoves); moveIndex++ {
			newBoardMoves[moveIndex] = append([]int{piece}, newBoardMoves[moveIndex]...)
		}

		moves = append(moves, newBoardMoves...)
	}

	return moves
}

func (b *CheckersBoard) generateMovesForBlackPiece(piece int, isDoubleJump bool, hasJump bool) [][]int {
	var moves [][]int

	row, col := pieceToIndex(piece)

	if b.board[row][col] == BlackKing {
		return b.generateMovesForKing(piece, isDoubleJump, hasJump)
	}

	// move up and left
	square := b.board[row-1][col-1]
	if square != Invalid {

		if square == Empty && !isDoubleJump && !hasJump {
			squareNumber := squareNumbers[row-1][col-1]
			currentSquare := piece
			move := []int{currentSquare, squareNumber}
			moves = append(moves, move)
		} else if square == White || square == WhiteKing {
			// check if we can jump
			if b.board[row-2][col-2] == Empty {

				newBoard := b.GenerateDeepCopy() // make a copy of the board
				newBoard.board[row-2][col-2] = b.board[row][col]
				newBoard.board[row-1][col-1] = Empty
				newBoard.board[row][col] = Empty
				newBoardSquare := squareNumbers[row-2][col-2]

				// check if we need to king the piece
				if row-2 == 1 {
					newBoard.board[row-2][col-2] = BlackKing
				}

				newBoardMoves := newBoard.generateMovesForBlackPiece(newBoardSquare, true, hasJump)

				// If there are no double jumps, then add the single jump to the list
				if len(newBoardMoves) == 0 {
					moves = append(moves, []int{piece, newBoardSquare})
				}

				// add the current move to new board moves (accounts for the double jump)
				for moveIndex := 0; moveIndex < len(newBoardMoves); moveIndex++ {
					newBoardMoves[moveIndex] = append([]int{piece}, newBoardMoves[moveIndex]...)
				}

				moves = append(moves, newBoardMoves...)
			}
		}

	}

	// move up and right
	square = b.board[row-1][col+1]
	if square != Invalid {

		if square == Empty && !isDoubleJump && !hasJump {
			squareNumber := squareNumbers[row-1][col+1]
			currentSquare := piece
			move := []int{currentSquare, squareNumber}
			moves = append(moves, move)
		} else if square == White || square == WhiteKing {
			// check if we can jump
			if b.board[row-2][col+2] == Empty {

				newBoard := b.GenerateDeepCopy() // make a copy of the board
				newBoard.board[row-2][col+2] = b.board[row][col]
				newBoard.board[row-1][col+1] = Empty
				newBoard.board[row][col] = Empty
				newBoardSquare := squareNumbers[row-2][col+2]

				// check if we need to king the piece
				if row-2 == 1 {
					newBoard.board[row-2][col+2] = BlackKing
				}

				newBoardMoves := newBoard.generateMovesForBlackPiece(newBoardSquare, true, hasJump)

				// If there are no double jumps, then add the single jump to the list
				if len(newBoardMoves) == 0 {
					moves = append(moves, []int{piece, newBoardSquare})
				}

				// add the current move to new board moves (accounts for the double jump)
				for moveIndex := 0; moveIndex < len(newBoardMoves); moveIndex++ {
					newBoardMoves[moveIndex] = append([]int{piece}, newBoardMoves[moveIndex]...)
				}

				moves = append(moves, newBoardMoves...)
			}
		}

	}

	return moves

}

func (b *CheckersBoard) generateMovesForWhitePeice(piece int, isDoubleJump bool, hasJump bool) [][]int {
	var moves [][]int

	row, col := pieceToIndex(piece)

	if b.board[row][col] == WhiteKing {
		return b.generateMovesForKing(piece, isDoubleJump, hasJump)
	}

	// move down and left
	square := b.board[row+1][col-1]
	if square != Invalid {

		if square == Empty && !isDoubleJump && !hasJump {
			squareNumber := squareNumbers[row+1][col-1]
			currentSquare := piece
			move := []int{currentSquare, squareNumber}
			moves = append(moves, move)
		} else if square == Black || square == BlackKing {
			// check if we can jump
			if b.board[row+2][col-2] == Empty {

				newBoard := b.GenerateDeepCopy() // make a copy of the board
				newBoard.board[row+2][col-2] = b.board[row][col]
				newBoard.board[row+1][col-1] = Empty
				newBoard.board[row][col] = Empty
				newBoardSquare := squareNumbers[row+2][col-2]

				// check if we need to king the piece
				if row+2 == 8 {
					newBoard.board[row+2][col-2] = WhiteKing
				}

				newBoardMoves := newBoard.generateMovesForWhitePeice(newBoardSquare, true, hasJump)

				// If there are no double jumps, then add the single jump to the list
				if len(newBoardMoves) == 0 {
					moves = append(moves, []int{piece, newBoardSquare})
				}

				// add the current move to new board moves (accounts for the double jump)
				for moveIndex := 0; moveIndex < len(newBoardMoves); moveIndex++ {
					newBoardMoves[moveIndex] = append([]int{piece}, newBoardMoves[moveIndex]...)
				}

				moves = append(moves, newBoardMoves...)
			}
		}

	}

	// move down and right
	square = b.board[row+1][col+1]
	if square != Invalid {

		if square == Empty && !isDoubleJump && !hasJump {
			squareNumber := squareNumbers[row+1][col+1]
			currentSquare := piece
			move := []int{currentSquare, squareNumber}
			moves = append(moves, move)
		} else if square == Black || square == BlackKing {
			// check if we can jump
			if b.board[row+2][col+2] == Empty {

				newBoard := b.GenerateDeepCopy() // make a copy of the board
				newBoard.board[row+2][col+2] = b.board[row][col]
				newBoard.board[row+1][col+1] = Empty
				newBoard.board[row][col] = Empty
				newBoardSquare := squareNumbers[row+2][col+2]

				// check if we need to king the piece
				if row+2 == 8 {
					newBoard.board[row+2][col+2] = WhiteKing
				}

				newBoardMoves := newBoard.generateMovesForWhitePeice(newBoardSquare, true, hasJump)

				// If there are no double jumps, then add the single jump to the list
				if len(newBoardMoves) == 0 {
					moves = append(moves, []int{piece, newBoardSquare})
				}

				// add the current move to new board moves (accounts for the double jump)
				for moveIndex := 0; moveIndex < len(newBoardMoves); moveIndex++ {
					newBoardMoves[moveIndex] = append([]int{piece}, newBoardMoves[moveIndex]...)
				}

				moves = append(moves, newBoardMoves...)
			}
		}

	}

	return moves

}

func (b *CheckersBoard) generateMovesForPiece(piece int, isDoubleJump bool, hasJump bool) [][]int {

	row, col := pieceToIndex(piece)

	if b.board[row][col] == Black || b.board[row][col] == BlackKing {
		return b.generateMovesForBlackPiece(piece, isDoubleJump, hasJump)
	}

	if b.board[row][col] == White || b.board[row][col] == WhiteKing {
		return b.generateMovesForWhitePeice(piece, isDoubleJump, hasJump)
	}

	return make([][]int, 0)

}

func (b *CheckersBoard) GenerateMoves() [][]int {

	moves := make([][]int, 0)

	hasJump := b.turnHasJump()

	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {

			if (b.board[row][col] == White || b.board[row][col] == WhiteKing) && b.turn == White {
				moves = append(moves, b.generateMovesForPiece(squareNumbers[row][col], false, hasJump)...)
			}

			if (b.board[row][col] == Black || b.board[row][col] == BlackKing) && b.turn == Black {
				moves = append(moves, b.generateMovesForPiece(squareNumbers[row][col], false, hasJump)...)
			}

		}
	}

	return moves

}

func (b *CheckersBoard) MakeMove(move []int) bool {
	moveSuccess := b.makeMoveHelper(move, false)

	if moveSuccess {
		b.moves = append(b.moves, move)
	}

	return moveSuccess
}

func (b *CheckersBoard) makeMoveHelper(move []int, secondJump bool) bool {

	possibleMoves := b.GenerateMoves()

	validMove := false

	for i := 0; i < len(possibleMoves); i++ {
		if len(possibleMoves[i]) != len(move) {
			continue
		}

		for j := 0; j < len(possibleMoves[i]); j++ {
			if possibleMoves[i][j] != move[j] {
				break
			}

			if j == len(possibleMoves[i])-1 {
				validMove = true
			}
		}

		if validMove {
			break
		}
	}

	if !validMove {
		return false
	}

	startPiece := move[0]
	nextPiece := move[1]

	startRow, startCol := pieceToIndex(startPiece)
	nextRow, nextCol := pieceToIndex(nextPiece)

	if math.Abs(float64(startRow-nextRow)) == 1 {
		b.board[nextRow][nextCol] = b.board[startRow][startCol]
		b.board[startRow][startCol] = Empty

		if nextRow == 1 || nextRow == 8 {
			if b.turn == Black {
				b.board[nextRow][nextCol] = BlackKing
			} else {
				b.board[nextRow][nextCol] = WhiteKing
			}
		}

		if b.turn == Black {
			b.turn = White
		} else {
			b.turn = Black
		}

		return true
	} else {
		b.board[nextRow][nextCol] = b.board[startRow][startCol]
		b.board[startRow][startCol] = Empty

		jumpedRow := (startRow + nextRow) / 2
		jumpedCol := (startCol + nextCol) / 2

		b.board[jumpedRow][jumpedCol] = Empty

		if nextRow == 1 || nextRow == 8 {
			if b.turn == Black {
				b.board[nextRow][nextCol] = BlackKing
			} else {
				b.board[nextRow][nextCol] = WhiteKing
			}
		}

		if len(move) > 2 {
			remainingSteps := make([]int, len(move)-1)
			for i := 1; i < len(move); i++ {
				remainingSteps[i-1] = move[i]
			}

			b.makeMoveHelper(remainingSteps, true)
		}

		if !secondJump {

			if b.turn == Black {
				b.turn = White
			} else {
				b.turn = Black
			}

		}

		return true
	}

}

// Return a copy of the moves that have been made.
func (b *CheckersBoard) GetGameMoveHistory() [][]int {
	return b.moves
}

func (b *CheckersBoard) ToString() string {

	boardStr := ""

	for row := 1; row < 9; row++ {
		for col := 1; col < 9; col++ {
			switch b.board[row][col] {
			case Black:
				boardStr += " b"
			case White:
				boardStr += " w"
			case BlackKing:
				boardStr += " B"
			case WhiteKing:
				boardStr += " W"
			default:
				boardStr += " _"
			}
		}
		boardStr += "\n"
	}

	return boardStr
}
