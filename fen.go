package chessgo

import (
	"fmt"
	"strconv"
	"strings"
)

// parsePiecesFen takes the field of a fen pieces string as paramater, it returns the arrangement of pieces if its
// correct, else it returns an error
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

// NewCastlingFromFen accepts a string with the castling options in fen notation, oficial notation is strict but this will work even out of order and with repetitions, it is not meant to validate
//
//	'KQkq' => All castling positions available
//	'Kq' => White can castle kingside, black can castle queenside
//	'-' => No castling options available
func parseCastlingFen(castlingfen string) (Castling, error) {
	var zero Castling
	if len(castlingfen) > 4 {
		return zero, fmt.Errorf("too many characters in castlingfen: %d (max. 4)", len(castlingfen))
	}
	if castlingfen == "" {
		return zero, fmt.Errorf("fen castling string is empty")
	}
	if castlingfen == "-" {
		return zero, nil
	}
	castling := Castling(0b0000)
	for _, ch := range castlingfen {
		switch ch {
		case 'K':
			castling |= WhiteKingSideCastling
		case 'Q':
			castling |= WhiteQueenSideCastling
		case 'k':
			castling |= BlackKingSideCastling
		case 'q':
			castling |= BlackQueenSideCastling
		default:
			return zero, fmt.Errorf("invalid character in fen castling string %c", ch)
		}
	}

	return castling, nil
}

func parseEnPassantFen(square string) (int, error) {
	if square == "-" {
		return -1, nil
	}
	if len(square) != 2 {
		return 0, fmt.Errorf("enpassant field must have two characters, got: %s", square)

	}
	row := int(square[0] - 'a')
	if row < 0 || row > 7 {
		return 0, fmt.Errorf("row out of range in square %s", square)
	}
	column := int(square[1] - '1')
	if column < 0 || column > 7 {
		return 0, fmt.Errorf("column out of range in square %s", square)
	}
	return 8*row + column, nil
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
		return Position{}, fmt.Errorf("in parsePiecesFen: %w", err)
	}

	var color Color
	switch fields[1] {
	case "w":
		color = White
	case "b":
		color = Black
	default:
		return Position{}, fmt.Errorf("invalid current color field: %s", fields[1])
	}

	castling, err := parseCastlingFen(fields[2])
	if err != nil {
		return Position{}, fmt.Errorf("in parseCastlingFen: %w", err)
	}

	epsquare, err := parseEnPassantFen(fields[3])
	if err != nil {
		return Position{}, fmt.Errorf("in parseEnPassantFen: %w", err)
	}
	halfmove, err := strconv.Atoi(fields[4])
	if err != nil {
		return Position{}, fmt.Errorf("when parsing half move clock: %w", err)
	}
	fullmove, err := strconv.Atoi(fields[5])
	if err != nil {
		return Position{}, fmt.Errorf("when parsing full move counter: %w", err)
	}

	return Position{
		board:           pieces,
		sideToMove:      color,
		castling:        castling,
		epSquare:        epsquare,
		halfMoveClock:   halfmove,
		fullMoveCounter: fullmove,
	}, nil
}
