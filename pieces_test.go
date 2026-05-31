package chessgo

import (
	"testing"
)

// The next tests are probably a little bit too verbose, enum conversions probably are better done with maps

func Test_pieceFromByte(t *testing.T) {
	tests := []struct {
		name      string
		pieceByte byte
		want      Piece
		wantErr   bool
	}{
		{
			name:      "empty",
			pieceByte: ' ',
			want:      Empty,
		},
		{
			name:      "white king",
			pieceByte: 'K',
			want:      WhiteKing,
		},
		{
			name:      "white queen",
			pieceByte: 'Q',
			want:      WhiteQueen,
		},
		{
			name:      "white bishop",
			pieceByte: 'B',
			want:      WhiteBishop,
		},
		{
			name:      "white knight",
			pieceByte: 'N',
			want:      WhiteKnight,
		},
		{
			name:      "white rook",
			pieceByte: 'R',
			want:      WhiteRook,
		},
		{
			name:      "white pawn",
			pieceByte: 'P',
			want:      WhitePawn,
		},
		{
			name:      "black king",
			pieceByte: 'k',
			want:      BlackKing,
		},
		{
			name:      "black queen",
			pieceByte: 'q',
			want:      BlackQueen,
		},
		{
			name:      "black bishop",
			pieceByte: 'b',
			want:      BlackBishop,
		},
		{
			name:      "black knight",
			pieceByte: 'n',
			want:      BlackKnight,
		},
		{
			name:      "black rook",
			pieceByte: 'r',
			want:      BlackRook,
		},
		{
			name:      "black pawn",
			pieceByte: 'p',
			want:      BlackPawn,
		},
		{
			name:      "invalid piece",
			pieceByte: 'x',
			want:      Empty,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pieceFromByte(tt.pieceByte)
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
		{
			name:  "invalid piece",
			piece: Piece(211),
			want:  "?",
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
