package chessgo

import (
	"testing"
)

func TestNewStartingPosition(t *testing.T) {
	expected := [64]Piece{
		BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook,
		BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn,
		WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook,
	}
	pos := NewStartingPosition()
	for i, want := range expected {
		got := pos.PieceAt(i)
		if got != want {
			t.Errorf("square %c%d: want %v, got %v", 'A'+i/8, i%8+1, want, got)
		}
	}

	if pos.SideToMove() != White {
		t.Errorf("expected White to move")
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
