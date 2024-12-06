package part_two

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(name string) []string {
	var out []string
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("Something went wrong with opening file: ", err)
		return []string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		out = append(out, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Something went wrong with scanner: ", err)
	}

	return out
}

var width int = -1
var height int = -1

func inArea(lines []string, x, y int) bool {
	if width == -1 {
		width = len((lines)[0])
	}
	if height == -1 {
		height = len(lines)
	}
	return x >= 0 && y >= 0 && y < height && x < width
}

func startingPosition(lines []string) (int, int) {
	for y, line := range lines {
		for x, ch := range line {
			if string(ch) == "^" {
				return x, y
			}
		}
	}
	return 0, 0
}

func testMap(lines []string) bool {
	var timesSeen map[string]int = make(map[string]int)
	looped := false
	traverse(lines, func(x, y int, char string, direction complex64) bool {
		pos := encodePos(x, y)
		timesSeen[pos] += 1
		// 10 is arbitrary but a smaller number isn't that much faster and 10 seemed big enough lol
		if timesSeen[pos] >= 10 {
			looped = true
			return false
		}
		return true
	})
	return looped
}

func addObstacle(lines []string, x, y int) (out []string) {
	for cy, line := range lines {
		newLine := ""
		for cx, ch := range line {
			var char string = string(ch)
			if cx == x && cy == y {
				char = "#"
			}
			newLine = newLine + char
		}
		out = append(out, newLine)
	}
	return
}

func allVisitedPositions(lines []string) (visitedPositions map[string]bool) {
	visitedPositions = make(map[string]bool)
	traverse(lines, func(x, y int, char string, _ complex64) bool {
		pos := encodePos(x, y)
		visitedPositions[pos] = true
		return true
	})
	return
}

func traverse(lines []string, atPoint func(x, y int, char string, direction complex64) bool) {
	direction := complex(0, -1)
	x, y := startingPosition(lines)
	for {
		cont := atPoint(x, y, string(lines[y][x]), complex64(direction))
		if !cont {
			break
		}
		newX, newY := x+int(real(direction)), y+int(imag(direction))
		if !inArea(lines, newX, newY) {
			break
		}
		front := string(lines[newY][newX])
		if front == "#" {
			direction = direction * complex(0, 1)
			continue
		}

		x, y = newX, newY
	}
}

func encodePos(x, y int) (out string) {
	out = strconv.Itoa(x) + "," + strconv.Itoa(y)
	return
}

func decodePos(pos string) (x, y int) {
	temp := strings.Split(pos, ",")
	x, _ = strconv.Atoi(temp[0])
	y, _ = strconv.Atoi(temp[1])
	return
}

func PartTwo() {
	lines := readFile("input.txt")
	possiblePlaces := allVisitedPositions(lines)
	var total int
	for pos, _ := range possiblePlaces {
		x, y := decodePos(pos)
		newLines := addObstacle(lines, x, y)
		if testMap(newLines) {
			total++
		}
	}
	fmt.Println("Total: ", total)
}
