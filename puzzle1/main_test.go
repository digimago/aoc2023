package main

import (
	_ "embed"
	"testing"
)

func TestStringFlip(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		// TODO: Add test cases.

		{
			name: "abc",
			args: "abc",
			want: "cba"},

		{
			name: "abc123",
			args: "abc123",
			want: "321cba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringFlip(tt.args); got != tt.want {
				t.Errorf("StringFlip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLast(t *testing.T) {

	tests := []struct {
		name string
		args string
		want int
	}{
		// TODO: Add test cases.
		{
			name: "abc1",
			args: "abc1",
			want: 1,
		},
		{
			name: "abctwo1",
			args: "abctwo1",
			want: 1,
		},
		{
			name: "abc1one",
			args: "abc1one",
			want: 1,
		},
		{
			name: "abc1two",
			args: "abc1two",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLast(tt.args); got != tt.want {
				t.Errorf("getLast() = %v, want %v", got, tt.want)
			}
		})
	}
}
