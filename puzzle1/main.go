package main

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"

	"fmt"
)

//go:embed input.txt
var input string

func getFirst(s string) int {

	pattern := regexp.MustCompile(`(?i)(?:)((one|two|three|four|five|six|seven|eight|nine)|(\d))`)

	matches := pattern.FindAllString(s, -1)
	// for _, match := range matches {
	// 	fmt.Println(match)
	// }
	//fmt.Println(matches[0])
	return convWordToInt(matches[0])
}

func StringFlip(s string) string {
	r := make([]string, len(s))
	for i := len(s); i > 0; i-- {
		r = append(r, string(s[i-1]))
	}
	return strings.Join(r, "")
}

func getLast(s string) int {
	s = StringFlip(s)
	p := StringFlip("one|two|three|four|five|six|seven|eight|nine")
	px := fmt.Sprintf(`(?i)(?:)((%v)|(\d))`, p)
	pattern := regexp.MustCompile(px)

	matches := pattern.FindAllString(s, -1)

	return convWordToInt(StringFlip(matches[0]))
}

func main() {
	var calibration int64
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		// got a line, now process the runes

		var cal []string
		if l == "" {
			continue
		}

		cal = append(cal, fmt.Sprint(getFirst(l)))
		cal = append(cal, fmt.Sprint(getLast(l)))
		y := strings.Join(cal, "")
		c, _ := strconv.ParseInt(y, 10, 64)
		calibration = calibration + c

		fmt.Printf("c: %v - calibration: %v\n", c, calibration)
	}
	fmt.Println(calibration)
}

func convWordToInt(s string) int {
	var i int
	switch s {
	case
		"zero":
		i = 0
	case
		"one":
		i = 1
	case
		"two":
		i = 2
	case
		"three":
		i = 3
	case
		"four":
		i = 4
	case
		"five":
		i = 5
	case
		"six":
		i = 6
	case
		"seven":
		i = 7
	case
		"eight":
		i = 8
	case
		"nine":
		i = 9

	default:
		i, _ = strconv.Atoi(s)

	}
	return i
}
