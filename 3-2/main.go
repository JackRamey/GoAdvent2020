package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/JackRamey/GoAdvent2020/pkg/iohelpers"
)

type slope struct {
	slide int
	drop int
}

func main() {
	// Determine filepath
	cwd, err := os.Getwd()
	check(err)
	filepath := path.Join(cwd, "3-2", "input", "data")
	// Read input
	lines, err := iohelpers.FileToLines(filepath)
	check(err)

	totalHits := 1
	for _, input := range []slope{{1,1}, {3,1}, {5,1}, {7,1}, {1,2}} {
		numHits := countTreesHit(lines, len(lines[0]), input)
		fmt.Printf("Drop: %d, Slide %d, Trees hit: %d\n", input.drop, input.slide, numHits)
		totalHits = totalHits * numHits
	}

	fmt.Printf("Number of trees hit: %d\n", totalHits)
}

func countTreesHit(treeMap []string, width int, s slope) int {
	cursor := 0
	countHits := 0
	for idx, line := range treeMap {
		if idx % s.drop != 0 {
			continue
		}
		if strings.Split(line, "")[cursor%width] == "#" {
			countHits += 1
		}
		cursor += s.slide
	}
	return countHits
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
