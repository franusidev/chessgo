package chessgo

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

var testingPositions = map[string]Position{
	"starting": {
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
	},
	"enpassant": {
		board: [64]Piece{
			BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook,
			BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, WhitePawn, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			WhitePawn, WhitePawn, WhitePawn, WhitePawn, Empty, WhitePawn, WhitePawn, WhitePawn,
			WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook,
		},
		sideToMove:      Black,
		castling:        0b1111,
		epSquare:        4*8 + 2, // e3
		halfMoveClock:   0,
		fullMoveCounter: 1,
	},
	"empty": {
		board: [64]Piece{
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		},
		sideToMove:      White,
		castling:        0b0000,
		epSquare:        -1,
		halfMoveClock:   0,
		fullMoveCounter: 1,
	},
}

func comparePieces(t *testing.T, want [64]Piece, got [64]Piece) {
	t.Helper()
	for i, wantPiece := range want {
		gotPiece := got[i]
		if gotPiece != wantPiece {
			t.Errorf("square %c%d: want '%v', got '%v'", 'A'+i%8, 8-i/8, wantPiece, gotPiece)
		}
	}
}

func comparePositions(t *testing.T, want Position, got Position) {
	t.Helper()
	wantPieces := want.board
	gotPieces := got.board
	comparePieces(t, wantPieces, gotPieces)
	wantInfo := want.Info()
	gotInfo := got.Info()
	// TODO change cmp.diff to assert
	if diff := cmp.Diff(wantInfo, gotInfo); diff != "" {
		t.Errorf("wrong position info fields (-want +got):\n%s", diff)
	}
}

func TestParsePiecesFen(t *testing.T) {
	testsValid := []struct {
		name       string
		piecesfen  string
		wantPieces [64]Piece
		wantErr    bool
	}{
		{
			name:       "starting position",
			piecesfen:  "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			wantPieces: testingPositions["starting"].board,
		},
		{
			name:       "after first move",
			piecesfen:  "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR",
			wantPieces: testingPositions["enpassant"].board,
		},
		{
			name:       "empty board",
			piecesfen:  "8/8/8/8/8/8/8/8",
			wantPieces: testingPositions["empty"].board,
		},
		{
			name:      "wrong number of rows",
			piecesfen: "8/8/8/8/pppppppp/8/8/8/8",
			wantErr:   true,
		},
		{
			name:      "too many empty pieces in a row",
			piecesfen: "9/8/8/8/8/8/8/8",
			wantErr:   true,
		},
		{
			name:      "too many pieces in a row",
			piecesfen: "ppppppppp/8/8/8/8/8/8/8",
			wantErr:   true,
		},
		{
			name:      "too many mixed pieces and empty",
			piecesfen: "pp6p/8/8/8/8/8/8/8",
			wantErr:   true,
		},
		{
			name:      "too little empty pieces in a row",
			piecesfen: "9/8/8/8/8/8/8/8",
			wantErr:   true,
		},
		{
			name:      "too little pieces in a row",
			piecesfen: "pppppp/8/8/8/8/8/8/8",
			wantErr:   true,
		},
		{
			name:      "too little mixed pieces and empty",
			piecesfen: "pp2p/8/8/8/8/8/8/8",
			wantErr:   true,
		},
		{
			name:      "made up piece id",
			piecesfen: "ppppxppp/8/8/8/8/8/8/8",
			wantErr:   true,
		},
		{
			name:      "space instead of number for empty",
			piecesfen: "pppp ppp/3 4/8/8/8/8/8/8",
			wantErr:   true,
		},
	}
	for _, tt := range testsValid {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := parsePiecesFen(tt.piecesfen)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("parsePiecesFen did not return an error")
				}
			} else {
				if err != nil {
					t.Fatalf("parsePiecesFen returned error: %s", err)
				}
			}
			comparePieces(t, tt.wantPieces, got)
		})
	}
}

func TestNewCastlingFromFen(t *testing.T) {
	testsValid := []struct {
		name        string
		castlingfen string
		want        Castling
		wantErr     bool
	}{
		{
			name:        "all available castling options",
			castlingfen: "KQkq",
			want:        Castling(0b1111),
		},
		{
			name:        "only a couple of castling options",
			castlingfen: "Kk",
			want:        Castling(0b1010),
		},
		{
			name:        "no castling options",
			castlingfen: "-",
			want:        Castling(0b0000),
		},
		{
			name:        "invalid character",
			castlingfen: "Kqf",
			wantErr:     true,
		},
		{
			name:        "empty string",
			castlingfen: "",
			wantErr:     true,
		},
		{
			name:        "too many characters",
			castlingfen: "KKQQqk",
			wantErr:     true,
		},
	}
	for _, tt := range testsValid {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := parseCastlingFen(tt.castlingfen)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("should return an error")
				}
			} else {
				if err != nil {
					t.Fatalf("should not return error: %s", err)
				}
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("wrong castling parsing (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_parseEnPassantFen(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		square  string
		want    int
		wantErr bool
	}{
		{
			name:   "no en passant",
			square: "-",
			want:   -1,
		},
		{
			name:   "the first square",
			square: "a1",
			want:   0,
		},
		{
			name:   "square in the middle",
			square: "e5",
			want:   4*8 + 4,
		},
		{
			name:   "the last square",
			square: "h8",
			want:   63,
		},
		{
			name:    "too many characters",
			square:  "h8a2",
			wantErr: true,
		},
		{
			name:    "too little characters",
			square:  "h",
			wantErr: true,
		},
		{
			name:    "out of range",
			square:  "q9",
			wantErr: true,
		},
		{
			name:    "incorrect order",
			square:  "4e",
			wantErr: true,
		},
		{
			name:    "empty",
			square:  "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := parseEnPassantFen(tt.square)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("parseEnPassantFen() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("parseEnPassantFen() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("parseEnPassantFen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPositionFromFen(t *testing.T) {
	testsValid := []struct {
		name    string
		fen     string
		want    Position
		wantErr bool
	}{
		{
			name: "starting position",
			fen:  "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			want: testingPositions["starting"],
		},
		{
			name: "en passant",
			fen:  "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
			want: testingPositions["enpassant"],
		},
		{
			name: "no queen castling right",
			fen:  "8/8/8/8/8/8/8/8 b Kkq - 0 0",
			want: Position{
				castling: Castling(0b1011),
				epSquare: -1,
			},
		},
		{
			name: "halfMove and fullMove numbers",
			fen:  "8/8/8/8/8/8/8/8 b - - 12 89",
			want: Position{
				halfMoveClock:   12,
				fullMoveCounter: 89,
				epSquare:        -1,
			},
		},
	}
	for _, tt := range testsValid {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPositionFromFen(tt.fen)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("NewPositionFromFen did not return an error")
				}
			} else {
				if err != nil {
					t.Fatalf("NewPositionFromFen returned error: %s", err)
				}
			}
			comparePositions(t, tt.want, got)
		})
	}
}
