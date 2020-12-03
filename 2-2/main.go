package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"time"

	"gopkg.in/gookit/color.v1"

	"github.com/JackRamey/GoAdvent2020/pkg/iohelpers"
)

var (
	fade = color.New(color.FgDarkGray).Render
	pass = color.New(color.FgBlack, color.BgGreen).Render
	okay = color.New(color.FgBlack, color.BgCyan).Render
	fail = color.New(color.FgBlack, color.BgRed).Render
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
		pos1, pos2, char, pw := parseLine(line)
		var count = 0
		var prettypw = pw
		var char1, char2 string
		for idx, c := range pw {
			switch {
			case idx+1 == pos1:
				char1 = string(c)
				if char == string(c) {
					count += 1
				}
			case idx+1 == pos2:
				char2 = string(c)
				if char == string(c) {
					count += 2
				}
			}
		}
		switch count {
		case 0:
			prettypw = fade(pw[:pos1-1]) + fail(char1) + fade(pw[pos1:pos2]) + fail(char2) + fade(pw[pos2:])
		case 1:
			prettypw = pw[:pos1-1] + pass(char1) + pw[pos1:pos2] + okay(char2) + pw[pos2:]
			numValid += 1
		case 2:
			prettypw = pw[:pos1-1] + okay(char1) + pw[pos1:pos2] + pass(char2) + pw[pos2:]
			numValid += 1
		case 3:
			prettypw = fade(pw[:pos1-1]) + fail(char1) + fade(pw[pos1:pos2]) + fail(char2) + fade(pw[pos2:])
		}
		fmt.Printf("%2d-%2d %s: %-25s\n", pos1, pos2, char, prettypw)
	}

	fmt.Printf("Number of valid passwords: %d\n", numValid)
	fmt.Println("Execution Time: ", time.Since(start))
}

func parseLine(line string) (pos1, pos2 int, character, password string) {
	var err error
	pattern := `(\d*)-(\d*)\s(\w):\s(\w*)$`
	matcher := regexp.MustCompile(pattern)

	results := matcher.FindStringSubmatch(line)
	minStr, maxStr := results[1], results[2]
	pos1, err = strconv.Atoi(minStr)
	check(err)
	pos2, err = strconv.Atoi(maxStr)
	check(err)
	character = results[3]
	password = results[4]

	return pos1, pos2, character, password
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
