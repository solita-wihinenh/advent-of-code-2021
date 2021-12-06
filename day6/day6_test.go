package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	fishes := readInput("test_part1.txt")
	days := 80
	result := part1(fishes, days)
	var want int = 5934
	if result != want {
		t.Errorf("Incorrect result, got: %d, want: %d\n", result, want)
	}
}

func test_part2(t *testing.T) {
	fishes := readInput("test_part2.txt")
	days := 256
	result := part2(fishes, days)
	var want int = 26984457539
	if result != want {
		t.Errorf("Incorrect result, got: %d, want: %d\n", result, want)
	}
}
