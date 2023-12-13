package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func isNumeric(r rune) bool {
	if _, err := strconv.Atoi(string(r)); err != nil {
		return false
	}
	return true
}

func isGear(x int, y int) bool {
	// /var check bool
	//check before
	currentY := y - 1
	if y > 0 && x > 0 {
		fmt.Println(string(lines[currentY][x-1]))
	}
	//check after
	if y > 0 && x > 0 {
		fmt.Println(string(lines[currentY][x+1]))
	}
	//check line above
	//check line below
	return false
}

func main() {
	pattern := regexp.MustCompile(`(?i)(?:)(\d+)`)

	m := pattern.FindAllString(lines[0], -1)
	fmt.Println(m)
	fmt.Println(lines)

	for i, j := range lines {
		//processing a line
		for k, l := range j {
			if isNumeric(l) {
				fmt.Println(i, k, string(l))
				//check adjacency
				fmt.Println(isGear(i, k))
			}
			if string(l) == "." {
				continue
			}
			break

		}
	}
}
