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
		{
			name: "Card 1",
			args: args{l: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"},
			want: 8,
		},
		{
			name: "Card 2",
			args: args{l: "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"},
			want: 2,
		},
		{
			name: "Card 3",
			args: args{l: "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"},
			want: 2,
		},
		{
			name: "Card 4",
			args: args{l: "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"},
			want: 1,
		},
		{
			name: "Card 5",
			args: args{l: "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"},
			want: 0,
		},
		{
			name: "Card 6",
			args: args{l: "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreLine(tt.args.l); got != tt.want {
				t.Errorf("scoreLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
