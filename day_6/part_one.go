package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func inArea(lines *[]string, x, y int) bool {
	width := len((*lines)[0])
	height := len(*lines)
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

func main() {
	fmt.Println("Part one:")

	lines := readFile("input.txt")
	total := 0
	direction := complex(0, -1)
	x, y := startingPosition(lines)
	var visited map[string]bool
	visited = make(map[string]bool)

	for {
		pos := string(x) + string(y)
		if !visited[pos] {
			visited[pos] = true
			total++
		}

		newX, newY := x+int(real(direction)), y+int(imag(direction))
		if !inArea(&lines, newX, newY) {
			break
		}
		front := string(lines[newY][newX])
		if front == "#" {
			direction = direction * complex(0, 1)
			continue
		}

		x, y = newX, newY
	}

	fmt.Println("Distinct positions: ", total)

}
