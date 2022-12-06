package main

import "testing"

func Test_findMarker(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: struct{ line string }{line: "nppdvjthqldpwncqszvftbrmjlhg"}, want: 23},
		{name: "test2", args: struct{ line string }{line: "bvwbjplbgvbhsrlpgdmjqwftvncz"}, want: 5},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMarker(tt.args.line); got != tt.want {
				t.Errorf("findMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}
