package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	lines := readInput("test_part1.txt")
	result := part1(lines)
	var want int = 5
	if result != want {
		t.Errorf("Incorrect result, got: %d, want: %d\n", result, want)
	}
}

func TestPart2(t *testing.T) {
	lines := readInput("test_part2.txt")
	result := part2(lines)
	var want int = 12
	if result != want {
		t.Errorf("Incorrect result, got: %d, want: %d\n", result, want)
	}
}
