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

	input := readInput("input.txt")
	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Reading lines from file failed: %v\n", err)
	}

	return lines
}

func part1(lines []string) int64 {
	sums := getSums(lines)
	gamma, epsilon := getGammaAndEpsilon(sums)
	gammaRate, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilon, 2, 64)
	return gammaRate * epsilonRate
}

func getSums(lines []string) []int {
	var sums []int
	for i := 0; i < len(lines[0]); i++ {
		sum := 0
		for _, line := range lines {
			if line[i] == '1' {
				sum++
			} else {
				sum--
			}
		}
		sums = append(sums, sum)
	}
	return sums
}

func getGammaAndEpsilon(sums []int) (string, string) {
	var gamma strings.Builder
	var epsilon strings.Builder
	for _, sum := range sums {
		if sum > 0 {
			gamma.WriteString("1")
			epsilon.WriteString("0")
		} else {
			gamma.WriteString("0")
			epsilon.WriteString("1")
		}
	}
	return gamma.String(), epsilon.String()
}

func part2(lines []string) int64 {
	oxygen := filter(lines, 0, '1', '0')[0]
	co2 := filter(lines, 0, '0', '1')[0]
	oxygenRating, _ := strconv.ParseInt(oxygen, 2, 64)
	co2Rating, _ := strconv.ParseInt(co2, 2, 64)
	return oxygenRating * co2Rating
}

func filter(lines []string, pos int, mostCommon byte, leastCommon byte) []string {
	sums := getSums(lines)
	var filtered []string

	for _, line := range lines {
		if (sums[pos] >= 0 && line[pos] == mostCommon) ||
			(sums[pos] < 0 && line[pos] == leastCommon) {
			filtered = append(filtered, line)
		}
	}

	if len(filtered) > 1 {
		filtered = filter(filtered, pos+1, mostCommon, leastCommon)
	}

	return filtered
}
