package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ParseSettings takes a Settings object to modify,
// and a command tuple of the format "settings timebank 10000"
func ParseSettings(settings *Settings, command []string) {
	switch command[1] {
	case "timebank":
		time, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, errorStr+"Unable to convert command argument to int. Error:", err)
		}
		(*settings).timebank = time
	case "time_per_move":
		time, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, errorStr+"Unable to convert command argument to int. Error:", err)
		}
		(*settings).timePerMove = time
	case "player_names":
		names := strings.Split(command[2], ",")
		if len(names) != 2 {
			fmt.Fprintln(os.Stderr, errorStr+"player_names was unable to parse into []string of length 2. Detail: names=", names)
		}
		(*settings).playerNames = names
	case "your_bot":
		(*settings).yourBot = command[2]
	case "your_botid":
		ID, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, errorStr+"Unable to parse bot id. Error: ", err, " Detail: ", command)
		}
		(*settings).yourBotID = ID
	case "field_width":
		width, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, errorStr+"Unable to parse width. Error: ", err, " Detail: ", command)
		}
		(*settings).fieldWidth = width
	case "field_height":
		height, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, errorStr+"Unable to parse height. Error: ", err, " Detail: ", command)
		}
		(*settings).fieldHeight = height
	case "max_rounds":
		rounds, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, errorStr+"Unable to parse max rounds. Error: ", err, " Detail: ", command)
		}
		(*settings).maxRounds = rounds
	default:
		fmt.Fprintln(os.Stderr, errorStr+"Unhandled settings field. Detail:", command)
	}
}

// ParseUpdate takes a Settings object to modify,
// and a command tuple of the format "update game round 0"
func ParseUpdate(state *State, command []string, fieldWidth int) {
	switch command[2] {
	case "round":
		round, err := strconv.Atoi(command[3])
		if err != nil {
			fmt.Fprintln(os.Stderr, errorStr+"Unable to parse round. Error:", err, "Detail: ", command)
		}
		(*state).round = round
	case "field":
		field1D := strings.Split(command[3], ",")

		if fieldWidth > 0 {
			for i := 0; i < len(field1D); i += fieldWidth {
				(*state).field.field[i/fieldWidth] = field1D[i : i+fieldWidth]
			}
		}
	case "snippets":
		snippets, err := strconv.Atoi(command[3])
		if err != nil {
			fmt.Fprintln(os.Stderr, errorStr+"Unable to parse snippets. Error:", err, "Detail: ", command)
		}
		if command[1] == "player0" {
			(*state).players[0].snippets = snippets
		} else if command[1] == "player1" {
			(*state).players[1].snippets = snippets
		} else {
			fmt.Fprintln(os.Stderr, errorStr+"Unhandled player enountered in update. Player: ", command[1])
		}
	case "bombs":
		bombs, err := strconv.Atoi(command[3])
		if err != nil {
			fmt.Fprintln(os.Stderr, errorStr+"Unable to parse bombs. Error:", err, "Detail: ", command)
		}
		if command[1] == "player0" {
			(*state).players[0].bombs = bombs
		} else if command[1] == "player1" {
			(*state).players[1].bombs = bombs
		} else {
			fmt.Fprintln(os.Stderr, errorStr+"Unhandled player enountered in update. Player: ", command[1])
		}
	default:
		fmt.Fprintln(os.Stderr, errorStr+"Unhandled update type. Detail:", command)
	}
}

// ParseAction takes a State object to modify,
// and a command tuple of the format "action character t"
// returns either "character" or "move"
func ParseAction(state *State, command []string) (commandType string) {
	timeRemaining, err := strconv.Atoi(command[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, errorStr+"Unable to parse time remaining. Detail: ", command)
	}
	(*state).timeRemaining = timeRemaining

	return command[1]
}
