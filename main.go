package main

import (
	"fmt"
	"github.com/seanpont/connect4/connect4"
)

func main() {
	fmt.Println("Connect4!")
	onePlayer := promptOnePlayer()
	if (onePlayer) {
		onePlayerMode()
		return
	}
	p1 := promptToken(1)
	p2 := promptToken(2)

	g := connect4.NewGame()
	won := false
	for !won {
		printBoard(g, p1, p2)
		columnIndex := 0
		fmt.Scanf("%d", &columnIndex)
		w, err := g.Play(columnIndex)
		if (err != nil) {
			fmt.Println("Invalid move!")
		}
		won = w
	}
	printBoard(g, p1, p2)
	winner := g.Turn() % 2 + 1
	fmt.Printf("Congratulations player %v!\n", winner)
}

func onePlayerMode() {
	p1 := promptToken(1)
	p2 := promptToken(2)
	depth := promptDepth()
	g := connect4.NewGame()
	won := false
	for !won {
		printBoard(g, p1, p2)
		if g.Turn() == byte(1) {
			columnIndex := 0
			fmt.Scanf("%d", &columnIndex)
			w, err := g.Play(columnIndex)
			if (err != nil) {
				fmt.Println("Invalid move!")
			}
			won = w
		} else {
			move := connect4.Solve(g, depth)
			w, err := g.Play(move)
			if (err != nil) {
				fmt.Println("Invalid move!")
			}
			won = w
		}
	}
	printBoard(g, p1, p2)
	winner := g.Turn() % 2 + 1
	fmt.Printf("Congratulations player %v!\n", winner)
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

func promptOnePlayer() bool {
	fmt.Printf("Enter number of players (1 or 2): ")
	var numPlayers int
	_, err := fmt.Scanf("%d", &numPlayers)
	if err != nil || numPlayers < 1 || numPlayers > 2 {
		fmt.Println("Please enter 1 or 2")
		return promptOnePlayer()
	}
	return numPlayers == 1
}

func promptDepth() int {
	fmt.Printf("Depth of ai search: ")
	var depth int
	fmt.Scanf("%d", &depth)
	return depth
}

func printBoard(g *connect4.Game, p1 string, p2 string) {
	fmt.Print("\033[2J\033[1;1H")
	s := ""
	for rowIndex := connect4.ROWS-1; rowIndex >= 0; rowIndex-- {
		s += "|"
		for _, val := range(g.Board()[rowIndex]) {
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
	if g.Turn() == byte(2) {
		p1 = p2
	}
	s += fmt.Sprintf("\nPlayer %v(%v): ", g.Turn(), p1)
	fmt.Print(s)
}










