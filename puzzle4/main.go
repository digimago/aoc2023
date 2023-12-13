package main

import (
	_ "embed"
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

func main() {}
