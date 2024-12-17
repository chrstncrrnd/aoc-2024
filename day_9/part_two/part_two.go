package part_two

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

type Fragment struct {
	id   int
	size int
}

func processInput(input string) []Fragment {
	out := make([]Fragment, 0)
	for i, char := range input {
		size, _ := strconv.Atoi(string(char))
		id := -1
		if i&1 == 0 {
			id = i / 2
		}
		out = append(out, Fragment{id, size})
	}
	return out
}

func processFragments(fragments []Fragment) []Fragment {
	out := make([]Fragment, 0)
	for j, frag := range fragments {
		if frag.id == -1 {
			sizeAvailable := frag.size
			for i := len(fragments) - 1; i >= j; i-- {
				lookingAt := fragments[i]
				if lookingAt.id == -1 {
					continue
				}
				if sizeAvailable == 0 {
					break
				}
				if lookingAt.size <= sizeAvailable {
					out = append(out, lookingAt)
					sizeAvailable -= lookingAt.size
					fragments[i].id = -1
				}
			}
			if sizeAvailable > 0 {
				out = append(out, Fragment{-1, sizeAvailable})
			}
		} else {
			out = append(out, frag)
		}
	}
	return out
}

func calculateChecksum(fragments []Fragment) int {
	checksum := 0
	index := 0
	for _, frag := range fragments {
		if frag.id != -1 {
			for j := 0; j < frag.size; j++ {
				checksum += (index + j) * frag.id
			}
		}
		index += frag.size
	}
	return checksum
}

func PartTwo() {
	contents, _ := readFile("./input.txt")
	fragments := processInput(contents)
	processed := processFragments(fragments)
	fmt.Println("Checksum: ", calculateChecksum(processed))
}
