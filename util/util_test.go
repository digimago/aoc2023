package util

import (
	"fmt"
	"testing"
)

func TestIsNumber(t *testing.T) {

	for _, tt := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"} {
		t.Run(fmt.Sprintf("Testing %v", tt), func(t *testing.T) {
			r := []rune(tt)
			if got := IsNumber(r[0]); got != true {
				t.Errorf("IsNumber() = %v, want %v", got, true)
			}
		})
	}
}
