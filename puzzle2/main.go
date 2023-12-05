package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var lines []string

const Red int = 12
const Green int = 13
const Blue int = 14

type Game struct {
	number int
	moves  []Move
	valid  bool
	power  Move
}

type Move struct {
	red   int
	green int
	blue  int
}

func init() {
	if input == "" {
		os.Exit(1)
	}

	lines = strings.Split(input, "\r\n")
}

func lineToGame(s string) Game {
	var game Game
	line := strings.Split(s, ":")
	g := strings.Split(line[0], " ")
	game.number, _ = strconv.Atoi(g[1])
	game = makeMoves(line[1], game)
	return game
}

func checkInvalidMove(m Move) bool {

	if m.red > Red || m.green > Green || m.blue > Blue {
		return true
	}
	return false
}

func makeMoves(s string, g Game) Game {
	var moves []Move

	g.valid = true
	p := strings.Split(s, ";")
	for _, j := range p {
		var m Move

		k := strings.Split(j, ",")
		for _, l := range k {
			c := strings.Split(strings.TrimSpace(l), " ")
			if strings.Contains(l, "red") {
				//red
				m.red, _ = strconv.Atoi(c[0])
				if m.red > g.power.red {
					g.power.red = m.red
				}
			}
			if strings.Contains(l, "green") {
				//green
				m.green, _ = strconv.Atoi(c[0])
				if m.green > g.power.green {
					g.power.green = m.green
				}
			}
			if strings.Contains(l, "blue") {
				//blue
				m.blue, _ = strconv.Atoi(c[0])
				if m.blue > g.power.blue {
					g.power.blue = m.blue
				}
			}

		}

		if checkInvalidMove(m) {
			// fmt.Printf("Invalid move detected: %v for move %v\n", g.number, m)
			g.valid = false
		}
		moves = append(moves, m)

	}
	g.moves = moves
	return g
}

func main() {

	var sum int

	var gsum int
	for _, l := range lines {
		var g Game
		if len(l) < 6 {
			continue
		}
		g = lineToGame(l)
		if g.valid {
			sum = sum + g.number
		}
		gsum = gsum + (g.power.blue * g.power.green * g.power.red)

	}
	fmt.Printf("Result part 1: %v\n", sum)
	fmt.Printf("Result part 2: %v\n", gsum)
}
