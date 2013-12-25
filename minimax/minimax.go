package minimax

import (
	"sort"
	"fmt"
	"strings"
)

type Game interface {
	Copy() Game
	LegalMoves() []Move
	Play(Move) (bool, error)
}

type Move interface {}

type GameState struct {
	Game
	Move
	Utility int
}

func (gs GameState) String() string {
	return fmt.Sprintf("{Move: %v, Utility: %v}", gs.Move, gs.Utility)
}

type ByUtility []GameState

func (a ByUtility) Len() int           { return len(a) }
func (a ByUtility) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUtility) Less(i, j int) bool { return a[i].Utility < a[j].Utility }
func (a ByUtility) Max() GameState {
	i := 0
	for j := range a { if a.Less(i, j) { i = j } }
	return a[i]
}
func (a ByUtility) Min() GameState {
	i := 0
	for j := range a { if a.Less(j, i) { i = j } }
	return a[i]
}

// Returns the 'utility' of the game state with respect to the player
// that just played. For instace, if p1 has just captured a pawn (all else equal),
// then it should return 1. If p1 was down a pawn but made no capture, 
// then it should return -1
type UtilityFunction func(Game) int

type Solver struct {
	UtilityFunction
	Depth int
	Verbose bool
}

func (s *Solver) Solve(game Game) (move Move, utility int) {
	gameState := s.explore(game, 0)
	return gameState.Move, gameState.Utility
}

func (s *Solver) explore(game Game, depth int) GameState {
	spaces := strings.Repeat(" ", depth * 4)
	if (s.Verbose) {
		gameStr := spaces+strings.Replace(fmt.Sprint(game), "\n", "\n"+spaces, -1)
		fmt.Println(spaces + "Exploring game:\n", gameStr)
		fmt.Println(spaces + "Depth:", depth)
	}
	legalMoves := game.LegalMoves()
	gameStates := make([]GameState, 0, len(legalMoves))
	if (s.Verbose) {
		fmt.Println(spaces + "Legal moves:", legalMoves)
	}
	for _, move := range legalMoves {
		if (s.Verbose) {
			fmt.Println(spaces + "Exploring move:", move)
		}
		game1 := game.Copy()
		over, err := game1.Play(move)
		if (err != nil) { panic(err) }
		gameState := GameState{ game1, move, s.UtilityFunction(game1) }
		if over {
			if (s.Verbose) {
				fmt.Println(spaces + "Terminal game state found:", gameState)
			}
			return gameState
		} 
		gameStates = append(gameStates, gameState)			
	}
	sort.Sort(sort.Reverse(ByUtility(gameStates)))
	if s.Verbose {
		fmt.Println(spaces + "Game states:")
		for _, gameState := range gameStates { 
			fmt.Println( spaces, gameState) 
		}
	}
	if depth < s.Depth {
		for i := range gameStates {
			gameStates[i].Utility = 
				-1 * s.explore(gameStates[i].Game, depth+1).Utility
			if (s.Verbose) {
				gameState := gameStates[i]
				fmt.Printf(spaces + "Utility of move %v after search: %v\n", 
					gameState.Move, gameState.Utility)
			}
		}
		if (s.Verbose) {
			fmt.Println(spaces + "Game states after exploration:")
			for _, gameState := range gameStates { 
				fmt.Println( spaces, gameState) 
			}
		}
	}
	best := ByUtility(gameStates).Max()
	if (s.Verbose) {
		fmt.Printf(spaces + "Best move: %v, Utility: %v\n", best.Move, best.Utility)
	}
	return best
}

















