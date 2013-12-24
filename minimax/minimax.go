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
	return fmt.Sprintf("{Game: %v, Move: %v, Utility: %v}", 
		gs.Game, gs.Move, gs.Utility)
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

type MinimaxSolver struct {
	UtilityFunction
	depth int
}

func (s *MinimaxSolver) Solve(game Game) (move Move, utility int) {
	gameState := s.explore(game, 0)
	return gameState.Move, gameState.Utility
}

func (s *MinimaxSolver) explore(game Game, depth int) GameState {
	spaces := strings.Repeat(" ", depth * 4)
	fmt.Println(spaces + "Exploring game:", game)
	fmt.Println(spaces + "Depth:", depth)
	gameStates := make([]GameState, 0)

	for move := range game.LegalMoves() {
		game1 := game.Copy()
		over, err := game1.Play(move)
		if (err != nil) { panic(err) }
		gameState := GameState{ game1, move, s.UtilityFunction(game1) }
		if over {
			fmt.Println(spaces + "Terminal game state found:", gameState)
			return gameState
		} 
		gameStates = append(gameStates, gameState)			
	}
	sort.Sort(sort.Reverse(ByUtility(gameStates)))
	fmt.Println(spaces + "Game states:")
	for _, gameState := range gameStates { fmt.Println( spaces, gameState) }
	if depth < s.depth {
		for _, gameState := range gameStates {
			gameState.Utility = -1 * s.explore(gameState.Game, depth+1).Utility
			fmt.Println(spaces + "Utility after search:", gameState.Utility)
		}
		fmt.Println(spaces + "Game states after exploration:")
		for _, gameState := range gameStates { fmt.Println( spaces, gameState) }
	}
	fmt.Println(spaces + "Best move:", ByUtility(gameStates).Max().Move)
	return ByUtility(gameStates).Max()
}

















