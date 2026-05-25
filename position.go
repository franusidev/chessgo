package chessgo

type Piece byte

const (
	Empty Piece = iota

	WhiteKing
	WhiteQueen
	WhiteBishop
	WhiteKnight
	WhiteRook
	WhitePawn

	BlackKing
	BlackQueen
	BlackBishop
	BlackKnight
	BlackRook
	BlackPawn
)

func (p Piece) String() string {
	pstring := map[Piece]string{
		Empty:       " ",
		WhiteKing:   "K",
		WhiteQueen:  "Q",
		WhiteBishop: "B",
		WhiteKnight: "N",
		WhiteRook:   "R",
		WhitePawn:   "P",

		BlackKing:   "k",
		BlackQueen:  "q",
		BlackBishop: "b",
		BlackKnight: "n",
		BlackRook:   "r",
		BlackPawn:   "p",
	}
	return pstring[p]
}

// The Color of the pieces
type Color bool

const (
	White = true
	Black = false
)

func (c Color) String() string {
	if c == White {
		return "White"
	} else {
		return "Black"
	}
}

// A Position describing a current state in a chess game
type Position struct {
	board      [64]Piece
	sideToMove Color
}

// NewStartingPosition creates a new starting position with all pieces in their place
func NewStartingPosition() Position {
	return Position{
		board: [64]Piece{
			BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook,
			BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn,
			WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook,
		},
		sideToMove: White,
	}
}

func NewPositionFromFen(fen string) Position {
	return Position{}
}

// PieceAt returns the piece at a given location
//
// The flat representation of a chessboard assumes that in a starting position:
//   - 0-7 is black back rank
//   - 8-15 is black pawns
//   - 16-47 is empty
//   - 48-55 is white pawns
//   - 56-63 is white back rank
func (p Position) PieceAt(x int) Piece {
	return p.board[x]
}

// SideToMove returns the Color to move next
func (p Position) SideToMove() Color {
	return p.sideToMove
}
