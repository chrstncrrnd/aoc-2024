package part_one

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

func expand(input string) ([]int, int){
  chars := strings.Split(input, "")
  out := make([]int, 0)

  gaps := 0
  free := false
  currentId := 0
  for _, char := range chars{
    num, _ := strconv.Atoi(char)
    for j := 0; j < num; j ++ {
      if free{
        out = append(out, -1)
        gaps ++
      }else{
        out = append(out, currentId)
      }
    }
    if !free {
      currentId ++
    }
    free = !free
  }

  return out, gaps
}

func checksum(expanded []int, gaps int) int {
  csum := 0
  endTaken := 1
  for i := 0; i < len(expanded) - gaps; i ++{
    num := expanded[i]
    if num != -1 {
      csum += num * i
    }else{
      end := expanded[len(expanded) - endTaken]
      for end == -1 {
        endTaken ++
        end = expanded[len(expanded) - endTaken]
      }
      csum += expanded[len(expanded) - endTaken] * i
      endTaken ++
    }
  }
  return csum
}

func PartOne() {
  contents, err := readFile("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  expanded, gaps := expand(contents)
  cs := checksum(expanded, gaps)
  fmt.Println("Checksum: ", cs)

}


