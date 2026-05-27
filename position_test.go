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
		fullMoveCounter: 0,
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
		fullMoveCounter: 0,
	},
}

func comparePieces(t *testing.T, want [64]Piece, got [64]Piece) {
	for i, wantPiece := range want {
		gotPiece := got[i]
		if gotPiece != wantPiece {
			t.Errorf("square %c%d: want '%v', got '%v'", 'A'+i/8, i%8+1, wantPiece, gotPiece)
		}
	}
}
func comparePositions(t *testing.T, want Position, got Position) {
	wantPieces := want.board
	gotPieces := got.board
	comparePieces(t, wantPieces, gotPieces)
	wantInfo := want.Info()
	gotInfo := got.Info()
	if diff := cmp.Diff(wantInfo, gotInfo); diff != "" {
		t.Errorf("wrong info for starting position (-want +got):\n%s", diff)
	}
}

func TestNewStartingPosition(t *testing.T) {
	expected := testingPositions["starting"]
	pos := NewStartingPosition()
	comparePositions(t, expected, pos)
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

func TestNewPositionFromFen(t *testing.T) {
	testsValid := []struct {
		name    string
		fen     string
		want    Position
		wantErr bool
	}{
		{
			name:    "starting position",
			fen:     "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			want:    testingPositions["starting"],
			wantErr: false,
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

func TestPieceFromRune(t *testing.T) {
	tests := []struct {
		name      string
		pieceRune rune
		want      Piece
		wantErr   bool
	}{
		{
			name:      "empty",
			pieceRune: ' ',
			want:      Empty,
		},
		{
			name:      "white king",
			pieceRune: 'K',
			want:      WhiteKing,
		},
		{
			name:      "white queen",
			pieceRune: 'Q',
			want:      WhiteQueen,
		},
		{
			name:      "white bishop",
			pieceRune: 'B',
			want:      WhiteBishop,
		},
		{
			name:      "white knight",
			pieceRune: 'N',
			want:      WhiteKnight,
		},
		{
			name:      "white rook",
			pieceRune: 'R',
			want:      WhiteRook,
		},
		{
			name:      "white pawn",
			pieceRune: 'P',
			want:      WhitePawn,
		},
		{
			name:      "black king",
			pieceRune: 'k',
			want:      BlackKing,
		},
		{
			name:      "black queen",
			pieceRune: 'q',
			want:      BlackQueen,
		},
		{
			name:      "black bishop",
			pieceRune: 'b',
			want:      BlackBishop,
		},
		{
			name:      "black knight",
			pieceRune: 'n',
			want:      BlackKnight,
		},
		{
			name:      "black rook",
			pieceRune: 'r',
			want:      BlackRook,
		},
		{
			name:      "black pawn",
			pieceRune: 'p',
			want:      BlackPawn,
		},
		{
			name:      "invalid piece",
			pieceRune: 'x',
			want:      Empty,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PieceFromRune(tt.pieceRune)
			if (err != nil) != tt.wantErr {
				t.Fatalf("error state mismatch: got err=%v, wantErr=%v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}

}
func TestPieceString(t *testing.T) {
	tests := []struct {
		name  string
		piece Piece
		want  string
	}{
		{
			name:  "empty",
			piece: Empty,
			want:  " ",
		},
		{
			name:  "white king",
			piece: WhiteKing,
			want:  "K",
		},
		{
			name:  "white queen",
			piece: WhiteQueen,
			want:  "Q",
		},
		{
			name:  "white bishop",
			piece: WhiteBishop,
			want:  "B",
		},
		{
			name:  "white knight",
			piece: WhiteKnight,
			want:  "N",
		},
		{
			name:  "white rook",
			piece: WhiteRook,
			want:  "R",
		},
		{
			name:  "white pawn",
			piece: WhitePawn,
			want:  "P",
		},
		{
			name:  "black king",
			piece: BlackKing,
			want:  "k",
		},
		{
			name:  "black queen",
			piece: BlackQueen,
			want:  "q",
		},
		{
			name:  "black bishop",
			piece: BlackBishop,
			want:  "b",
		},
		{
			name:  "black knight",
			piece: BlackKnight,
			want:  "n",
		},
		{
			name:  "black rook",
			piece: BlackRook,
			want:  "r",
		},
		{
			name:  "blackpawn",
			piece: BlackPawn,
			want:  "p",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.piece.String()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}

}
func TestColorString(t *testing.T) {
	tests := []struct {
		name  string
		color Color
		want  string
	}{
		{
			name:  "black",
			color: Black,
			want:  "Black",
		},
		{
			name:  "white",
			color: White,
			want:  "White",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.color.String()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}

}
