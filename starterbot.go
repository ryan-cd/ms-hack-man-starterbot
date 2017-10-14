package main

import (
	"bufio"
	"fmt"
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
	moves     [4]string
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
	settings.moves = [4]string{"up", "left", "down", "right"}
	processInput()

}

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
			if settings.fieldHeight > 0 && settings.fieldWidth > 0 && len(game.field.field) != settings.fieldHeight {
				game.field.Initialize(settings.fieldWidth, settings.fieldHeight)
			}
		case "update":
			ParseUpdate(&game, command, settings.fieldWidth)
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
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, errorStr+"Scan error: %s\n", err)
	}
}

// DoMove sends intended movement to the engine
func DoMove() {
	//fmt.Println(settings.moves[rand.Intn(len(settings.moves))])
	fmt.Fprintf(os.Stdout, "pass\n")
}
