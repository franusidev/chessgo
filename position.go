package chessgo

// A Position describing a current state in a chess game.
// Currently uses mailbox implementation, probably will change to bitboard eventually
type Position struct {
	board      [64]Piece
	sideToMove Color

	castling Castling

	epSquare        int
	halfMoveClock   int
	fullMoveCounter int
}

// PositionInfo contains an exported struct that abstracts the metadata of the current position
type PositionInfo struct {
	SideToMove      Color
	Castling        Castling
	EpSquare        int
	HalfMoveClock   int
	FullMoveCounter int
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
		sideToMove:      White,
		castling:        0b1111,
		epSquare:        -1,
		halfMoveClock:   0,
		fullMoveCounter: 1,
	}
}

// Info returns a view of the current state of the position aside from the board
func (p Position) Info() PositionInfo {
	return PositionInfo{
		SideToMove:      p.sideToMove,
		Castling:        p.castling,
		EpSquare:        p.epSquare,
		HalfMoveClock:   p.halfMoveClock,
		FullMoveCounter: p.fullMoveCounter,
	}
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
