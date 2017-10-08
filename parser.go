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
	fmt.Println("INFO: Parsing settings: ", command)
	switch command[1] {
	case "timebank":
		time, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to convert command argument to int. Error:", err)
		}
		(*settings).timebank = time
	case "time_per_move":
		time, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to convert command argument to int. Error:", err)
		}
		(*settings).timePerMove = time
	case "player_names":
		names := strings.Split(command[2], ",")
		if len(names) != 2 {
			fmt.Fprintln(os.Stderr, "player_names was unable to parse into []string of length 2. Detail: names=", names)
		}
		(*settings).playerNames = names
	case "your_bot":
		(*settings).yourBot = command[2]
	case "your_botid":
		ID, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to parse bot id. Error: ", err, " Detail: ", command)
		}
		(*settings).yourBotID = ID
	case "field_width":
		width, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to parse width. Error: ", err, " Detail: ", command)
		}
		(*settings).fieldWidth = width
	case "field_height":
		height, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to parse height. Error: ", err, " Detail: ", command)
		}
		(*settings).fieldHeight = height
	case "max_rounds":
		rounds, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to parse max rounds. Error: ", err, " Detail: ", command)
		}
		(*settings).maxRounds = rounds
	default:
		fmt.Fprintln(os.Stderr, "Unhandled settings field. Detail:", command)
	}
}

// ParseUpdate takes a Settings object to modify,
// and a command tuple of the format "update game round 0"
func ParseUpdate(settings *Settings, command []string) {
	fmt.Println("INFO: Parsing update: ", command)
	switch command[2] {
	//case "round":
	}
}

// ParseAction takes a State object to modify,
// and a command tuple of the format "action character t"
func ParseAction(state *State, command []string) (commandType string) {
	fmt.Println("INFO: Parsing action: ", command)
	timeRemaining, err := strconv.Atoi(command[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Unable to parse time remaining. Detail: ", command)
	}
	(*state).timeRemaining = timeRemaining

	return command[1]
}
