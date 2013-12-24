package connect4

import (
	"testing"
	"github.com/seanpont/assert"
	_"fmt"
)

func Test_LegalMoves(t *testing.T) {
	assert := assert.Assert(t)
	g := NewGame()
	assert.Equal(g.LegalMoves(),
		[]byte{1,2,3,4,5,6,7})

	for row:=0; row<ROWS/2; row++ {
		g.Play(1); g.Play(1); g.Play(4); g.Play(4)
	}
	assert.Equal(g.LegalMoves(),
		[]byte{2,3,5,6,7})
}

func Test_ConnectCounter(t *testing.T) {
	assert := assert.Assert(t)
	cc := new(ConnectCounter)
	assert.False(cc.incr(true), "f1")
	assert.False(cc.incr(true), "f2")
	assert.False(cc.incr(true), "f3")
	assert.True(cc.incr(true), "t")
}

func Test_Play_returns_true_when_wins_right_diagonal(t *testing.T) {
	assert := assert.Assert(t)
	g := NewGame()

	g.SetRows(
		"0001000",
		"0012100",
		"0021200",
		"0212121",
		"0121210")
	win, err := g.Play(6)
	assert.Nil(err, "error after player 1 played final piece")
	assert.True(win, "player 1 should have won")
}

func Test_Play_returns_true_when_wins_left_diagonal(t *testing.T) {
	assert := assert.Assert(t)
	g := NewGame()

	g.SetRows(
		"0000000",
		"0001200",
		"0012200",
		"0121200")
	win, err := g.Play(5)
	assert.Nil(err, "error after player 1 played final piece")
	assert.True(win, "player 1 should have won")
}

func Test_Play_returns_true_when_wins_column(t *testing.T) {
	assert := assert.Assert(t)
	g := NewGame()

	players := []int {1, 2}
	for i:=1; i<4; i++ {
		for _, p := range players {
			won, err := g.Play(p) // player 1 plays column 1, player 2 column 2
			assert.Nil(err, "error after player %d played", p)
			assert.False(won, "player %v should not have won", p)
		}
	}

	won, err := g.Play(1)
	assert.Nil(err, "error after player 1 played final piece")
	assert.True(won, "player 1 should have won")
}

func Test_Play_returns_true_when_wins_in_row(t *testing.T) {
	assert := assert.Assert(t)
	g := NewGame()
	
	players := []int {1, 2}
	for i:=1; i<4; i++ {
		for _, p := range players {
			won, err := g.Play(i)
			assert.False(won, "player %v should not have won", p)
			assert.Nil(err, "error after player %d played", p)
		}
	}
	won, err := g.Play(4)
	assert.Nil(err, "error when player 1 played final piece")
	assert.True(won, "player 1 should have won")
}

func TestBoard_turn_switches_after_each_play(t *testing.T) {
	assert := assert.Assert(t)
	g := NewGame()
	assert.Equal(g.Turn(), byte(1));
	g.Play(1)
	assert.Equal(g.Turn(), byte(2));
	g.Play(1)
	assert.Equal(g.Turn(), byte(1));
}

func TestSetRows_sets_rows(t *testing.T) {
	assert := assert.Assert(t)
	g := NewGame()

	g.SetRows(
		"   1112",
		"  11212")

	expectedBoard := [][]byte {
		{0,0,1,1,2,1,2},
		{0,0,0,1,1,1,2},
		{0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0}}

	assert.Equal(g.board, expectedBoard)
}
