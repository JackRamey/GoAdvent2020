package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/JackRamey/GoAdvent2020/pkg/iohelpers"
)

func main() {
	// Determine filepath
	cwd, err := os.Getwd()
	check(err)
	filepath := path.Join(cwd, "3-1", "input", "data")
	// Read input
	lines, err := iohelpers.FileToLines(filepath)
	check(err)

	numHits := countTreesHit(lines, len(lines[0]), 1, 3)

	fmt.Printf("Number of trees hit: %d\n", numHits)
}

func countTreesHit(treeMap []string, width, drop, slide int) int {
	cursor := 0
	countHits := 0
	for idx, line := range treeMap {
		if idx % drop != 0 {
			continue
		}
		if strings.Split(line, "")[cursor%width] == "#" {
			countHits += 1
		}
		cursor += slide
	}
	return countHits
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
