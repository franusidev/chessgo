package main

import (
	"fmt"
	"github.com/franusidev/chessgo"
)

func displayPosition(pos chessgo.Position) {
	printSeparator := func() {
		fmt.Println("   --- --- --- --- --- --- --- --- ")
	}
	for i := range 8 {
		printSeparator()
		fmt.Printf("%d ", 8-i)
		for j := range 8 {
			fmt.Printf("| %v ", pos.PieceAt(i*8+j))
		}
		fmt.Println("|")
	}
	printSeparator()
	fmt.Println("    A   B   C   D   E   F   G   H  ")
	fmt.Println()
	fmt.Printf("%v to move\n", pos.SideToMove())
}

func main() {
	pos := chessgo.NewStartingPosition()
	displayPosition(pos)
}
