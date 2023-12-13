package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
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
	pattern := regexp.MustCompile(`(?i)(?:)(\d+)`)

	m := pattern.FindAllString(lines[0], -1)
	fmt.Println(m)
}
