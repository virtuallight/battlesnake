package main

// This file can be a nice home for your Battlesnake logic and related helper functions.
//
// We have started this for you, with a function to help remove the 'neck' direction
// from the list of possible moves!

import (
	"log"
	"math/rand"
)

// This function is called when you register your Battlesnake on play.battlesnake.com
// See https://docs.battlesnake.com/guides/getting-started#step-4-register-your-battlesnake
// It controls your Battlesnake appearance and author permissions.
// For customization options, see https://docs.battlesnake.com/references/personalization
// TIP: If you open your Battlesnake URL in browser you should see this data.
func info() BattlesnakeInfoResponse {
	log.Println("INFO")
	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "virtuallight",
		Color:      "#88ff88",
		Head:       "smile",
		Tail:       "bolt",
	}
}

// This function is called everytime your Battlesnake is entered into a game.
// The provided GameState contains information about the game that's about to be played.
// It's purely for informational purposes, you don't have to make any decisions here.
func start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// This function is called when a game your Battlesnake was in has ended.
// It's purely for informational purposes, you don't have to make any decisions here.
func end(state GameState) {
	log.Printf("%s END\n\n", state.Game.ID)
}

// This function is called on every turn of a game. Use the provided GameState to decide
// where to move -- valid moves are "up", "down", "left", or "right".
// We've provided some code and comments to get you started.
func move(state GameState) BattlesnakeMoveResponse {

	myHead := state.You.Head
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height

	// Step 0 - Init `possibleMoves` map
	// Step 1 - Don't hit walls.
	possibleMoves := map[string]bool{
		"up":    myHead.Y < boardHeight-1,
		"down":  myHead.Y > 0,
		"left":  myHead.X > 0,
		"right": myHead.X < boardWidth-1,
	}

	// Step 2 - Don't hit yourself.
	// Step 3 - Don't collide with others.
	toTheLeft := Coord{X: myHead.X - 1, Y: myHead.Y}
	toTheRight := Coord{X: myHead.X + 1, Y: myHead.Y}
	toTheTop := Coord{X: myHead.X, Y: myHead.Y + 1}
	toTheBottom := Coord{X: myHead.X, Y: myHead.Y - 1}

	for _, snake := range state.Board.Snakes {
		for _, bodyPartCoord := range snake.Body {
			if bodyPartCoord == toTheLeft {
				possibleMoves["left"] = false
			} else if bodyPartCoord == toTheRight {
				possibleMoves["right"] = false
			} else if bodyPartCoord == toTheTop {
				possibleMoves["up"] = false
			} else if bodyPartCoord == toTheBottom {
				possibleMoves["down"] = false
			}
		}
	}

	// TODO: Step 4 - Find food.
	// Use information in GameState to seek out and find food.

	// Finally, choose a move from the available safe moves.
	// TODO: Step 5 - Select a move to make based on strategy, rather than random.
	var nextMove string

	safeMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}
	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}
