package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInput("test_part1.txt")
	result := part1(input)
	want := 7
	if result != want {
		t.Errorf("Incorrect amount of increments, got: %d, want: %d\n", result, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInput("test_part2.txt")
	result := part2(input)
	want := 5
	if result != want {
		t.Errorf("Incorrect amount of increments, got: %d, want: %d\n", result, want)
	}
}
