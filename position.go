package chessgo

import (
	"fmt"
	"strings"
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

type Castling uint8

func NewCastlingFromFen(castlingfen string) (Castling, error) {
	castling := Castling(0b0000)
	return castling, nil
}

// A Position describing a current state in a chess game
type Position struct {
	board      [64]Piece
	sideToMove Color

	castling uint8

	epSquare        int
	halfMoveClock   int
	fullMoveCounter int
}

// PositionInfo contains an exported struct that abstracts the metadata of the current position
type PositionInfo struct {
	SideToMove      Color
	Castling        uint8
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
		fullMoveCounter: 0,
	}
}

func parsePiecesFen(fenpieces string) ([64]Piece, error) {
	pieces := [64]Piece{}
	rows := strings.Split(fenpieces, "/")
	if len(rows) != 8 {
		return [64]Piece{}, fmt.Errorf("wrong number of rows in fen string pieces: (want:8 got:%d)", len(rows))
	}
	for i, row := range rows {
		currentrow := i * 8
		currentpiece := 0
		for _, ch := range row {
			if ch >= '0' && ch <= '9' {
				emptySpaces := int(ch - '0')
				for range emptySpaces {
					pieces[currentrow+currentpiece] = Empty
					currentpiece++
				}
			} else {
				piece, err := PieceFromRune(ch)
				if err != nil {
					return [64]Piece{}, fmt.Errorf("error parsing piece %d at row %d: %w", currentpiece, i, err)
				}
				if piece == Empty {
					return [64]Piece{}, fmt.Errorf("error parsing piece %d at row %d: empty space instead of number", currentpiece, i)
				}
				pieces[currentrow+currentpiece] = piece
				currentpiece++
			}
		}
		if currentpiece != 8 {
			return [64]Piece{}, fmt.Errorf("incorrect number of pieces in row %d (want:8 got:%d)", i, currentpiece)
		}
	}
	return pieces, nil
}

// NewPositionFromFen loads a fen string and returns a position from it
//
// An example fen string: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
func NewPositionFromFen(fen string) (Position, error) {
	fields := strings.Split(fen, " ")
	if len(fields) != 6 {
		return Position{}, fmt.Errorf("wrong amount of fields in fen string: (want:6 got:%d)", len(fields))
	}
	pieces, err := parsePiecesFen(fields[0])
	if err != nil {
		return Position{}, err
	}

	return Position{
		board: pieces,
		// TODO implement full fen parsing
		sideToMove:      White,
		castling:        0b1111,
		epSquare:        -1,
		halfMoveClock:   0,
		fullMoveCounter: 0,
	}, nil
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
