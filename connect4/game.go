package connect4

import (
	"fmt"
	"errors"
)

const (
	ROWS = 6
	COLUMNS = 7
	CONNECT = 4
	empty = byte(0)
	p1 = byte(1)
	p2 = byte(2)
)

type Game struct {
	turn byte
	board [][]byte
}

func NewGame() *Game {
	g := new(Game)
	g.turn = 1
	g.board = make([][]byte, 6)
	for index := range g.board {
		g.board[index] = make([]byte, 7)
	}	
	return g
}

func (g *Game) Turn() byte {
	return g.turn
}

func (g *Game) Board() [][]byte {
	return g.board
}

func (g *Game) LegalMoves() []byte {
	moves := make([]byte, 0, 7)
	for col:=0; col<COLUMNS; col++ {
		if g.board[ROWS-1][col] == empty {
			moves = append(moves, byte(col+1))
		}
	}
	return moves
}

func (g *Game) Play(columnIndex int) (bool, error) {
	columnIndex -= 1
	if columnIndex < 0 || columnIndex >= COLUMNS {
		return false, errors.New(
			fmt.Sprintf("Column %d out of bounds [1,7]", columnIndex+1))
	}

	// find first open row
	rowIndex := 0
	for ; rowIndex < ROWS && g.board[rowIndex][columnIndex] > empty; rowIndex++ {}
	if rowIndex == ROWS {
		return false, errors.New("column full")
	}
	g.board[rowIndex][columnIndex] = g.turn

	//test for victory
	if g.testRow(rowIndex) || 
			g.testColumn(columnIndex) || 
			g.testLeftDiagonal(rowIndex, columnIndex) ||
			g.testRightDiagonal(rowIndex, columnIndex) {
		return true, nil	
	}

	g.turn += 1
	if g.turn == 3 {
		g.turn = 1
	}
	return false, nil
}

type ConnectCounter struct {
	int
}

func (c *ConnectCounter) incr(match bool) bool {
	if match {
		c.int += 1
		return c.int == CONNECT
	} else {
		c.int = 0
		return false
	}
}

func (g *Game) testRow(rowIndex int) bool {
	counter := new(ConnectCounter)
	for _, val := range g.board[rowIndex] {
		if counter.incr(val == g.turn) {
			return true
		}
	}
	return false
}

func (g *Game) testColumn(columnIndex int) bool {
	counter := new(ConnectCounter)
	for _, row := range g.board {
		if counter.incr(row[columnIndex] == g.turn) {
			return true
		}
	}
	return false
}

func (g *Game) testRightDiagonal(rowIndex, columnIndex int) bool {
	rightMax := ROWS - rowIndex - 1
	if columnIndex < rightMax {
		rightMax = columnIndex
	}
	rowIndex, columnIndex = rowIndex+rightMax, columnIndex-rightMax
	counter := new(ConnectCounter)
	for ; rowIndex >= 0 && columnIndex < COLUMNS; {
		if counter.incr(g.board[rowIndex][columnIndex] == g.turn) {
			return true
		}
		rowIndex -= 1
		columnIndex += 1
	}
	return false
}

func (g *Game) testLeftDiagonal(rowIndex, columnIndex int) bool {
	leftMax := rowIndex
	if columnIndex < rowIndex {
		leftMax = columnIndex
	}
	rowIndex, columnIndex = rowIndex-leftMax, columnIndex-leftMax
	counter := new(ConnectCounter)
	for ; rowIndex < ROWS && columnIndex < COLUMNS; {
		if counter.incr(g.board[rowIndex][columnIndex] == g.turn) {
			return true
		}
		rowIndex += 1
		columnIndex += 1
	}
	return false
}
 
func (g *Game) SetRows(rows ...string) {
	maxRowIndex := len(rows)-1
	for rowIndex, row := range rows {
		if len(row) > COLUMNS {
			panic("Row too long")
		}
		for columnIndex, val := range row {
			if val == '1' {
				g.board[maxRowIndex - rowIndex][columnIndex] = 1
			} else if val == '2' {
				g.board[maxRowIndex - rowIndex][columnIndex] = 2
			}
		}
	}
}

func (g *Game) String() string {
	s := ""
	for rowIndex := ROWS-1; rowIndex >= 0; rowIndex-- {
		s += fmt.Sprintf("%v\n", g.board[rowIndex])
	}
	s += fmt.Sprintf("Turn: %v", g.turn)
	return s
}


