package connect4

import (
	"testing"
	// "fmt"
	"github.com/seanpont/assert"
	"github.com/seanpont/connect4/minimax"
)

// SOLVER ====================================================================

func Test_Solve_aver_vertical_disaster(t *testing.T) {
	assert := assert.Assert(t)
	game := NewGame()
	game.SetRows(
		"1000000",
		"2010000",
		"2211000",
		"2211120")

	over, _ := game.Play(2)
	assert.False(over, "wtf")

	// game.turn = p2
	// move := SolveVerbose(game, 2)
	// assert.Equal(move, 3)

}

func Test_Solve_avert_disaster(t *testing.T) {
	assert := assert.Assert(t)
	game := NewGame()
	game.SetRows(
		"2000000",
		"2000000",
		"1110000")
	game.turn = p2
	move := Solve(game, 2)
	assert.Equal(move, 4)
}

// Utility Functions =========================================================

func Test_HasWonUtilityFunction_returns_1_when_won(t *testing.T) {
	assert := assert.Assert(t)
	game := NewGame()
	wrapper := GameWrapper{ game }
	game.SetRows("1110000")
	assert.Equal(HasWon(wrapper), 0)
	hasWon, _ := game.Play(4) // now it's player 2's turn, but 1 won
	assert.True(hasWon, "should have won")
	assert.Equal(HasWon(wrapper), 1)
}

func Test_HasWonUtilityFuncion_implements_correct_inverface(t *testing.T) {
	assert := assert.Assert(t)
	solver := minimax.Solver { HasWon, 4, false }
	assert.NotNil(solver, "solver not nil")
}

// C4Game Wrapper ============================================================

func Test_C4Game(t *testing.T) {
	assert := assert.Assert(t)
	game := GameWrapper{NewGame()}
	g := minimax.Game(game)
	assert.NotNil(g, "not nil")

	game1 := game.Copy()
	game1.Play(1)
	assert.NotEqual(game, game1)
}



