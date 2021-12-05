package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInput("test_part1.txt")
	result := part1(input)
	var want int64 = 4512
	if result != want {
		t.Errorf("Incorrect result, got: %d, want: %d\n", result, want)
	}
}

// func TestPart2(t *testing.T) {
// 	input := readInput("test_part2.txt")
// 	result := part2(input)
// 	var want int64 = 230
// 	if result != want {
// 		t.Errorf("Incorrect result, got: %d, want: %d\n", result, want)
// 	}
// }
