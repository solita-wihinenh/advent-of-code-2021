package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	p         Point
	direction int
	length    int
}

func (line *Line) GetPoints(diagonals bool) (points []Point) {
	for i := 0; i < line.length; i++ {
		x, y := line.p.x, line.p.y
		switch line.direction {
		case ASCENDING:
			if !diagonals {
				return
			}
			x = line.p.x + i
			y = line.p.y - i
		case DESCENDING:
			if !diagonals {
				return
			}
			x = line.p.x + i
			y = line.p.y + i
		case HORIZONTAL:
			x = line.p.x + i
		case VERTICAL:
			y = line.p.y + i
		}
		points = append(points, Point{x, y})
	}
	return
}

const (
	HORIZONTAL = iota
	VERTICAL
	ASCENDING
	DESCENDING
)

func main() {
	lines := readInput("input.txt")
	fmt.Printf("Part1: %d\n", part1(lines))
	fmt.Printf("Part1: %d\n", part2(lines))
}

func readInput(filename string) []Line {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []Line{}
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.FieldsFunc(text, func(r rune) bool {
			return r == ',' || r == ' '
		})
		x1, _ := strconv.Atoi(fields[0])
		y1, _ := strconv.Atoi(fields[1])
		x2, _ := strconv.Atoi(fields[3])
		y2, _ := strconv.Atoi(fields[4])

		px, py := 0, 0
		p2x, p2y := 0, 0
		direction := HORIZONTAL
		length := 0
		if x1 < x2 {
			px = x1
			py = y1
			p2x = x2
			p2y = y2
			direction = HORIZONTAL
			length = x2 - x1 + 1
		} else if x1 > x2 {
			px = x2
			py = y2
			p2x = x1
			p2y = y1
			direction = HORIZONTAL
			length = x1 - x2 + 1
		} else if y1 < y2 {
			px = x1
			py = y1
			p2x = x2
			p2y = y2
			direction = VERTICAL
			length = y2 - y1 + 1
		} else {
			px = x2
			py = y2
			p2x = x1
			p2y = y1
			direction = VERTICAL
			length = y1 - y2 + 1
		}

		if px != p2x && py != p2y {
			if px < p2x && py < p2y {
				direction = DESCENDING
			} else {
				direction = ASCENDING
			}
		}

		line := Line{Point{px, py}, direction, length}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Reading lines from file failed: %v\n", err)
	}

	return lines
}

func part1(lines []Line) int {
	linemap := make(map[Point]int)
	for _, line := range lines {
		points := line.GetPoints(false)
		for _, p := range points {
			v, ok := linemap[p]
			if ok {
				linemap[p] = v + 1
			} else {
				linemap[p] = 1
			}
		}
	}

	return getDangerCount(&linemap)
}

func getDangerCount(linemap *map[Point]int) (count int) {
	for _, v := range *linemap {
		if v > 1 {
			count++
		}
	}
	return
}

func part2(lines []Line) int {
	linemap := make(map[Point]int)
	for _, line := range lines {
		points := line.GetPoints(true)
		for _, p := range points {
			v, ok := linemap[p]
			if ok {
				linemap[p] = v + 1
			} else {
				linemap[p] = 1
			}
		}
	}

	return getDangerCount(&linemap)
}
