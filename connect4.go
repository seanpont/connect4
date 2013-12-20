package main

import (
	"fmt"
)

func main() {
	fmt.Println("Connect4!")

	b := NewBoard()
	won := false
	for !won {
		fmt.Println(b)
		columnIndex := 0
		fmt.Scanf("%d", &columnIndex)
		w, err := b.Play(columnIndex)
		if (err != nil) {
			fmt.Println("Invalid move!")
		}
		won = w
	}
	fmt.Println("Congradulations player", b.turn)
}

