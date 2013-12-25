package connect4

import (
	_"fmt"
	"github.com/seanpont/connect4/minimax"
)

func Solve(game *Game, depth int, verbose ...bool) int {
	isVerbose := verbose != nil && len(verbose) == 1 && verbose[0]
	solver := &minimax.Solver{ HasWon, depth, isVerbose }
	move, _ := solver.Solve( GameWrapper{game} )
	return move.(int)
}

func SolveVerbose(game *Game, depth int) int {
	return Solve(game, depth, true)
}

// Wrapper around connect4 game to make it match minimax.Game interface
type GameWrapper struct {
	*Game
}

func (g GameWrapper) Copy() minimax.Game {
	return GameWrapper{g.Game.Copy()}
}

func (g GameWrapper) LegalMoves() []minimax.Move {
	moves := g.Game.LegalMoves()
	moves2 := make([]minimax.Move, 0, len(moves))
	for _, move := range moves {
		moves2 = append(moves2, minimax.Move(move))
	}
	return moves2
}

func (g GameWrapper) Play(move minimax.Move) (bool, error) {
	return g.Game.Play(move.(int))
}

// UTILITY FUNCTION ==========================================================

func HasWon(g minimax.Game) int {
	game := g.(GameWrapper).Game
	player := game.turn % 2 + 1 // get other player
	for i:=0; i<ROWS; i++ {
		if game.testRow(player, i) { return 1 }
	}
	for i:=0; i<COLUMNS; i++ {
		if game.testColumn(player, i) { return 1 }
	}
	for i:=0; i<COLUMNS-3; i++ {
		if game.testLeftDiagonal(player, 2, i) { return 1 }
	}
	for i:=3; i>COLUMNS; i++ {
		if game.testRightDiagonal(player, 0, i) { return 1 }
	}
	return 0
}























