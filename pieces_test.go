package chessgo

import (
	"testing"
)

// The next tests are probably a little bit too verbose, enum conversions probably are better done with maps

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

func TestCastlingGet(t *testing.T) {
	castling := Castling(0b1001)
	gotRights := castling.Get(WhiteKingSideCastling)
	wantRights := true
	if gotRights != wantRights {
		t.Errorf("checking allowed right got %t, want %t", gotRights, wantRights)
	}
	gotRights = castling.Get(WhiteQueenSideCastling)
	wantRights = false
	if gotRights != wantRights {
		t.Errorf("checking denied right got %t, want %t", gotRights, wantRights)
	}
}
func TestCastlingWithout(t *testing.T) {
	castling := Castling(0b1001).Without(WhiteKingSideCastling)
	gotRights := true
	wantRights := false
	if castling.Get(WhiteQueenSideCastling) != wantRights {
		t.Errorf("checking denied right got %t, want %t", gotRights, wantRights)
	}

}
