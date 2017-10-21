package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var errorStr = "ERROR:"
var infoStr = "INFO:"

// Settings contains all initial game settings
type Settings struct {
	timebank, timePerMove, yourBotID, fieldWidth, fieldHeight, maxRounds int
	playerNames                                                          []string
	yourBot                                                              string

	character string // bixie or bixiette
}

// State holds the game features that update
type State struct {
	players       [2]Player
	field         Field
	round         int
	timeRemaining int
}

// Player holds state information for a player
type Player struct {
	snippets, bombs int
}

var scanner *bufio.Scanner
var settings = Settings{}
var game = State{}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	settings.character = "bixiette"
	processInput()

}

// Main game loop to handle communication from and to engine
func processInput() {
	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")

		if len(command) < 3 {
			return
		}
		switch command[0] {
		case "settings":
			ParseSettings(&settings, command)
			// Initialize field if it hasn't been, and the dimensions are known
			if settings.fieldHeight > 0 && settings.fieldWidth > 0 && len(game.field.GetBoard()) != settings.fieldHeight {
				game.field.Initialize(settings.fieldWidth, settings.fieldHeight, settings.yourBotID)
			}
		case "update":
			ParseUpdate(&game, command)
		case "action":
			switch ParseAction(&game, command) {
			case "character":
				fmt.Println(settings.character) // send the engine who we want to play as
			case "move":
				DoMove()
			default:
				fmt.Fprintf(os.Stderr, errorStr+"Unrecognized return from ParseAction")
			}
		default:
			fmt.Fprintf(os.Stderr, infoStr+"Received unhandled command type: %s\n", command)
			fmt.Fprintf(os.Stderr, infoStr+"Settings: %+v\n", settings)
			fmt.Fprintf(os.Stderr, infoStr+"State: %+v\n", game)
			game.field.PrintBoard()
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, errorStr+"Scan error: %s\n", err)
	}
}

// DoMove sends intended movement to the engine
func DoMove() {
	var validMoves []string
	if game.field.CanMove(UP) {
		validMoves = append(validMoves, "up")
	}
	if game.field.CanMove(LEFT) {
		validMoves = append(validMoves, "left")
	}
	if game.field.CanMove(DOWN) {
		validMoves = append(validMoves, "down")
	}
	if game.field.CanMove(RIGHT) {
		validMoves = append(validMoves, "right")
	}

	if len(validMoves) == 0 {
		fmt.Fprintln(os.Stdout, "pass")
	} else {
		fmt.Fprintln(os.Stdout, validMoves[rand.Intn(len(validMoves))])
	}
}
