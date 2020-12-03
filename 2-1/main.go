package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gopkg.in/gookit/color.v1"

	"github.com/JackRamey/GoAdvent2020/pkg/iohelpers"
)

var (
	fail = color.FgRed.Render
	pass = color.FgGreen.Render
)

func main() {
	start := time.Now()
	// Determine filepath
	cwd, err := os.Getwd()
	check(err)
	filepath := path.Join(cwd, "2-1", "input", "data")
	// Read input
	lines, err := iohelpers.FileToLines(filepath)
	check(err)

	numValid := 0
	for _, line := range lines {
		line := line
		min, max, char, pw := parseLine(line)
		count := strings.Count(pw, char)
		fmt.Printf("%2d-%2d %s: %-25s Count: ", min, max, char, pw)
		if count >= min && count <= max {
			fmt.Printf("%s\n", fail(count))
 			numValid += 1
		} else {
			fmt.Printf("%s\n", pass(count))
		}
	}

	fmt.Printf("Number of valid passwords: %d\n", numValid)
	fmt.Println("Execution Time: ", time.Since(start))
}

func parseLine(line string) (min, max int, character, password string) {
	var err error
	pattern := `(\d*)-(\d*)\s(\w):\s(\w*)$`
	matcher := regexp.MustCompile(pattern)

	results := matcher.FindStringSubmatch(line)
	minStr, maxStr := results[1], results[2]
	min, err = strconv.Atoi(minStr)
	check(err)
	max, err = strconv.Atoi(maxStr)
	check(err)
	character = results[3]
	password = results[4]

	return min, max, character, password
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

