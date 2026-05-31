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

func PieceFromRune(r rune) (Piece, error) {
	// TODO change map ineficient implementation to switch
	rpiece := map[rune]Piece{
		' ': Empty,
		'K': WhiteKing,
		'Q': WhiteQueen,
		'B': WhiteBishop,
		'N': WhiteKnight,
		'R': WhiteRook,
		'P': WhitePawn,

		'k': BlackKing,
		'q': BlackQueen,
		'b': BlackBishop,
		'n': BlackKnight,
		'r': BlackRook,
		'p': BlackPawn,
	}
	piece, ok := rpiece[r]
	if !ok {
		return Empty, fmt.Errorf("%q is not a valid piece", r)
	}
	return piece, nil
}

func (p Piece) String() string {
	// TODO change map ineficient implementation to switch
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
