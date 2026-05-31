package main

import (
	"bufio"
	"fmt"
	"github.com/franusidev/chessgo"
	"os"
	"strings"
)

func renderPosition(pos chessgo.Position) string {
	var b strings.Builder
	printSeparator := func() {
		b.WriteString("   --- --- --- --- --- --- --- --- \n")
	}
	for i := range 8 {
		printSeparator()
		fmt.Fprintf(&b, "%d ", 8-i)
		for j := range 8 {
			fmt.Fprintf(&b, "| %v ", pos.PieceAt(i*8+j))
		}
		b.WriteString("|\n")
	}
	printSeparator()
	b.WriteString("    A   B   C   D   E   F   G   H  \n\n")
	posinfo := pos.Info()
	fmt.Fprintf(&b, "%v to move\n", posinfo.SideToMove)
	fmt.Fprintf(&b, "Available castling: %v\n", posinfo.Castling)
	fmt.Fprintf(&b, "Moves: %v HalfMoveClock: %v\n", posinfo.FullMoveCounter, posinfo.HalfMoveClock)
	return b.String()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("must provide as argument one fen file to read")
		os.Exit(1)
	}
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := bufio.NewScanner(file)
	if !s.Scan() {
		fmt.Println("File is empty")
		os.Exit(1)
	}
	pos, err := chessgo.NewPositionFromFen(s.Bytes())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(renderPosition(pos))
}
