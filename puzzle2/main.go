package main

import (
	_ "embed"
	"fmt"
	"os"
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

func lineToGame(s string) Move {
	var move Move
	line := strings.Split(s, ";")
}

func main() {
	var result string
	fmt.Printf("Result: %v", result)
}
