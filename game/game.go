package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Game struct {
	board       [8][8]rune // [x][y]
	moveLog     []string
	pieces      map[string]rune
	currentMove string
}

var emptySpaceRune int = 160

// using pieces as parameter allows for interopability with different unicode chars
func NewBoard(pieces map[string]rune) [8][8]rune {
	var board [8][8]rune
	board[0] = [8]rune{
		pieces["WROOK"],
		pieces["WKNIGHT"],
		pieces["WBISHOP"],
		pieces["WQUEEN"],
		pieces["WKING"],
		pieces["WBISHOP"],
		pieces["WKNIGHT"],
		pieces["WROOK"],
	}
	board[1] = [8]rune{
		pieces["WPAWN"],
		pieces["WPAWN"],
		pieces["WPAWN"],
		pieces["WPAWN"],
		pieces["WPAWN"],
		pieces["WPAWN"],
		pieces["WPAWN"],
		pieces["WPAWN"],
	}
	board[2] = [8]rune{
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
	}
	board[3] = [8]rune{
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
	}
	board[4] = [8]rune{
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
	}
	board[5] = [8]rune{
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
		rune(emptySpaceRune),
	}
	board[6] = [8]rune{
		pieces["BPAWN"],
		pieces["BPAWN"],
		pieces["BPAWN"],
		pieces["BPAWN"],
		pieces["BPAWN"],
		pieces["BPAWN"],
		pieces["BPAWN"],
		pieces["BPAWN"],
	}
	board[7] = [8]rune{
		pieces["BROOK"],
		pieces["BKNIGHT"],
		pieces["BBISHOP"],
		pieces["BQUEEN"],
		pieces["BKING"],
		pieces["BBISHOP"],
		pieces["BKNIGHT"],
		pieces["BROOK"],
	}

	return board
}

func (g *Game) ReturnBoard() [8][8]rune {
	return g.board
}

func NewGame() *Game {
	pieces := map[string]rune{
		"WKING":   9818,
		"WQUEEN":  9819,
		"WROOK":   9820,
		"WBISHOP": 9821,
		"WKNIGHT": 9822,
		"WPAWN":   9823,
		"BKING":   9812,
		"BQUEEN":  9813,
		"BROOK":   9814,
		"BBISHOP": 9815,
		"BKNIGHT": 9816,
		"BPAWN":   9817,
	}

	board := NewBoard(pieces)

	return &Game{
		pieces:      pieces,
		board:       board,
		currentMove: "white",
	}
}

func (g *Game) Play() {
	for {
		fmt.Println(g.currentMove, "to move: ")

		if g.currentMove == "white" {
			g.PrintWhiteBoard()
		} else {
			g.PrintBlackBoard()
		}

		fmt.Print("What is your move: ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		if input.Text() == "end" {
			fmt.Println()
			fmt.Print(g.moveLog)
			break
		}

		g.Move(input.Text())
		fmt.Println()
	}
}

func (g *Game) SwitchMove() {
	// switch current move
	if g.currentMove == "white" {
		g.currentMove = "black"
	} else {
		g.currentMove = "white"
	}
}

func (g *Game) Move(move string) {
	// move is structured: e2e4
	// parse move, verify move is valid, change board, add to move log, check if move is checkmate, switch toMove
	coords := parseMove(move)

	// move board
	g.board[coords[2]][coords[3]] = g.board[coords[0]][coords[1]]
	g.board[coords[0]][coords[1]] = rune(emptySpaceRune)

	g.moveLog = append(g.moveLog, move)
	g.SwitchMove()

}

func (g *Game) PrintWhiteBoard() {
	for i := 7; i >= 0; i-- {
		for j := 0; j <= 7; j++ {
			fmt.Print(string(g.board[i][j]))
		}
		fmt.Println()
	}
}

func (g *Game) PrintBlackBoard() {
	for i := 0; i <= 7; i++ {
		for j := 7; j >= 0; j-- {
			fmt.Print(string(g.board[i][j]))
		}
		fmt.Println()
	}
}

func parseMove(move string) [4]int {
	moveLetterStartY := move[0:1]
	moveStartX, _ := strconv.Atoi(move[1:2])
	moveStartX -= 1
	moveLetterEndY := move[2:3]
	moveEndX, _ := strconv.Atoi(move[3:4])
	moveEndX -= 1

	// returns start y, start x, end y, end x
	var moveStartY, moveEndY int

	switch moveLetterStartY {
	case "a":
		moveStartY = 0
	case "b":
		moveStartY = 1
	case "c":
		moveStartY = 2
	case "d":
		moveStartY = 3
	case "e":
		moveStartY = 4
	case "f":
		moveStartY = 5
	case "g":
		moveStartY = 6
	case "h":
		moveStartY = 7
	}
	switch moveLetterEndY {
	case "a":
		moveEndY = 0
	case "b":
		moveEndY = 1
	case "c":
		moveEndY = 2
	case "d":
		moveEndY = 3
	case "e":
		moveEndY = 4
	case "f":
		moveEndY = 5
	case "g":
		moveEndY = 6
	case "h":
		moveEndY = 7
	}

	return [4]int{moveStartX, moveStartY, moveEndX, moveEndY}
}
