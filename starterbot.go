package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Settings contains all initial game settings
type Settings struct {
	timebank, timePerMove, yourBotID, fieldWidth, fieldHeight, maxRounds int
	playerNames                                                          []string
	yourBot                                                              string

	character string // bixie or bixiette
}

// State holds the player and field info
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
	//game.players[0], game.players[1] = Player{}, Player{}
	settings.character = "bixiette"
	for {
		processInput()
	}
}

func processInput() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Scan error: %s\n", err)
		}
	}

	command := strings.Split(scanner.Text(), " ")
	if len(command) < 3 {
		fmt.Fprintf(os.Stderr, "Invalid command block length: %s\n", command)
	}
	switch command[0] {
	case "settings":
		ParseSettings(&settings, command)
	case "update":
		ParseUpdate(&settings, command)
	case "action":
		switch ParseAction(&game, command) {
		case "character":
			fmt.Println(settings.character) // send the engine who we want to play as
		case "move":
			DoMove()
		default:
			fmt.Fprintf(os.Stderr, "ERROR: Unrecognized return from ParseAction")
		}
	default:
		fmt.Fprintf(os.Stderr, "Received unhandled command type: %s\n", command)
		fmt.Fprintf(os.Stderr, "Settings: %+#v\n", settings)
	}

}

func DoMove() {
	fmt.Println("right")
}
