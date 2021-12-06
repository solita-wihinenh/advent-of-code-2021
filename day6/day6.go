package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fishes := readInput("input.txt")
	fmt.Printf("Part1: %d\n", part1(fishes, 80))
	fmt.Printf("Part1: %d\n", part2(fishes, 256))
}

func readInput(filename string) map[int]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fishes := make(map[int]int)
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Split(text, ",")
		for _, field := range fields {
			fish, _ := strconv.Atoi(field)
			addFish(&fishes, fish, 1)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Reading lines from file failed: %v\n", err)
	}

	return fishes
}

func part1(fishes map[int]int, days int) int {
	for i := 0; i < days; i++ {
		fishes = nextDay(fishes)
	}
	return countFishes(&fishes)
}

func nextDay(old map[int]int) map[int]int {
	new := make(map[int]int)
	for fish, amount := range old {
		timer := fish - 1
		if timer < 0 {
			addFish(&new, 6, amount)
			addFish(&new, 8, amount)
		} else {
			addFish(&new, timer, amount)
		}
	}
	return new
}

func addFish(fishes *map[int]int, fish, amount int) {
	v, ok := (*fishes)[fish]
	if ok {
		(*fishes)[fish] = v + amount
	} else {
		(*fishes)[fish] = amount
	}
}

func countFishes(fishes *map[int]int) (sum int) {
	for _, amount := range *fishes {
		sum += amount
	}
	return
}

func part2(fishes map[int]int, days int) int {
	return part1(fishes, days)
}
