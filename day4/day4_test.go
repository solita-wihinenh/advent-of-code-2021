package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	draws, tables := readInput("test_part1.txt")
	result := part1(draws, tables)
	var want int = 4512
	if result != want {
		t.Errorf("Incorrect result, got: %d, want: %d\n", result, want)
	}
}

func TestPart2(t *testing.T) {
	draws, tables := readInput("test_part2.txt")
	result := part2(draws, tables)
	var want int = 1924
	if result != want {
		t.Errorf("Incorrect result, got: %d, want: %d\n", result, want)
	}
}
