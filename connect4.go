package main

import (
	"fmt"
)

func main() {
	fmt.Println("Connect4!")

	p1 := promptToken(1)
	p2 := promptToken(2)

	b := NewBoard()
	won := false
	for !won {
		printBoard(b, p1, p2)
		columnIndex := 0
		fmt.Scanf("%d", &columnIndex)
		w, err := b.Play(columnIndex)
		if (err != nil) {
			fmt.Println("Invalid move!")
		}
		won = w
	}
	printBoard(b, p1, p2)
	fmt.Printf("Congratulations player %v!\n", b.turn)
}

func promptToken(player int) string {
	fmt.Printf("Player %d: enter your token: ", player)
	var p1 string
	_, err := fmt.Scanln(&p1)
	if (err != nil) {
		panic("Could not read token")
	}
	if len(p1) > 1 {
		fmt.Println("Token must be one character")
		return promptToken(player)
	}
	return p1
}

func printBoard(b *Board, p1 string, p2 string) {
	s := "\n"
	for rowIndex := ROWS-1; rowIndex >= 0; rowIndex-- {
		s += "|"
		for _, val := range(b.board[rowIndex]) {
			s += " "
			if val == byte(1) {
				s += p1
			} else if val == byte(2) {
				s += p2
			} else {
				s += " "
			}
		}
		s += " |\n"
	}
	s += "|---------------|\n| 1 2 3 4 5 6 7 |"
	if b.turn == byte(2) {
		p1 = p2
	}
	s += fmt.Sprintf("\nPlayer %v(%v): ", b.turn, p1)
	fmt.Print(s)
}










