package chessgo

import (
	"fmt"
)

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

func pieceFromByte(r byte) (Piece, error) {
	switch r {
	case ' ':
		return Empty, nil
	case 'K':
		return WhiteKing, nil
	case 'Q':
		return WhiteQueen, nil
	case 'B':
		return WhiteBishop, nil
	case 'N':
		return WhiteKnight, nil
	case 'R':
		return WhiteRook, nil
	case 'P':
		return WhitePawn, nil
	case 'k':
		return BlackKing, nil
	case 'q':
		return BlackQueen, nil
	case 'b':
		return BlackBishop, nil
	case 'n':
		return BlackKnight, nil
	case 'r':
		return BlackRook, nil
	case 'p':
		return BlackPawn, nil
	default:
		return Empty, fmt.Errorf("%q is not a valid piece", r)
	}
}

func (p Piece) String() string {
	switch p {
	case Empty:
		return " "
	case WhiteKing:
		return "K"
	case WhiteQueen:
		return "Q"
	case WhiteBishop:
		return "B"
	case WhiteKnight:
		return "N"
	case WhiteRook:
		return "R"
	case WhitePawn:
		return "P"

	case BlackKing:
		return "k"
	case BlackQueen:
		return "q"
	case BlackBishop:
		return "b"
	case BlackKnight:
		return "n"
	case BlackRook:
		return "r"
	case BlackPawn:
		return "p"
	default:
		return "?"
	}
}

// The Color of the pieces
type Color bool

const (
	White Color = true
	Black Color = false
)

func (c Color) String() string {
	if c == White {
		return "White"
	} else {
		return "Black"
	}
}

// Castling contains the information of available castling options in the position
type Castling uint8

const (
	WhiteKingSideCastling  Castling = 0b1000
	WhiteQueenSideCastling Castling = 0b0100
	BlackKingSideCastling  Castling = 0b0010
	BlackQueenSideCastling Castling = 0b0001
)

// Get reports whether the castling right is set for that CastlingSide.
func (c Castling) Get(right Castling) bool {
	return c&right != 0
}

// Without allows to disable a castling right.
func (c Castling) Without(right Castling) Castling {
	return c &^ right
}
