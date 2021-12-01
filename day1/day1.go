package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	input := readInput("input.txt")
	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

func readInput(filename string) []int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []int64{}
	for scanner.Scan() {
		line, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Reading lines from file failed: %v\n", err)
	}

	return lines
}

func part1(depths []int64) int {
	total := 0
	prev := depths[0]
	for _, value := range depths[1:] {
		if value > prev {
			total++
		}
		prev = value
	}
	return total
}

func part2(depths []int64) int {
	total := 0
	prev := getArraySum(depths[:3])
	for i := 3; i < len(depths); i++ {
		sum := getArraySum(depths[i-2 : i+1])
		if sum > prev {
			total++
		}
		prev = sum
	}
	return total
}

func getArraySum(array []int64) int64 {
	var sum int64 = 0
	for _, value := range array {
		sum += value
	}
	return sum
}
