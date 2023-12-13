package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"strings"
)

//go:embed input.txt

var input string
var lines []string

func init() {
	if input == "" {
		os.Exit(1)
	}
	lines = strings.Split(input, "\n")
	if strings.Contains(input, "\r\n") {
		lines = strings.Split(input, "\r\n")
	}

}

func main() {
	var totalScore int16
	for _, l := range lines {
		totalScore = totalScore + scoreLine(l)
	}
	fmt.Printf("total score: %v", totalScore)
}

func scoreLine(l string) int16 {

	s := strings.Split(l, "|")
	var lineScore float64 = 0
	var scoreCount float64 = -1
	if len(l) <= 1 {
		return int16(lineScore)
	}
	wins := strings.Split(s[0], ":")

	myList := numList(s[1])

	for _, a := range myList {
		if a == " " || a == "" {
			continue
		}
		if checkInSlice(wins[1], a) {
			// fmt.Printf("found %v in: %v\n", a, wins[1])
			scoreCount++
		}
	}
	lineScore = math.Pow(2, scoreCount)
	return int16(lineScore)
}

func checkInSlice(s string, c string) bool {
	k := strings.Split(s, " ")
	for _, x := range k {
		if x == c {
			return true
		}
	}
	return false
}

func numList(s string) []string {
	x := strings.Split(s, " ")
	return x
}
