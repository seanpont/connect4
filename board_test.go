package main

import (
	"testing"
	"github.com/seanpont/assert"
	"fmt"
)

func Test_Play_returns_true_when_wins_right_diagonal(t *testing.T) {
	assert := assert.Assert(t)
	b := NewBoard()

	b.SetRows(
		"0001000",
		"0012100",
		"0021200",
		"0212121",
		"0121210")
	win, err := b.Play(6)
	assert.Nil(err, "error after player 1 played final piece")
	assert.True(win, "player 1 should have won")
}

func Test_Play_returns_true_when_wins_left_diagonal(t *testing.T) {
	assert := assert.Assert(t)
	b := NewBoard()

	b.SetRows(
		"0000000",
		"0001200",
		"0012200",
		"0121200")
	win, err := b.Play(5)
	assert.Nil(err, "error after player 1 played final piece")
	assert.True(win, "player 1 should have won")
}

func Test_Play_returns_true_when_wins_column(t *testing.T) {
	assert := assert.Assert(t)
	b := NewBoard()

	players := []int {1, 2}
	for i:=1; i<4; i++ {
		for _, p := range players {
			won, err := b.Play(p) // player 1 plays column 1, player 2 column 2
			assert.Nil(err, "error after player %d played", p)
			assert.False(won, "player %v should not have won", p)
		}
	}

	won, err := b.Play(1)
	assert.Nil(err, "error after player 1 played final piece")
	assert.True(won, "player 1 should have won")
}

func Test_Play_returns_true_when_wins_in_row(t *testing.T) {
	assert := assert.Assert(t)
	b := NewBoard()
	
	players := []int {1, 2}
	for i:=1; i<4; i++ {
		for _, p := range players {
			won, err := b.Play(i)
			assert.False(won, "player %v should not have won", p)
			assert.Nil(err, "error after player %d played", p)
		}
	}
	won, err := b.Play(4)
	assert.Nil(err, "error when player 1 played final piece")
	assert.True(won, "player 1 should have won")
}

func TestBoard_turn_switches_after_each_play(t *testing.T) {
	assert := assert.Assert(t)
	b := NewBoard()
	
	assert.Equal(b.Turn(), byte(1));
	b.Play(1)
	assert.Equal(b.Turn(), byte(2));
	b.Play(1)
	assert.Equal(b.Turn(), byte(1));
}

func TestSetRows_sets_rows(t *testing.T) {
	assert := assert.Assert(t)
	b := NewBoard()

	b.SetRows(
		"   1112",
		"  11212")

	expectedBoard := [][]byte {
		{0,0,1,1,2,1,2},
		{0,0,0,1,1,1,2},
		{0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0}}

	assert.Equal(b.board, expectedBoard)
}
