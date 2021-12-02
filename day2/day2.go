package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	name     string
	argument int
}

func main() {

	input := readInput("input.txt")
	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

func readInput(filename string) []Command {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	commands := []Command{}
	for scanner.Scan() {
		command := parseCommand(scanner.Text())
		commands = append(commands, command)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Reading lines from file failed: %v\n", err)
	}

	return commands
}

func parseCommand(line string) Command {
	split := strings.Split(line, " ")
	if len(split) != 2 {
		log.Fatalf("Could not parse command from line: %s\n", line)
	}

	name := split[0]
	argument, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	return Command{name, argument}
}

func part1(commands []Command) int {
	depth, position := 0, 0
	for _, command := range commands {
		switch {
		case command.name == "down":
			depth += command.argument
		case command.name == "up":
			depth -= command.argument
		case command.name == "forward":
			position += command.argument
		}
	}
	return depth * position
}

func part2(commands []Command) int {
	depth, position, aim := 0, 0, 0
	for _, command := range commands {
		switch {
		case command.name == "down":
			aim += command.argument
		case command.name == "up":
			aim -= command.argument
		case command.name == "forward":
			position += command.argument
			depth += aim * command.argument
			if depth < 0 {
				depth = 0
			}
		}
	}
	return depth * position
}
