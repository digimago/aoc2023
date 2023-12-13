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
	for _, l := range lines[0:4] {
		totalScore = totalScore + scoreLine(l)
	}
	fmt.Printf("total score: %v", totalScore)
}

func scoreLine(l string) int16 {
	s := strings.Split(l, "|")
	var lineScore int16 = 0
	//fmt.Println(s[0])
	wins := strings.Split(s[0], ":")
	fmt.Println(wins[1])
	myList := numList(s[1])

	for _, a := range myList {
		if strings.Contains(wins[1], a) {
			fmt.Printf("found %v in: %v\n", a, wins[1])
		}
	}
	return lineScore
}

func numList(s string) []string {
	x := strings.Split(s, " ")
	return x
}
