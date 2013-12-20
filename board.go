package main

import (
	"fmt"
	"errors"
)

const (
	ROWS = 6
	COLUMNS = 7
	CONNECT = 4
)

type Board struct {
	turn byte
	board [][]byte
}

func NewBoard() *Board {
	b := new(Board)
	b.turn = 1
	b.board = make([][]byte, 6)
	for index, _ := range b.board {
		b.board[index] = make([]byte, 7)
	}	
	return b
}

func (b *Board) Turn() byte {
	return b.turn
}

func (b *Board) Play(columnIndex int) (bool, error) {
	columnIndex -= 1
	if columnIndex < 0 || columnIndex >= COLUMNS {
		return false, errors.New(
			fmt.Sprintf("Column %d out of bounds [1,7]", columnIndex+1))
	}

	// find first open row
	rowIndex := 0
	for ; rowIndex < ROWS && b.board[rowIndex][columnIndex] > 0; rowIndex++ {}
	if rowIndex == ROWS {
		return false, errors.New("column full")
	}
	b.board[rowIndex][columnIndex] = b.turn

	//test for victory
	if b.testRow(rowIndex) || 
			b.testColumn(columnIndex) || 
			b.testLeftDiagonal(rowIndex, columnIndex) ||
			b.testRightDiagonal(rowIndex, columnIndex) {
		return true, nil	
	}

	b.turn += 1
	if b.turn == 3 {
		b.turn = 1
	}
	return false, nil
}

func (b *Board) testRow(rowIndex int) bool {
	consecutive := 0
	for _, val := range b.board[rowIndex] {
		if val == b.turn {
			consecutive += 1
			if consecutive == CONNECT {
				return true
			}
		} else {
			consecutive = 0
		}
	}
	return false
}

func (b *Board) testColumn(columnIndex int) bool {
	consecutive := 0
	for _, row := range b.board {
		if row[columnIndex] == b.turn {
			consecutive += 1
			if consecutive == CONNECT {
				return true
			}
		} else {
			consecutive = 0
		}
	}
	return false
}

func (b *Board) testRightDiagonal(rowIndex, columnIndex int) bool {
	rightMax := ROWS - rowIndex - 1
	if columnIndex < rightMax {
		rightMax = columnIndex
	}
	rowIndex, columnIndex = rowIndex+rightMax, columnIndex-rightMax
	consecutive := 0
	for ; rowIndex >= 0 && columnIndex < COLUMNS; {
		if b.board[rowIndex][columnIndex] == b.turn {
			consecutive += 1
			if consecutive == CONNECT {
				return true
			}
		} else {
			consecutive = 0
		}
		rowIndex -= 1
		columnIndex += 1
	}
	return false
}

func (b *Board) testLeftDiagonal(rowIndex, columnIndex int) bool {
	leftMax := rowIndex
	if columnIndex < rowIndex {
		leftMax = columnIndex
	}
	rowIndex, columnIndex = rowIndex-leftMax, columnIndex-leftMax
	consecutive := 0
	for ; rowIndex < ROWS && columnIndex < COLUMNS; {
		if b.board[rowIndex][columnIndex] == b.turn {
			consecutive += 1
			if consecutive == CONNECT {
				return true
			}
		} else {
			consecutive = 0
		}
		rowIndex += 1
		columnIndex += 1
	}
	return false
}
 
func (b *Board) SetRows(rows ...string) {
	maxRowIndex := len(rows)-1
	for rowIndex, row := range rows {
		for columnIndex, val := range row {
			if val == '1' {
				b.board[maxRowIndex - rowIndex][columnIndex] = 1
			} else if val == '2' {
				b.board[maxRowIndex - rowIndex][columnIndex] = 2
			}
		}
	}
}

func (b *Board) String() string {
	s := ""
	for rowIndex := ROWS-1; rowIndex >= 0; rowIndex-- {
		s += fmt.Sprintf("%v\n", b.board[rowIndex])
	}
	s += fmt.Sprintf("Turn: %v", b.turn)
	return s
}


