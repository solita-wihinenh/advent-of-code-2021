package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Table struct {
	Numbers [][]string
	Marks   [][]bool
	Bingo   bool
}

func (t *Table) GetUnmarkedSum() int {
	sum := 0
	for x, row := range t.Numbers {
		for y, pos := range row {
			if !t.Marks[x][y] {
				number, _ := strconv.Atoi(pos)
				sum += number
			}
		}
	}
	return sum
}

func (t *Table) MarkNumber(draw string) bool {
	for x, row := range t.Numbers {
		for y, pos := range row {
			if draw == pos {
				t.Marks[x][y] = true
				return t.checkBingo(x, y)
			}
		}
	}
	return false
}

func (t *Table) checkBingo(x, y int) bool {
	return t.checkCol(x) || t.checkRow(y)
}

func (t *Table) checkCol(x int) bool {
	for y := 0; y < len(t.Numbers); y++ {
		if !t.Marks[x][y] {
			return false
		}
	}
	t.Bingo = true
	return true
}

func (t *Table) checkRow(y int) bool {
	for x := 0; x < len(t.Numbers[y]); x++ {
		if !t.Marks[x][y] {
			return false
		}
	}
	t.Bingo = true
	return true
}

func main() {
	draws, tables_p1 := readInput("input.txt")
	fmt.Printf("Part1: %d\n", part1(draws, tables_p1))
	draws, tables_p2 := readInput("input.txt")
	fmt.Printf("Part2: %d\n", part2(draws, tables_p2))
}

func readInput(filename string) ([]string, []Table) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	draws := strings.Split(scanner.Text(), ",")
	scanner.Scan()

	tables := []Table{}
	table := Table{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			tables = append(tables, table)
			table = Table{}
		} else {
			numbers := strings.Fields(line)
			table.Numbers = append(table.Numbers, numbers)
			marks := []bool{}
			for i := 0; i < len(numbers); i++ {
				marks = append(marks, false)
			}
			table.Marks = append(table.Marks, marks)
		}
	}
	tables = append(tables, table)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Reading lines from file failed: %v\n", err)
	}

	return draws, tables
}

func part1(draws []string, tables []Table) int {
	for _, draw := range draws {
		for i := range tables {
			if tables[i].MarkNumber(draw) {
				drawNumber, _ := strconv.Atoi(draw)
				return drawNumber * tables[i].GetUnmarkedSum()
			}
		}
	}
	return 0
}

func part2(draws []string, tables []Table) int {
	tablesWon := 0
	tableCount := len(tables)
	for _, draw := range draws {
		for i := range tables {
			if !tables[i].Bingo && tables[i].MarkNumber(draw) {
				tablesWon++
				if tablesWon == tableCount {
					drawNumber, _ := strconv.Atoi(draw)
					return drawNumber * tables[i].GetUnmarkedSum()
				}
			}
		}
	}
	return 0
}
