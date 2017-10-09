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
	moves     [4]string
}

// State holds the game features that update
type State struct {
	players       [2]Player
	field         [][]string
	round         int
	timeRemaining int
}

// Player holds state information for a player
type Player struct {
	snippets, bombs int
}

var settings = Settings{}
var game = State{}

func main() {
	settings.character = "bixiette"
	settings.moves = [4]string{"up", "left", "down", "right"}
	for {
		processInput()
	}
}

func processInput() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, errorStr+"Scan error: %s\n", err)
		}
	}

	command := strings.Split(scanner.Text(), " ")
	if len(command) < 3 {
		fmt.Fprintf(os.Stderr, errorStr+"Invalid command block length: %s\n", command)
	}
	switch command[0] {
	case "settings":
		ParseSettings(&settings, command)
		// Initialize field if it hasn't been, and the dimensions are known
		if settings.fieldHeight > 0 && settings.fieldWidth > 0 && len(game.field) != settings.fieldHeight {
			for i := 0; i < settings.fieldHeight; i++ {
				boardRow := make([]string, settings.fieldWidth)
				game.field = append(game.field, boardRow)
			}
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
	}

}

// DoMove sends intended movement to the engine
func DoMove() {
	fmt.Println(settings.moves[rand.Intn(len(settings.moves))])
}
