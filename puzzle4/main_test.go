package main

import (
	_ "embed"
	"testing"
)

func Test_scoreLine(t *testing.T) {
	type args struct {
		l string
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreLine(tt.args.l); got != tt.want {
				t.Errorf("scoreLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
