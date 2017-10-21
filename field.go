package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var playerID int
var playerPos Point

// MoveType : The type for movement directions
type MoveType int

// Movement direction constants
const (
	UP    MoveType = 1 << iota
	LEFT  MoveType = 1 << iota
	DOWN  MoveType = 1 << iota
	RIGHT MoveType = 1 << iota
	PASS  MoveType = 1 << iota
)

// NodeType : The type for board cells
type NodeType int

// Constants for node types. If multiple objects are on a cell,
// the cell will be the bitwise or of the relevant types
const (
	EMPTY   NodeType = 1 << iota
	WALL    NodeType = 1 << iota
	PLAYER  NodeType = 1 << iota
	SPAWN   NodeType = 1 << iota
	GATE    NodeType = 1 << iota
	BUG     NodeType = 1 << iota
	MINE    NodeType = 1 << iota
	SNIPPET NodeType = 1 << iota
)

// Node represents a field cell
type Node struct {
	str      string
	nodeType NodeType
	modifier string //e.g. count down value
	coords   Point
}

// Point : Coordinates for a field location
type Point struct {
	x int
	y int
}

// NewPoint constructs a Point object
func NewPoint(x, y int) *Point {
	point := new(Point)
	point.x = x
	point.y = y
	return point
}

func (node Node) String() string {
	return fmt.Sprintf("%s ", node.str)
}

// NewNode constructs a Node object
func NewNode(str string, x, y int) *Node {
	node := new(Node)
	node.coords = *NewPoint(x, y)
	node.str = str
	node.modifier = ""
	for _, token := range strings.Split(str, ";") {
		if len(token) > 1 {
			node.modifier = token[1:]
		}

		switch token[0] {
		case '.':
			node.nodeType |= EMPTY
		case 'x':
			node.nodeType |= WALL
		case 'P':
			node.nodeType |= PLAYER
			id, err := strconv.Atoi(node.modifier)
			if err != nil {
				fmt.Fprintln(os.Stderr, errorStr+"Unable to convert player id to an int. Error:", err, "Detail: ", node.modifier)
			}
			if id == playerID {
				playerPos = *NewPoint(x, y)
			}
		case 'S':
			node.nodeType |= SPAWN
		case 'G':
			node.nodeType |= GATE
		case 'E':
			node.nodeType |= BUG
		case 'B':
			node.nodeType |= MINE
		case 'C':
			node.nodeType |= SNIPPET
		default:
			fmt.Fprintln(os.Stderr, errorStr+"Unable to parse node. Detail: ", token[0])
		}
	}
	return node
}

// Field stores the board data
type Field struct {
	board  [][]Node
	width  int
	height int
}

// PrintBoard : Debug function to see board representation
func (field *Field) PrintBoard() {
	fmt.Fprintln(os.Stderr, infoStr+"Current board:")
	for y := field.height - 1; y >= 0; y-- {
		for x := 0; x < field.width; x++ {
			fmt.Fprint(os.Stderr, field.board[x][y])
		}
		fmt.Fprintln(os.Stderr)
	}
}

// Initialize : Allocates a board with the specified dimensions, and records the bot ID
func (field *Field) Initialize(width, height, botID int) {
	playerID = botID
	if height > 0 && width > 0 && len(field.board) != height {
		field.width = width
		field.height = height
		for i := 0; i < width; i++ {
			boardCol := make([]Node, height)
			field.board = append(field.board, boardCol)
		}
	}
}

// GetBoard returns the board
func (field *Field) GetBoard() (board [][]Node) {
	return field.board
}

// SetField sets the board from a 1d input string received from the engine
func (field *Field) SetField(field1D []string) {
	for y := field.height - 1; y >= 0; y-- {
		for x := 0; x < field.width; x++ {
			field.board[x][y] = (*NewNode(field1D[field.width*(field.height-1-y)+x], x, y))
		}
	}
}

// CanMove takes a direction, and returns whether the player can move there
func (field *Field) CanMove(move MoveType) bool {
	switch move {
	case UP:
		if field.canMoveTo(playerPos.x, playerPos.y+1) {
			return true
		}
	case LEFT:
		if field.canMoveTo(playerPos.x-1, playerPos.y) {
			return true
		}
	case DOWN:
		if field.canMoveTo(playerPos.x, playerPos.y-1) {
			return true
		}
	case RIGHT:
		if field.canMoveTo(playerPos.x+1, playerPos.y) {
			return true
		}
	case PASS:
		return true
	default:
		fmt.Fprintln(os.Stderr, errorStr+"Trying to move to unhandled location. Detail:", move)
	}
	return false
}

func (field *Field) canMoveTo(x, y int) bool {
	if x < 0 || y < 0 || x >= field.width || y >= field.height {
		return false
	}

	nodeType := (*field).board[x][y].nodeType
	switch nodeType {
	case WALL & nodeType:
		return false
	case BUG & nodeType:
		return false
	}
	return true
}
