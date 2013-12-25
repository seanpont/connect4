package minimax

import (
	"testing"
	"github.com/seanpont/assert"
	"fmt"
)

// Note to self: the pointer to HeapGame implements the Game interface
type HeapGame struct {
	utility []int
	index int
	turn bool
	turns int
}

func (g *HeapGame) String() string {
	return fmt.Sprintf("index: %v, utility: %v, turn: %v",
		g.index, g.utility[g.index], g.turn)
}

func NewHeapGame(utility ...int) *HeapGame {
	game := new(HeapGame)
	game.utility = make([]int, 1, len(utility) + 1) // index 0 has value 0
	game.utility = append(game.utility, utility...)
	game.index = 1
	return game
}

func (game *HeapGame) Copy() Game {
	copy := *game
	return &copy
}

func (g *HeapGame) LegalMoves() []Move {
	return []Move { 0, 1 }
} 

func (g *HeapGame) Play(m Move) (bool, error) {
	g.turns++
	g.turn = !g.turn
	move, _ := m.(int)
	g.index = g.index * 2 + move
	isOver := g.index*2 > len(g.utility)
	return isOver, nil
}

func HeapUtilityFunction(g Game) int {
	hg, _ := g.(*HeapGame)
	return hg.utility[hg.index]
}

func TestMinimax(t *testing.T) {
	assert := assert.Assert(t)
	//                     p1   [--p2--]  [------p1------]
	game := NewHeapGame(0, 0,1, 0,1,-1,0, 0,1,-1,0,1,2,0,1 )
	fmt.Println(game.utility)
	// best move: 1
	solver := &Solver{ HeapUtilityFunction, 2 }
	move, utility := solver.Solve(game)
	moveInt := move.(int)
	assert.Equal(moveInt, 1)
	assert.Equal(utility, 1)
}

// Make sure I wrote heap game correctly before I try to win it.
func TestHeapGame(t *testing.T) {
	assert := assert.Assert(t)
	assert.Nil(nil, "good")

	f := HeapUtilityFunction

	game := NewHeapGame(1, 2, 3, 4, 5, 6, 7, 8)
	assert.Equal(f(game), 1)

	legalMoves := game.LegalMoves()
	move := legalMoves[0]
	game1 := game.Copy()
	over, _ := game1.Play(move)
	assert.False(over, "game not yet over")
	assert.NotEqual(game, game1)
	assert.Equal(f(game1), 2)

	game2 := game1.Copy()
	over, _ = game2.Play(1)
	assert.Equal(f(game2), 5)
	assert.True(over, "game should be over")
}



