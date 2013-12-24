package connect4

import (
	// "fmt"
	// "github.com/seanpont/connect4/minimax"
)

// func (g *Connect4Game) LegalMoves() []Move {

// }

// func (g *Connect4Game) Play(m Move) (Game, bool) {
// 	board1 := *g.board
// 	var move int = m.(int)
// 	board1.Play(moveInt)
// 	return 
// }

// type ObjectiveFunction func(*Board) (int, int)

// func HasWonObjFunc(b *Board) (int, int) {
// 	return hasWon(b, p1), hasWon(b, p2)
// }

// func hasWon(b *Board, p byte) int {
// 	for row:=0; row<ROWS; row++ {
// 		if maxConsecutiveInARow(b, row, p) == CONNECT {
// 			return 1
// 		}
// 	}
// 	for column:=0; column<COLUMNS; column++ {
// 		if maxConsecutiveInAColumn(b, column, p) == CONNECT {
// 			return 1
// 		}
// 	}
// 	return 0
// }

// type Solver interface {
// 	Solve(Board) int
// }

// type Minimax struct {
// 	f ObjectiveFunction
// }

// func (s *Minimax) Solve(b Board) int {
// 	max := b.Turn()
// 	min := p1
// 	if (max == min) { min = p2 }

// 	// for col := 1; col <= COLUMNS; col++ {


// 	return 1
// }

// // HELPER METHODS ============================================================

// func maxConsecutiveInARow(b *Board, row int, p byte) int {
// 	m := 0
// 	consecutive := 0
// 	for column:=0; column<COLUMNS; column++ {
// 		if p == b.board[row][column] {
// 			consecutive++
// 		} else {
// 			m = max(m, consecutive)
// 			consecutive = 0
// 		}
// 	}
// 	return m
// }

// func maxConsecutiveInAColumn(b *Board, column int, p byte) int {
// 	m := 0
// 	consecutive := 0
// 	for row:=0; row<ROWS; row++ {
// 		if p == b.board[row][column] {
// 			consecutive++
// 		} else {
// 			m = max(m, consecutive)
// 			consecutive = 0
// 		}
// 	}
// 	return m
// }

// func max(a, b int) int {
// 	if a > b { return a }
// 	return b
// }



























