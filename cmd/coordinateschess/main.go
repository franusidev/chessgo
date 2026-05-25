package main

import (
	"fmt"
	"github.com/franusidev/chessgo"
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
	fmt.Fprintf(&b, "%v to move\n", pos.SideToMove())
	return b.String()
}

func main() {
	pos := chessgo.NewStartingPosition()
	fmt.Println(renderPosition(pos))
}
